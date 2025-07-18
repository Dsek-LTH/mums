package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPhaddergruppSettings(c echo.Context) error {
	return c.Render(http.StatusOK, "phaddergrupp-settings", map[string]interface{}{})
}

func PostPhaddergruppSettings(c echo.Context) error {
	return c.Render(http.StatusOK, "phaddergrupp-settings", map[string]interface{}{})
}
