package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func Phaddergrupp(c echo.Context) error {
    return c.Render(http.StatusOK, "phaddergrupp", map[string]interface{}{})
}

