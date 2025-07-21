package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/config"
	"github.com/Dsek-LTH/mums/internal/context"
	"github.com/Dsek-LTH/mums/internal/db"
	"github.com/Dsek-LTH/mums/internal/roles"
)

type phaddergruppPageData struct {
	IsLoggedIn                bool
	AllowedErrorCodes         []int
	PhaddergruppID            int64
	PhaddergruppRole          roles.PhaddergruppRole
	MumsAvailable             int64
	db.UserProfileData
	db.PhaddergruppData 
	PhaddergruppUserSummaries []db.PhaddergruppUserSummary
	MumsCapacityReached       bool
	MumsPurchaseQuantities    []int
	InviteURLN0lla            string
	InviteURLPhadder          string
}

func mumsPurchaseQuantities(mumsAvailable, mumsCapacityPerUser int64) []int {
	remainingMumsCapacity := mumsCapacityPerUser - mumsAvailable

	var purchaseQuantities []int
	for qty := 1; qty <= min(config.MumsMaxPurchaseQuantity, int(remainingMumsCapacity)); qty++ {
		purchaseQuantities = append(purchaseQuantities, qty)
	}

	return purchaseQuantities
}

func GetPhaddergrupp(c echo.Context) error {
	userAccountID := auth.GetUserAccountID(c)
	phaddergruppID := auth.GetPhaddergruppID(c)
	phaddergruppRole := auth.GetPhaddergruppRole(c)
	database := db.GetDB(c)

	mumsAvailable, err := database.ReadMumsAvailable(database, userAccountID, phaddergruppID)
	if err != nil {
		c.Logger().Errorf("Database error during mums available read: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}
	
	phaddergruppUserSummaries, err := database.ReadPhaddergruppUserSummariesByPhaddergruppID(database, phaddergruppID)
	if err != nil {
		c.Logger().Errorf("Database error during phaddergrupp user summary read: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}
	
	purchaseQuantities := mumsPurchaseQuantities(mumsAvailable, config.MumsMaxPurchaseQuantity)

	inviteTokens, err := database.ReadPhaddergruppInviteTokensByPhaddergruppID(database, phaddergruppID)
	if err != nil {
		c.Logger().Errorf("Database error during invite tokens read read: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	inviteURLN0lla := config.PhaddergruppInviteURLBase + inviteTokens.N0lla
	inviteURLPhadder := config.PhaddergruppInviteURLBase + inviteTokens.Phadder

	pageData := phaddergruppPageData{
		IsLoggedIn: auth.GetIsLoggedIn(c),
		AllowedErrorCodes: []int{http.StatusInternalServerError},
		PhaddergruppID: phaddergruppID,
		PhaddergruppRole: phaddergruppRole,
		MumsAvailable: mumsAvailable,
		UserProfileData: context.GetUserProfile(c),
		PhaddergruppData: context.GetPhaddergrupp(c),
		PhaddergruppUserSummaries: phaddergruppUserSummaries,
		MumsCapacityReached: mumsAvailable >= config.MumsMaxPurchaseQuantity,
		MumsPurchaseQuantities: purchaseQuantities,
		InviteURLN0lla: inviteURLN0lla,
		InviteURLPhadder: inviteURLPhadder,
	}
	return c.Render(http.StatusOK, "phaddergrupp", pageData)
}

func PostPhaddergrupp(c echo.Context) error {
	return nil
}
