package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/martin2844/link-shorten/pkg/links"
)

// Handlers are like controllers. They handle the request and return a response.
// The logic here could be further abstracted into a service layer.
func getLinkHandler(c echo.Context) error {
	short := c.Param("short")
	original, err := links.GetLink(short)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, original)
}

func getAllLinksHandler(c echo.Context) error {
	links, err := links.GetAllLinks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, links)
}
