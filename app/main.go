package main

import (
	"app/server"
	"log"
	"net/http"
)

func main() {
	// Create a new instance of the server
	s := server.NewServer()

	// Start the server
	log.Fatal(http.ListenAndServe(":3456", s.Router))
}
