package server

import (
	"github.com/labstack/echo/v4"
	"github.com/tempor1s/mosaic/handlers"
)

// Routes will register all of the routes on the server structure
func (s *Server) Routes() {
	h := handlers.New(s.Database)
	// hello route for status checks
	s.Echo.GET("/status", h.Hello)
	// setup api group
	v1 := s.Echo.Group("/api/v1")
	// enable auth middleware (all routes after this point will require authorization)
	mw := getJwtMiddleware()
	v1.Use(echo.WrapMiddleware(mw.Handler))

	// upload image to the server (gets bound to logged in account)
	v1.POST("/upload", h.Upload)
	// get all images for logged in account
	v1.GET("/images", h.Images)
}
