package handlers

import "go.mongodb.org/mongo-driver/mongo"

// Handlers is the structure that all handlers hang off of. This is so they are able to access dependencies like the DB
type Handlers struct {
	Database *mongo.Database
}

// New will return a new handlers structure
func New(db *mongo.Database) *Handlers {
	return &Handlers{
		Database: db,
	}
}
