package main

import (
	"log"
	"net/http"

	"github.com/ktamashun/flickrphotosearch/api-gateway/app"
)

const (
	restPort = ":18080"
)

func main() {
	// JSONAPI
	router := app.CreateRouter()
	log.Fatal(http.ListenAndServe(restPort, router))
}
