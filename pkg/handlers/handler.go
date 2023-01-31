package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/martin2844/link-shorten/pkg/middleware"
)

func RegisterBaseGroup(g *echo.Group) {
	g.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	g.GET("/all", getAllLinksHandler)

	g.GET("/:short", getLinkHandler)
}

func RegisterCreateGroup(g *echo.Group) {
	// Limits requests to 2 per minute per IP address.
	rateLimiter := middleware.NewCustomRateLimiter(2, time.Minute)
	g.Use(rateLimiter.Middleware)
	g.POST("", createLinkHandler)
}
