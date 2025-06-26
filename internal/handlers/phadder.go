package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func Phadder(c echo.Context) error {
    return c.Render(http.StatusOK, "phadder", map[string]interface{}{})
}

