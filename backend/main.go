package main

import (
	"log"
	"net/http"

	"github.com/ktamashun/flickrphotosearch/backend/app"
)

const (
	restPort = ":8080"
)

func main() {
	// JSONAPI
	router := app.CreateRouter()
	log.Fatal(http.ListenAndServe(restPort, router))
}
