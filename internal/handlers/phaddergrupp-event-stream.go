package handlers

import (
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/config"
	"github.com/Dsek-LTH/mums/internal/context"
	"github.com/Dsek-LTH/mums/internal/db"
	"github.com/Dsek-LTH/mums/internal/roles"
	"github.com/Dsek-LTH/mums/pkg/httpx"
)

type templateDataMumsAvailableWidget struct {
	PhaddergruppID         int64
	db.PhaddergruppData 
	MumsAvailable          int64
	HasMumsAvailable 	   bool
	MumsCapacityReached    bool
	MumsPurchaseQuantities []int
}

func emitMumsAvailableWidgetUpdate(c echo.Context, eventData db.MumsAvailableUpdate) error {
	phaddergruppID := auth.GetPhaddergruppID(c)
	phaddergruppData := context.GetPhaddergrupp(c)
	purchaseQuantities := mumsPurchaseQuantities(eventData.MumsAvailable, phaddergruppData.MumsCapacityPerUser)

	templateData := templateDataMumsAvailableWidget {
		PhaddergruppID:         phaddergruppID,
		PhaddergruppData:       phaddergruppData,
		MumsAvailable:          eventData.MumsAvailable,
		HasMumsAvailable: 		eventData.MumsAvailable > 0,
		MumsCapacityReached:    eventData.MumsAvailable >= phaddergruppData.MumsCapacityPerUser,
		MumsPurchaseQuantities: purchaseQuantities,
	}

	var sb strings.Builder
	if err := c.Echo().Renderer.Render(&sb, "phaddergrupp#mums-available-widget", templateData, c); err != nil {
		c.Logger().Errorf("template render error: %v", err)
		return nil
	}

	return httpx.EmitSSE(c, "mums-available-widget-update", sb.String())
}

type mumsAvailableBadgeTemplateData struct {
	UserAccountID int64
	DoOOB 	      bool
	MumsAvailable int64
}

func emitMumsAvailableBadgeUpdate(c echo.Context, eventData db.MumsAvailableUpdate) {
	templateData := mumsAvailableBadgeTemplateData{
		UserAccountID: eventData.UserAccountID,
		DoOOB:         true,
		MumsAvailable: eventData.MumsAvailable,
	}

	var sb strings.Builder
	if err := c.Echo().Renderer.Render(&sb, "phaddergrupp#mums-available-badge", templateData, c); err != nil {
		c.Logger().Errorf("template render error: %v", err)
		return
	}

	httpx.EmitSSE(c, "mums-available-badge-update", sb.String())
}

func handlePhaddergruppEvent(c echo.Context, event db.DBEvent) {
	if event.Type != db.DBUpdate || event.Table != "phaddergrupp_mappings" {
		return
	}

	eventData := event.Data.(db.MumsAvailableUpdate)

	phaddergruppID := auth.GetPhaddergruppID(c)

	if eventData.PhaddergruppID != phaddergruppID {
		return
	}
	
	if eventData.UserAccountID == auth.GetUserAccountID(c) {
		emitMumsAvailableWidgetUpdate(c, eventData)
	}

	if auth.GetPhaddergruppRole(c) == roles.Phadder {
		emitMumsAvailableBadgeUpdate(c, eventData)
	}
}

func GetPhaddergruppEventStream(c echo.Context) error {
	database := db.GetDB(c)

	httpx.SetupSSE(c)

	subID, events := database.Subscribe(config.DBEventChannelBufferSize)
	defer database.Unsubscribe(subID)

	for {
		select {
		case <-c.Request().Context().Done():
			return nil
		case event, ok := <-events:
			if !ok {
				return nil
			}

			handlePhaddergruppEvent(c, event)
		}
	}
}
