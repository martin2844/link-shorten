package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lynxsecurity/viper"
	"github.com/martin2844/link-shorten/pkg/db"
	"github.com/martin2844/link-shorten/pkg/handlers"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	fmt.Println(viper.Get("PORT"))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! test")
	})

	db.Init()

	//Gets first link
	var link db.Link
	db.Instance.First(&link, 1)

	//Print the first entry
	fmt.Println(link)

	// base group
	g := e.Group("")
	g2 := e.Group("/create")
	handlers.RegisterBaseGroup(g)
	handlers.RegisterCreateGroup(g2)

	e.Logger.Fatal(e.Start(":" + viper.GetString("PORT")))
}
