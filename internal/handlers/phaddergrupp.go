package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type phaddergruppPageData struct {
	IsLoggedIn                bool
	AllowedErrorCodes         []int
	UserProfileName           string
	PhaddergruppName          string
	Errors                    map[string][]string
}

func GetPhaddergrupp(c echo.Context) error {
	pageData := phaddergruppPageData{}
	return c.Render(http.StatusOK, "phaddergrupp", pageData)
}

func PostPhaddergrupp(c echo.Context) error {
	return nil
}
