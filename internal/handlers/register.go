package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
)

func GetRegister(c echo.Context) error {
    return c.Render(http.StatusOK, "register", map[string]any{})
}

func PostRegister(ss *auth.SessionStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
