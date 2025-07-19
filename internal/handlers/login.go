package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/db"
	"github.com/Dsek-LTH/mums/pkg/password"
)

type loginPageData struct {
	IsLoggedIn        bool
	AllowedErrorCodes []int
	Email             string
	Errors            map[string][]string
}

func GetLogin(c echo.Context) error {
	pageData := loginPageData{
		IsLoggedIn:        false,
		AllowedErrorCodes: []int{http.StatusUnauthorized, http.StatusInternalServerError},
	}
	return c.Render(http.StatusOK, "login", pageData)
}

func PostLogin(ss *auth.SessionStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		userEmail := c.FormValue("email")
		userPassword := c.FormValue("password")

		unexpectedError := func() error {
			pageData := registerPageData{
				IsLoggedIn: false,
				Email:      userEmail,
				Errors:     map[string][]string{"Generic": {"An unexpected error occurred. Please try again."}},
			}
			return c.Render(http.StatusInternalServerError, "login#form-fields", pageData)
		}

		database := db.GetDB(c)

		userCredentialsID, hashword, err := database.ReadUserCredentialsIDAndHashwordByEmail(database, userEmail)
		if err == sql.ErrNoRows || !password.Check(userPassword, hashword) {
			pageData := loginPageData{
				IsLoggedIn: false,
				Email:      userEmail,
				Errors:     map[string][]string{"Generic": {"Invalid email or password."}},
			}
			return c.Render(http.StatusUnauthorized, "login#form-fields", pageData)
		}
		if err != nil {
			c.Logger().Errorf("Database error during login for email %s: %v", userEmail, err)
			return unexpectedError()
		}

		userAccountID, err := database.ReadUserAccountIDByUserCredentialsID(database, userCredentialsID)
		if err != nil {
			c.Logger().Errorf("CRITICAL: Credentials found (ID: %d) but no matching user account.", userCredentialsID)
		}

		return auth.LoginUser(c, ss, userAccountID)
	}
}
