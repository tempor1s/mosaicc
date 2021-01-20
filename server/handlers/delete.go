package handlers

import (
	"context"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tempor1s/mosaic/auth"
	"github.com/tempor1s/mosaic/models"
)

// DeleteImage allows will delete an image on the users account based off of the
// object name (UploadedName)
func (h *Handlers) DeleteImage(c echo.Context) error {
	// get the user id from context (jwt)
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Println("User ID", userID)
	// get the object ID to delete
	objectID := html.EscapeString(c.Param("objectID"))
	log.Println("Object ID", objectID)

	// use the object ID to get the image from the database
	store := models.NewStore(h.Database)
	dbImage, err := store.GetImageByID(objectID)

	// if they do not have permissions, return an error
	if dbImage.UserID != userID {
		return echo.NewHTTPError(http.StatusForbidden, "you do not own that image")
	}

	// delete the image from GCP
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := h.StorageClient.Bucket(bucketName).Object(objectID)
	if err := o.Delete(ctx); err != nil {
		log.Println("did not delete from gcp. err:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete image")
	}

	log.Println("image deleted from GCP")

	err = store.DeleteImageByID(objectID)
	if err != nil {
		log.Println("did not delete from db. err:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete image")
	}

	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}
