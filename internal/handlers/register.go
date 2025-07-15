package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
    return c.Render(http.StatusOK, "register", map[string]interface{}{})
}
