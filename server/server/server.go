package server

import (
	"context"
	"log"
	"net/http"

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
func New(db *mongo.Database, appEngine bool) *Server {
	// create new google cloud storage client for uploads/downloads of images
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("failed to create storage client: %v\n", err)
	}

	e := echo.New()
	// if we are in app engine, wire router to main handler
	if appEngine {
		http.Handle("/", e)
	}

	s := &Server{
		Echo:          e,
		Database:      db,
		StorageClient: client,
	}

	// register routes
	s.Routes()

	return s
}

// Start will start the http server
func (s *Server) Start(port string) {
	// cors middleware
	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))

	// clean stop the storage client when the server stops
	defer s.StorageClient.Close()

	// start the server
	s.Echo.Logger.Fatal(s.Echo.Start(port))
}
