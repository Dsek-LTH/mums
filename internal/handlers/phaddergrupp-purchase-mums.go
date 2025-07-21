package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/context"
	"github.com/Dsek-LTH/mums/internal/db"
	"github.com/Dsek-LTH/mums/internal/roles"
	"github.com/Dsek-LTH/mums/pkg/httpx"
)

func PostPhaddergruppPurchaseMums(c echo.Context) error {
	mumsPurchaseQuantityStr := c.FormValue("mums-purchase-quantity")
	if mumsPurchaseQuantityStr == "" {
	    return echo.NewHTTPError(http.StatusBadRequest, "mums-purchase-quantity is required")
	}
	mumsPurchaseQuantity, err := strconv.ParseInt(mumsPurchaseQuantityStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "mums-purchase-quantity must be a valid integer")
	}

	database := db.GetDB(c)
	userAccountID := auth.GetUserAccountID(c)
	phaddergruppID := auth.GetPhaddergruppID(c)
	phaddergruppRole := auth.GetPhaddergruppRole(c)
	phaddergruppData := context.GetPhaddergrupp(c)

	mumsAvailable, err := database.ReadMumsAvailable(database, userAccountID, phaddergruppID) 	
	if err != nil {
		c.Logger().Errorf("Database error during mums available read: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	if mumsPurchaseQuantity + mumsAvailable > phaddergruppData.MumsCapacityPerUser {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf(
			"Purchase exceeds allowed limit: current=%d, requested=%d, max=%d",
			mumsAvailable, mumsPurchaseQuantity, phaddergruppData.MumsCapacityPerUser,
		))
	}

	var mumsPrice float64
    if phaddergruppRole == roles.N0lla {
        mumsPrice = phaddergruppData.MumsPriceN0lla
    } else {
        mumsPrice = phaddergruppData.MumsPricePhadder
    }

	finalMumsPrice := mumsPrice * float64(mumsPurchaseQuantity)

    // The biggest of thanks to "t-shirt danne" for the Swish URL <3
	// https://app.swish.nu/1/p/sw/?sw=<number>&amt=<amount>&cur=<currency>&msg=<message>
	swishURL := &url.URL{
		Scheme: "https",
		Host:   "app.swish.nu",
		Path:   "/1/p/sw/",
	}

	params := url.Values{}
	params.Set("sw", phaddergruppData.SwishRecipientNumber)
	params.Set("amt", fmt.Sprintf("%.2f", finalMumsPrice))
	params.Set("cur", phaddergruppData.MumsCurrency)
	params.Set("msg", fmt.Sprintf("%s - %d mums", phaddergruppData.Name, mumsPurchaseQuantity))

	encodedQuery := params.Encode()
	finalQuery := strings.ReplaceAll(encodedQuery, "+", "%20")

	swishURL.RawQuery = finalQuery

	return httpx.Redirect(c, http.StatusSeeOther, swishURL.String())
}
