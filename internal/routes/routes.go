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

	phaddergrupp.GET("/:id", handlers.GetPhaddergrupp)
	phaddergrupp.DELETE("/:id", handlers.DeletePhaddergrupp)
	phaddergrupp.GET("/:id/event-stream", handlers.GetPhaddergruppEventStream)
	phaddergrupp.POST("/:id/purchase-mums", handlers.PostPhaddergruppPurchaseMums)
	phaddergrupp.POST("/:id/mumsa", handlers.PostPhaddergruppMumsa)
	phaddergrupp.POST("/:id/kick", handlers.PostPhaddergruppKick, auth.RequirePhaddergruppRole(roles.Phadder))
	phaddergrupp.POST("/:id/mums/adjust", handlers.PostPhaddergruppMumsAdjust, auth.RequirePhaddergruppRole(roles.Phadder))
	phaddergrupp.GET("/:id/settings", handlers.GetPhaddergruppSettings, auth.RequirePhaddergruppRole(roles.Phadder))
	phaddergrupp.PATCH("/:id/settings", handlers.PatchPhaddergruppSettings, auth.RequirePhaddergruppRole(roles.Phadder))
}
