package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/handlers"
	"github.com/Dsek-LTH/mums/internal/roles"
)

func RegisterRoutes(e *echo.Echo, ss *auth.SessionStore) {
	e.Use(auth.SessionMiddleware(ss))

	e.GET("/login", handlers.GetLogin)
	e.POST("/login", handlers.PostLogin(ss))
	e.GET("/register", handlers.GetRegister)
	e.POST("/register", handlers.PostRegister(ss))
	e.GET("/about", handlers.GetAbout)
	e.GET("/work-in-progress", handlers.GetWorkInProgress)

	protected := e.Group("")
	protected.Use(auth.RequireSession())
	protected.Use(auth.UserAccountRBACMiddleware())

	protected.GET("/", handlers.GetHome)
	protected.POST("/logout", handlers.PostLogout(ss))
	protected.GET("/admin", handlers.GetAdmin, auth.RequireUserAccountRole(roles.Admin))
	protected.POST("/admin", handlers.PostAdmin, auth.RequireUserAccountRole(roles.Admin))
	protected.GET("/invite", handlers.GetInvite)

	phaddergrupp := protected.Group("/phaddergrupp", auth.PhaddergruppRBACMiddleware())

	phaddergrupp.GET("/:phaddergrupp_id", handlers.GetPhaddergrupp, auth.RequirePhaddergruppRole(roles.Nolla, roles.Phadder))
	phaddergrupp.POST("/:phaddergrupp_id", handlers.PostPhaddergrupp, auth.RequirePhaddergruppRole(roles.Phadder))
	phaddergrupp.GET("/:phaddergrupp_id/settings", handlers.GetPhaddergruppSettings, auth.RequirePhaddergruppRole(roles.Phadder))
	phaddergrupp.POST("/:phaddergrupp_id/settings", handlers.PostPhaddergruppSettings, auth.RequirePhaddergruppRole(roles.Phadder))
	phaddergrupp.GET("/:phaddergrupp_id/event_stream", handlers.StreamPhaddergruppEvents, auth.RequirePhaddergruppRole(roles.Nolla, roles.Phadder))
}
