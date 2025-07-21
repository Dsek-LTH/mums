package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/context"
	"github.com/Dsek-LTH/mums/internal/config"
	"github.com/Dsek-LTH/mums/internal/db"
)

func PostPhaddergruppMumsa(c echo.Context) error {
	database := db.GetDB(c)
	userAccountID := auth.GetUserAccountID(c)
	phaddergruppID := auth.GetPhaddergruppID(c)

	mumsAvailable, err := database.UpdateAdjustMumsAvailable(database, userAccountID, phaddergruppID, -1)
	if err != nil {
		c.Logger().Errorf("Database error during mums available update for user %s in phaddergrupp %s: %v", userAccountID, phaddergruppID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	phaddergruppData := context.GetPhaddergrupp(c)
	purchaseQuantities := mumsPurchaseQuantities(mumsAvailable, context.GetPhaddergrupp(c).MumsCapacityPerUser)

	pageData := phaddergruppPageData{
		PhaddergruppID: phaddergruppID,
		PhaddergruppData: phaddergruppData,
		MumsAvailable: mumsAvailable,
		MumsCapacityReached: mumsAvailable >= config.MumsMaxPurchaseQuantity,
		MumsPurchaseQuantities: purchaseQuantities,
	}
	return c.Render(http.StatusOK, "phaddergrupp#mums", pageData)
}
