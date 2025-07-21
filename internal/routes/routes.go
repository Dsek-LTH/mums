package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/context"
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

	protected := e.Group(
		"",
		auth.RequireSession(),
		auth.UserAccountRBACMiddleware(),
		context.InjectUserProfile(),
	)

	protected.GET("/", handlers.GetHome)
	protected.POST("/", handlers.PostHome)
	protected.POST("/logout", handlers.PostLogout(ss))
	protected.GET("/admin", handlers.GetAdmin, auth.RequireUserAccountRole(roles.Admin))
	protected.POST("/admin", handlers.PostAdmin, auth.RequireUserAccountRole(roles.Admin))
	protected.GET("/invite/:token", handlers.GetInvite)

	phaddergrupp := protected.Group(
		"/phaddergrupp",
		auth.PhaddergruppRBACMiddleware(),
		context.InjectPhaddergrupp(),
	)

	phaddergrupp.GET("/:id", handlers.GetPhaddergrupp, auth.RequirePhaddergruppRole(roles.N0lla, roles.Phadder))
	phaddergrupp.POST("/:id", handlers.PostPhaddergrupp, auth.RequirePhaddergruppRole(roles.Phadder))
	phaddergrupp.GET("/:id/settings", handlers.GetPhaddergruppSettings, auth.RequirePhaddergruppRole(roles.Phadder))
	phaddergrupp.POST("/:id/settings", handlers.PostPhaddergruppSettings, auth.RequirePhaddergruppRole(roles.Phadder))
	phaddergrupp.GET("/:id/event-stream", handlers.StreamPhaddergruppEvents, auth.RequirePhaddergruppRole(roles.N0lla, roles.Phadder))
}
