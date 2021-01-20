package handlers

import (
	"cloud.google.com/go/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handlers is the structure that all handlers hang off of. This is so they are able to access dependencies like the DB
type Handlers struct {
	Database      *mongo.Database
	StorageClient *storage.Client
}

// New will return a new handlers structure
func New(db *mongo.Database, storage *storage.Client) *Handlers {
	return &Handlers{
		Database:      db,
		StorageClient: storage,
	}
}
