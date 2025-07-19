package context

import (
	"fmt"
	"net/http"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/config"
	"github.com/Dsek-LTH/mums/internal/db"
	"github.com/labstack/echo/v4"
)

func InjectPhaddergrupp() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			phaddergruppID := auth.GetPhaddergruppID(c)
			database := db.GetDB(c)

			phaddergruppData, err := database.ReadPhaddergrupp(database, phaddergruppID)
			if err != nil {
				c.Logger().Errorf("Database error during phaddergrupp read: %v", err)
                return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
			}
			c.Set(config.CTXKeyPhaddergrupp, phaddergruppData)

			return next(c)
		}
	}
}

func GetPhaddergrupp(c echo.Context) db.PhaddergruppData {
	phaddergruppData, ok := c.Get(config.CTXKeyPhaddergrupp).(db.PhaddergruppData)
	if !ok {
		panic("config.CTXKeyPhaddergrupp is not set in context, was InjectPhaddergrupp not applied?")
	}

	return phaddergruppData
}
