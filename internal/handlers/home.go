package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
)

type homePageData struct {
	IsLoggedIn bool
	AllowedErrorCodes []int
}

func GetHome(c echo.Context) error {
	pageData := homePageData{
		IsLoggedIn: auth.GetIsLoggedIn(c),
	}
    return c.Render(http.StatusOK, "home", pageData)
}
