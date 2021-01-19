package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server represents the internal server structure and its deps
type Server struct {
	Echo     *echo.Echo
	Database *mongo.Database
}

// New will create a new server instance
func New(db *mongo.Database) *Server {
	return &Server{
		Echo:     echo.New(),
		Database: db,
	}
}

// Start will start the http server
func (s *Server) Start(port string) {
	// cors middleware
	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))

	// register the routes
	s.Routes()

	// start the server
	s.Echo.Logger.Fatal(s.Echo.Start(port))
}
