package handlers

import (
	"net/http"

	"flickrphotosearch/backend/app/models"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
	"strconv"
)

type PhotoSearchResponse struct {
	Photos    []*models.Photo `jsonapi:"relation,photos"`
	PageCount int             `jsonapi:"attr,pageCount"`
	Total     string          `jsonapi:"attr,total"`
}

// PhotoSearchHandler /api/v1/photos/search/{query}/{page}
func PhotoSearchHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	var vars = mux.Vars(request)
	page, err := strconv.Atoi(vars["page"])

	if err != nil {
		panic(err)
	}

	photos, pageCount, total := models.SearchPhoto(vars["query"], page)

	response := &PhotoSearchResponse{
		Photos:    photos,
		PageCount: pageCount,
		Total:     total,
	}

	if err := jsonapi.MarshalPayload(w, response); err != nil {
		panic(err)
	}
}
