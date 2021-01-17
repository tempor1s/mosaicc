package handlers

import (
	"html"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Images will return all the images that are on a users account
func (h *Handlers) Images(c echo.Context) error {
	return nil
}

// Image will return the URL to a specific image given its short code
func (h *Handlers) Image(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))

	return c.JSON(http.StatusOK, map[string]string{"image_id": id})
}
