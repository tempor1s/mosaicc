package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tempor1s/mosaic/auth"
	"github.com/tempor1s/mosaic/file"
	"github.com/tempor1s/mosaic/models"
)

const (
	bucketName = "mosaic-image-bucket"
)

// Upload will upload a file to the server and return the URL to the image
func (h *Handlers) Upload(c echo.Context) error {
	// get the user id from context (jwt)
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// get the image from the posted form
	image, err := c.FormFile("image")
	if err != nil {
		log.Println("failed to get image from form. err:", err)
		return echo.NewHTTPError(http.StatusBadRequest, "please provide a valid image file")
	}

	// open the file
	openedFile, err := image.Open()
	if err != nil {
		log.Println("something went wrong when opening the image. err:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error reading file")
	}

	// create in memory version of the file
	file := file.NewFile(openedFile, image)
	// give the file a nice name for uploading (length of 7 for now)
	file.GenerateName(7)

	// upload context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// read the file to get the bytes to upload

	// upload the image object with storage.Writer
	wc := h.StorageClient.Bucket(bucketName).Object(file.Fullname).NewWriter(ctx)
	if _, err := io.Copy(wc, openedFile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to copy file internally")
	}
	if err := wc.Close(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to copy file internally. close error")
	}

	log.Println("file uploaded")

	// craft the URLs to store in the database
	fullURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, file.Fullname)
	shortURL := fmt.Sprintf("https://i.benl.dev/%s", file.Fullname)

	// image model to store in the database
	newImg := models.Image{
		FullURL:      fullURL,
		ShortURL:     shortURL,
		UserID:       userID,         // set the user ID in the DB to be that from auth
		Name:         file.Basename,  // the new name of the file (generated)
		UploadedName: image.Filename, // the original name of the file
		UploadDate:   time.Now(),     // when the file was uploaded (for sorting)
	}

	store := models.NewStore(h.Database)
	err = store.CreateImage(newImg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create image in the database")
	}

	return c.JSON(http.StatusOK, newImg)
}
