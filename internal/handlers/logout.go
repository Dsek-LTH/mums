package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
)

func PostLogout(ss *auth.SessionStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth.LogoutUser(c, ss)

		c.Response().Header().Set("HX-Redirect", "/login")
		return c.NoContent(http.StatusOK)
	}
}
