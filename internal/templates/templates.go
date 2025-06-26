package templates

import (
	"embed"
	"errors"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:embed **/*.tmpl
var tmplFS embed.FS

type TemplateRenderer struct {
    templates map[string]*template.Template
}

func (tR *TemplateRenderer) Render(w io.Writer, name string, data any, _ echo.Context) error {
	tmpl, ok := tR.templates[name]
	if !ok {
		return errors.New("template not found")
	}

    return tmpl.ExecuteTemplate(w, "base", data)
}

func NewTemplateRenderer() *TemplateRenderer {
	templates := make(map[string]*template.Template)

	base, err := template.ParseFS(tmplFS, "layouts/base.tmpl")
	if err != nil {
		panic("failed to parse base layout: " + err.Error())
	}

	viewFiles, err := fs.Glob(tmplFS, "views/*.tmpl")
	if err != nil {
		panic("failed to glob view templates: " + err.Error())
	}

	for _, viewFile := range viewFiles {
		tmplName := strings.TrimSuffix(filepath.Base(viewFile), ".tmpl")

		baseCopy, err := base.Clone()
		if err != nil {
			panic("failed to clone base template: " + err.Error())
		}

		tmpl, err := baseCopy.ParseFS(tmplFS, viewFile, "partials/*.tmpl")
		if err != nil {
			panic("failed to parse view template: " + err.Error())
		}

		templates[tmplName] = tmpl
	}

	return &TemplateRenderer{templates}
}

func LoadTemplates(e *echo.Echo) {
    e.Renderer = NewTemplateRenderer()
}

