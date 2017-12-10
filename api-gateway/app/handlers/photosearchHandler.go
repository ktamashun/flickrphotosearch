package handlers

import (
	"net/http"

	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/ktamashun/flickrphotosearch/api-gateway/app/models"
	"github.com/ktamashun/flickrphotosearch/api-gateway/apis"
	"context"
	"google.golang.org/grpc"
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

	photoSearchRequest := &apis.PhotoSearchRequest{
		Page: int32(page),
		Query: vars["query"],
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("flickrphotosearch-backend:18000", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	photoClient := apis.NewPhotoClient(conn)
	photoSearchResponse, err := photoClient.Search(context.Background(), photoSearchRequest)
	if err != nil {
		panic(err)
	}

	photos := make([]*models.Photo, len(photoSearchResponse.Photos))

	for key, photo := range photoSearchResponse.Photos {
		photos[key] = &models.Photo{
			ID: photo.GetID(),
			ThumbnailUrl: photo.GetThumbnailUrl(),
			MediumUrl: photo.GetMediumUrl(),
			LargeUrl: photo.GetLargeUrl(),
			Owner: photo.GetOwner(),
			Title: photo.Title,
		}
	}

	response := &PhotoSearchResponse{
		Photos:    photos,
		PageCount: int(photoSearchResponse.PageCount),
		Total:     photoSearchResponse.Total,
	}

	if err := jsonapi.MarshalPayload(w, response); err != nil {
		panic(err)
	}
}
