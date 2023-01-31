package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/martin2844/link-shorten/pkg/links"
)

type (
	Link struct {
		Original string `json:"original"`
	}
)

func createLinkHandler(c echo.Context) error {
	u := new(Link)
	c.Bind(u)
	//Move to validate function --> Todo //TODO
	if u.Original == "" {
		return c.JSON(http.StatusBadRequest, "Invalid URL")
	}
	short, err := links.CreateLink(u.Original)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, short)
}
