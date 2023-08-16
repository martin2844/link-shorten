package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lynxsecurity/viper"
	"github.com/martin2844/link-shorten/pkg/db"
	"github.com/martin2844/link-shorten/pkg/handlers"
)

func main() {
	// Read in environment variables
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	// Start new echo instance
	e := echo.New()
	// Initialize test endpoints
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})
	// Initialize database
	db.Init()
	db.AutoMigrate()

	// Create Routing groups - similar to express router
	g := e.Group("")
	g2 := e.Group("/create")
	// Register handlers - pass the router we just created into the handler package to register the endpoints.
	handlers.RegisterBaseGroup(g)
	handlers.RegisterCreateGroup(g2)
	// Start Server
	e.Logger.Fatal(e.Start(":" + viper.GetString("PORT")))
}
