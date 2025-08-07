package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/memagu/mums/internal/config"
	"github.com/memagu/mums/internal/db"
	"github.com/memagu/mums/internal/roles"
)

func UserAccountRBACMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			database := db.GetDB(c)
			userAccountRoles, err := database.ReadUserAccountRoles(database, GetUserAccountID(c))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
			}
			c.Set(config.CTXKeyUserAccountRoles, userAccountRoles)

			isSuperAdmin := slices.Contains(userAccountRoles, roles.SuperAdmin)
			c.Set(config.CTXKeyIsSuperAdmin, isSuperAdmin)

			return next(c)
		}
	}
}

func GetUserAccountRoles(c echo.Context) []roles.UserAccountRole {
	userAccountRoles, ok := c.Get(config.CTXKeyUserAccountRoles).([]roles.UserAccountRole)
	if !ok {
		panic("config.CTXKeyUserAccountRoles is not set in context, was UserAccountRBACMiddleware not applied?")
	}

	return userAccountRoles
}

func GetIsSuperAdmin(c echo.Context) bool {
	isSuperAdmin, ok := c.Get(config.CTXKeyIsSuperAdmin).(bool)
	if !ok {
		panic("config.CTXKeyIsSuperAdmin is not set in context, was UserAccountRBACMiddleware not applied?")
	}

	return isSuperAdmin
}

func RequireUserAccountRole(allowedUserAccountRoles ...roles.UserAccountRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if GetIsSuperAdmin(c) {
				return next(c)
			}

			for _, userAccountRole := range GetUserAccountRoles(c) {
				if slices.Contains(allowedUserAccountRoles, userAccountRole) {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "Forbidden: User is missing a required user account role")
		}
	}
}

func PhaddergruppRBACMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			phaddergruppIDString := c.Param("id")
			if phaddergruppIDString == "" {
				return echo.NewHTTPError(http.StatusBadRequest, "Bad Request: Missing phaddergrupp-id parameter")
			}
			phaddergruppID, err := strconv.ParseInt(phaddergruppIDString, 10, 64)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Bad Request: Invalid phaddergrupp-id")
			}
			database := db.GetDB(c)
			phaddergruppRole, err := database.ReadPhaddergruppRole(database, GetUserAccountID(c), phaddergruppID)
			if err != nil {
				if err == sql.ErrNoRows {
					return echo.NewHTTPError(http.StatusForbidden, "Forbidden: User account does not have access to this phaddergrupp")
				}
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
			}

			c.Set(config.CTXKeyPhaddergruppID, phaddergruppID)
			c.Set(config.CTXKeyPhaddergruppRole, phaddergruppRole)

			return next(c)
		}
	}
}

func GetPhaddergruppID(c echo.Context) int64 {
	phaddergruppID, ok := c.Get(config.CTXKeyPhaddergruppID).(int64)
	if !ok {
		panic("config.CTXKeyPhaddergruppID is not set in context, was PhaddergruppRBACMiddleware not applied?")
	}

	return phaddergruppID
}

func GetPhaddergruppRole(c echo.Context) roles.PhaddergruppRole {
	phaddergruppRole, ok := c.Get(config.CTXKeyPhaddergruppRole).(roles.PhaddergruppRole)
	if !ok {
		panic("config.CTXKeyPhaddergruppRole is not set in context, was PhaddergruppRBACMiddleware not applied?")
	}

	return phaddergruppRole
}

func RequirePhaddergruppRole(allowedPhaddergruppRoles ...roles.PhaddergruppRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if GetIsSuperAdmin(c) {
				return next(c)
			}

			if slices.Contains(allowedPhaddergruppRoles, GetPhaddergruppRole(c)) {
				return next(c)
			}

			return echo.NewHTTPError(http.StatusForbidden, "Forbidden: User is missing a required phaddergrupp role")
		}
	}
}
