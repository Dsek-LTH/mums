package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/handlers"
	"github.com/Dsek-LTH/mums/internal/roles"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", handlers.GetHome)

	e.GET("/admin", handlers.GetAdmin, auth.RequireUserAccountRole(roles.Admin))
	e.POST("/admin", handlers.PostAdmin, auth.RequireUserAccountRole(roles.Admin))

	e.GET("/login", handlers.GetLogin)
	e.POST("/login", handlers.PostLogin)

	e.GET("/register", handlers.GetRegister)
	e.POST("/register", handlers.PostRegister)

	e.GET("/invite", handlers.GetInvite)

	phaddergrupp := e.Group("/phaddergrupp", auth.PhaddergruppRBACMiddleware())

	phaddergrupp.GET("/:phaddergrupp_id", handlers.GetPhaddergrupp, auth.RequirePhaddergruppRole(roles.Nolla, roles.Phadder))
	phaddergrupp.POST("/:phaddergrupp_id", handlers.PostPhaddergrupp, auth.RequirePhaddergruppRole(roles.Phadder))

	phaddergrupp.GET("/:phaddergrupp_id/settings", handlers.GetPhaddergruppSettings, auth.RequirePhaddergruppRole(roles.Phadder))
	phaddergrupp.POST("/:phaddergrupp_id/settings", handlers.PostPhaddergruppSettings, auth.RequirePhaddergruppRole(roles.Phadder))

	phaddergrupp.GET("/:phaddergrupp_id/event_stream", handlers.StreamPhaddergruppEvents, auth.RequirePhaddergruppRole(roles.Nolla, roles.Phadder))
}
