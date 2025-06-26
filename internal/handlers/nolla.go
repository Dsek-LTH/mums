package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func Nolla(c echo.Context) error {
    return c.Render(http.StatusOK, "nolla", map[string]interface{}{})
}

