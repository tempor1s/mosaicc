package main

import (
	"github.com/tempor1s/mosaic/server"
)

func main() {
	// create a new server and start it
	server := server.New()
	server.Start(":8080")
}
