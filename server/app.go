// +build appengine

// for when we are running on app engine (deployed)
package main

import (
	"github.com/tempor1s/mosaic/db"
	"github.com/tempor1s/mosaic/server"
)

func main() {
	// connect to db
	database := db.Connect()
	// create a new server and register handlers for app engine
	server.New(database, true)
}
