package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
)

func PostLogout(ss *auth.SessionStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		return auth.LogoutUser(c, ss)
	}
}
