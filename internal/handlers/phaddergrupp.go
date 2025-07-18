package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPhaddergrupp(c echo.Context) error {
	return c.Render(http.StatusOK, "phaddergrupp", map[string]interface{}{})
}

func PostPhaddergrupp(c echo.Context) error {
	return nil
}
