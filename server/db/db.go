package db

import (
	"context"
	"log"
	"time"

	"github.com/tempor1s/mosaic/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect will connect to the database and return a mongo database connection.
func Connect() *mongo.Database {
	// get db config from environment
	config := config.GetConfig()
	// mongodb credentials
	auth := options.Credential{
		Username: config.DbUsername,
		Password: config.DbPassword,
	}
	// connection uri for mongo
	uri := options.Client().ApplyURI(config.DbURI).SetAuth(auth)
	// contect for timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	// connect to the database
	client, err := mongo.Connect(ctx, uri)
	// fail if we cannot connect
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database.")

	// return the database connection connected to the mosaic database
	return client.Database("mosaic")
}
