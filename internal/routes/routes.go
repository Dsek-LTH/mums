package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/Dsek-LTH/mums/internal/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", handlers.Home)
	e.GET("/login", handlers.Login)
	e.GET("/admin", handlers.Admin)
	e.GET("/nolla", handlers.Nolla)
	e.GET("/phadder", handlers.Phadder)
}

