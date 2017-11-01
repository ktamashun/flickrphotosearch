package app

import (
	"flickrphotosearch/backend/app/handlers"
	"net/http"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes of Route items
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1/photos/search/{query}/{page}",
		handlers.PhotoSearchHandler,
	},
}
