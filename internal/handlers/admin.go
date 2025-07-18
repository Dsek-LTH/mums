package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAdmin(c echo.Context) error {
	return c.Render(http.StatusOK, "admin", map[string]interface{}{})
}

func PostAdmin(c echo.Context) error {
	return c.Render(http.StatusOK, "admin", map[string]interface{}{})
}
