package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/db"
	"github.com/Dsek-LTH/mums/pkg/password"
)

type registerPageData struct {
	IsLoggedIn        bool
	AllowedErrorCodes []int
	Name              string
	Email             string
	Errors            map[string][]string
}

func GetRegister(c echo.Context) error {
	pageData := registerPageData{
		IsLoggedIn: auth.GetIsLoggedIn(c),
		AllowedErrorCodes: []int{http.StatusInternalServerError, http.StatusConflict, http.StatusBadRequest},
	}
	return c.Render(http.StatusOK, "register", pageData)
}

func PostRegister(ss *auth.SessionStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.FormValue("name")
		userEmail := c.FormValue("email")
		userPassword := c.FormValue("password")
		userConfirmPassword := c.FormValue("confirm-password")

		unexpectedError := func() error {
			pageData := registerPageData{
				IsLoggedIn: false,
				Name:       userName,
				Email:      userEmail,
				Errors:     map[string][]string{"Generic": {"An unexpected error occurred. Please try again."}},
			}
			return c.Render(http.StatusInternalServerError, "register#form-fields", pageData)
		}

		fieldError := func(statusCode int, field string, errorMessages []string) error {
			pageData := registerPageData{
				IsLoggedIn: false,
				Name:       userName,
				Email:      userEmail,
				Errors:     map[string][]string{field: errorMessages},
			}
			return c.Render(statusCode, "register#form-fields", pageData)
		}

		database := db.GetDB(c)

		emailExists, err := database.ReadUserCredentialsExistsByEmail(database, userEmail)
		if err != nil {
			c.Logger().Errorf("Database error during email conflict check for email %s: %v", userEmail, err)
			return unexpectedError()
		}
		if emailExists {
			return fieldError(http.StatusConflict, "Email", []string{"Account with email already exists."}) 
		}

		if userPassword != userConfirmPassword {
			return fieldError(http.StatusBadRequest, "PasswordConfirm", []string{"Passwords do not match."}) 
		}	
		hashword, err := password.HashSecure(userPassword)
		if err == bcrypt.ErrPasswordTooLong {
			return fieldError(http.StatusBadRequest, "PasswordConfirm", []string{"Passwords length exceeds 72."}) 
		}
		if err != nil {
			c.Logger().Errorf("Password could not be hashed: %v", err)
			return unexpectedError()
		}

		tx, err := database.Begin()
		if err != nil {
			c.Logger().Errorf("Failed to begin transaction during user creation: %v", err)
			return unexpectedError()
		}
		defer tx.Rollback()
		userCredentialsID, err := database.CreateUserCredentials(tx, userEmail, hashword)
		if err != nil {
			c.Logger().Errorf("Database error during user credentials creation during user creation: %v", err)
			return unexpectedError()
		}
		userProfileID, err := database.CreateUserProfile(tx, userName)
		if err != nil {
			c.Logger().Errorf("Database error during user profile creation during user creation: %v", err)
			return unexpectedError()
		}
		userAccountID, err := database.CreateUserAccount(tx, userCredentialsID, userProfileID)
		if err != nil {
			c.Logger().Errorf("Database error during user account creation during user creation: %v", err)
			return unexpectedError()
		}
		err = tx.Commit()
		if err != nil {
			c.Logger().Errorf("Database error during user creation: %v", err)
			return unexpectedError()
		}
		return auth.LoginUser(c, ss, userAccountID)
	}
}
