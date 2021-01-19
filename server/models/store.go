package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Store is a way to interact with the database with various methods
type Store struct {
	Database *mongo.Database
}

// NewStore will return a new store with the given database for interaction (also allows us to add easy caching later)
func NewStore(db *mongo.Database) *Store {
	return &Store{
		Database: db,
	}
}
