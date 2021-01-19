package handlers

import (
	"html"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tempor1s/mosaic/models"
)

// Images will return all the images that are on a users account
func (h *Handlers) Images(c echo.Context) error {
	// TODO: pull userID from context
	userID := "auth0|6004a8273225f90077cfe83a"

	store := models.NewStore(h.Database)
	images, err := store.GetImagesByUser(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "could not get images")
	}

	return c.JSON(http.StatusOK, images)
}

// Image will return the URL to a specific image given its short code
func (h *Handlers) Image(c echo.Context) error {
	id := html.EscapeString(c.Param("id"))

	return c.JSON(http.StatusOK, map[string]string{"image_id": id})
}
