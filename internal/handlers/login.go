package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
	
	"github.com/Dsek-LTH/mums/internal/auth"
)

func GetLogin(c echo.Context) error {
    return c.Render(http.StatusOK, "login", map[string]any{})
}

func PostLogin(ss *auth.SessionStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

