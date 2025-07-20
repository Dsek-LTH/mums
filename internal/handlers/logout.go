package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/pkg/httpx"
)

func PostLogout(ss *auth.SessionStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth.LogoutUser(c, ss)

		return httpx.Redirect(c, http.StatusSeeOther, "/login")
	}
}
