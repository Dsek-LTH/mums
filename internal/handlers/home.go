package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/context"
	"github.com/Dsek-LTH/mums/internal/db"
	"github.com/Dsek-LTH/mums/internal/roles"
)

type homePageData struct {
	IsLoggedIn                            bool
	AllowedErrorCodes                     []int
	UserProfileName                       string
	UserPhaddergruppSummaries             []db.UserPhaddergruppSummary
	HasMoreThanOneUserPhaddergruppSummary bool
	PhaddergruppName                      string
	Errors                                map[string][]string
}

func GetHome(c echo.Context) error {
	database := db.GetDB(c)
	userAccountID := auth.GetUserAccountID(c)
	userProfile := context.GetUserProfile(c)

	userPhaddergruppSummaries, err := database.ReadUserPhaddergruppSummariesByUserAccountID(database, userAccountID)
	if err != nil {
		c.Logger().Errorf("Database error user phaddergrupp summaries read for user %s: %v", userAccountID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	pageData := homePageData{
		IsLoggedIn: auth.GetIsLoggedIn(c),
		AllowedErrorCodes: []int{http.StatusInternalServerError},
		UserProfileName: userProfile.Name,
		UserPhaddergruppSummaries: userPhaddergruppSummaries,
		HasMoreThanOneUserPhaddergruppSummary: len(userPhaddergruppSummaries) > 1,
	}
	return c.Render(http.StatusOK, "home", pageData)
}

func PostHome(c echo.Context) error {
	phaddergruppName := c.FormValue("phaddergrupp-name")

	unexpectedFormError := func() error {
		pageData := homePageData{
			PhaddergruppName: phaddergruppName,
			Errors:           map[string][]string{"Generic": {"An unexpected error occurred. Please try again."}},
		}
		return c.Render(http.StatusInternalServerError, "home#form-fields", pageData)
	}

	database := db.GetDB(c)

	phaddergruppID, err := database.CreatePhaddergrupp(database, phaddergruppName)
	if err != nil {
		c.Logger().Errorf("Database error during phaddergrupp creation: %v", err)
		return unexpectedFormError()	
	}
	userAccountID := auth.GetUserAccountID(c)
	err = database.CreatePhaddergruppMapping(database, userAccountID, phaddergruppID, roles.Phadder)
	if err != nil {
		c.Logger().Errorf("Database error during phaddergrupp mapping creation: %v", err)
		return unexpectedFormError()	
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/phaddergrupp/%d", phaddergruppID))
	return c.NoContent(http.StatusOK)
}
