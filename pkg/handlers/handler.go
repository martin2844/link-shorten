package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/martin2844/link-shorten/pkg/middleware"
)

// Registers the base group of endpoints, i.e. server url + /:short, /all, etc.
func RegisterBaseGroup(g *echo.Group) {
	g.GET("/all", getAllLinksHandler)
	g.GET("/:short", getLinkHandler)
}

// Registers the create group of endpoints, i.e. {server url}/create + "/one", "/many", etc.
func RegisterCreateGroup(g *echo.Group) {
	// Limits requests to 2 per minute per IP address.
	rateLimiter := middleware.NewCustomRateLimiter(2, time.Minute)
	g.Use(rateLimiter.Middleware)
	g.POST("/one", createLinkHandler)
	// TODO create /many endpoint max 10 links at time
}
