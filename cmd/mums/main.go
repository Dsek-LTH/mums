package main

import (
	"github.com/Dsek-LTH/mums/internal/auth"
	"github.com/Dsek-LTH/mums/internal/config"
	"github.com/Dsek-LTH/mums/internal/db"
	"github.com/Dsek-LTH/mums/internal/routes"
	"github.com/Dsek-LTH/mums/internal/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	database, err := db.NewDB(config.DBFilePath)
	if err != nil {
		panic(err)
	}
	e.Use(db.DBMiddleware(database))

	sessionStore := auth.NewSessionStore()
	e.Use(auth.SessionMiddleware(sessionStore))

	e.Use(auth.UserAccountRBACMiddleware())

	templates.LoadTemplates(e)

	routes.RegisterRoutes(e)

	e.Static("/static", "web/static")

	e.Start(config.ServerAddress)
}
