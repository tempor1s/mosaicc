package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Routes will register all of the routes on the server structure
func (s *Server) Routes() {
	// hello function
	s.Echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"msg": "hello, world"})
	})

	// setup api group
	v1 := s.Echo.Group("/api/v1")
	// enable auth middleware
	mw := getJwtMiddleware()
	v1.Use(echo.WrapMiddleware(mw.Handler))

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
}
