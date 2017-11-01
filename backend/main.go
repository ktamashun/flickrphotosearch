package main

import (
	"log"
	"net/http"

	"flickrphotosearch/backend/app"
)

const (
	restPort = ":8080"
)

func main() {
	// JSONAPI
	router := app.CreateRouter()
	log.Fatal(http.ListenAndServe(restPort, router))
}
