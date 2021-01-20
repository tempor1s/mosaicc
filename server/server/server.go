package server

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"

	"cloud.google.com/go/storage"
)

// Server represents the internal server structure and its deps
type Server struct {
	Echo          *echo.Echo
	Database      *mongo.Database
	StorageClient *storage.Client
}

// New will create a new server instance and its required dependencies
func New(db *mongo.Database) *Server {
	// create new google cloud storage client for uploads/downloads of images
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("failed to create storage client: %v\n", err)
	}

	return &Server{
		Echo:          echo.New(),
		Database:      db,
		StorageClient: client,
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

	// clean stop the storage client when the server stops
	defer s.StorageClient.Close()

	// start the server
	s.Echo.Logger.Fatal(s.Echo.Start(port))
}
