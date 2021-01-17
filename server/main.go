package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// hello function
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"msg": "hello, world"})
	})

	// TODO: add auth middleware here
	v1 := e.Group("/api/v1")

	// TODO: upload function
	v1.POST("/upload", func(c echo.Context) error {
		// TODO: return the URL of the uploaded image
		return c.JSON(http.StatusInternalServerError, "unimplemented")
	})

	// TODO: get all images for logged in account function
	v1.GET("/images", func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, "unimplemented")
	})

	// TODO: get single image by its name/id or something
	v1.GET("/image/:id", func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, "unimplemented")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
