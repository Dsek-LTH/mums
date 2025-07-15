package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func GetLogin(c echo.Context) error {
    return c.Render(http.StatusOK, "login", map[string]interface{}{})
}

func PostLogin(c echo.Context) error {
    return nil
}

