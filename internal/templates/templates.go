package templates

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/memagu/mums/internal/config"
	"github.com/labstack/echo/v4"
)

//go:embed **/*.tmpl
var tmplFS embed.FS

type TemplateRenderer struct {
	templates map[string]*template.Template
}

func dict(values ...any) (map[string]any, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("dict requires an even number of arguments")
	}
	dict := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func NewTemplateRenderer() *TemplateRenderer {
	funcMap := template.FuncMap{
		"dict": dict,
	}
	templates := make(map[string]*template.Template)

	base, err := template.New("base").Funcs(funcMap).ParseFS(tmplFS, "layouts/*.tmpl", "components/*.tmpl")
	if err != nil {
		panic("failed to parse base layouts and partials: " + err.Error())
	}

	pageFiles, err := fs.Glob(tmplFS, "pages/*.tmpl")
	if err != nil {
		panic("failed to glob page template files: " + err.Error())
	}

	for _, pageFile := range pageFiles {
		tmplName := strings.TrimSuffix(filepath.Base(pageFile), ".tmpl")

		baseCopy, err := base.Clone()
		if err != nil {
			panic("failed to clone base template for " + tmplName + ": " + err.Error())
		}

		tmpl, err := baseCopy.ParseFS(tmplFS, pageFile)
		if err != nil {
			panic("failed to parse page template " + pageFile + ": " + err.Error())
		}

		templates[tmplName] = tmpl
	}

	return &TemplateRenderer{templates}
}

// name format: pageName or pageName#blockName
func (tr *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	pageName, blockName, isBlockRender := strings.Cut(name, config.TemplateBlockRenderSeparator)
	if !isBlockRender {
		blockName = pageName
	}
	if blockName == "" {
		return fmt.Errorf("error rendering template '%s': block name cannot be empty", name)
	}

	tmpl, ok := tr.templates[pageName]
	if !ok {
		return fmt.Errorf("template with pageName '%s' not found", pageName)
	}

	return tmpl.ExecuteTemplate(w, blockName, data)
}

func LoadTemplates(e *echo.Echo) {
	e.Renderer = NewTemplateRenderer()
}
