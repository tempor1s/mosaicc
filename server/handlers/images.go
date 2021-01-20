package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tempor1s/mosaic/auth"
	"github.com/tempor1s/mosaic/models"
)

// Images will return all the images that are on a users account
func (h *Handlers) Images(c echo.Context) error {
	// get the user id from context (jwt)
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// get the images for a user from the database based on the ID pulled from context
	store := models.NewStore(h.Database)
	images, err := store.GetImagesByUser(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "could not get images")
	}

	return c.JSON(http.StatusOK, images)
}
