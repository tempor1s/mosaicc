// +build !appengine !appenginevm

package main

import (
	"github.com/tempor1s/mosaic/db"
	"github.com/tempor1s/mosaic/server"
)

func main() {
	// connect to db
	database := db.Connect()
	// create a new server and start it
	server := server.New(database, false)
	server.Start(":8080")
}
