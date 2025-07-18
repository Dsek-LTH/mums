package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
)

type aboutPageData struct {
	IsLoggedIn        bool
	AllowedErrorCodes []int
}

func GetAbout(c echo.Context) error {
	pageData := aboutPageData{
		IsLoggedIn: auth.GetIsLoggedIn(c),
	}
	return c.Render(http.StatusOK, "about", pageData)
}
