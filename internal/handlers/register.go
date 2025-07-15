package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func GetRegister(c echo.Context) error {
    return c.Render(http.StatusOK, "register", map[string]interface{}{})
}

func PostRegister(c echo.Context) error {
    return nil
}
