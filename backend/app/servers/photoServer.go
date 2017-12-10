package servers

import (
	"context"
	"github.com/ktamashun/flickrphotosearch/backend/apis"
	"github.com/ktamashun/flickrphotosearch/backend/app/models/photo_search"
)

type PhotoServer struct{}

func (s *PhotoServer) Search(ctx context.Context, request *apis.PhotoSearchRequest) (*apis.PhotoSearchResponse, error) {

	photos, pageCount, total := photo_search.SearchPhoto(request.GetQuery(), int(request.GetPage()))
	responsePhotos := make([]*apis.PhotoSearchResponse_Photo, len(photos))

	for key, photo := range photos {
		responsePhotos[key] = &apis.PhotoSearchResponse_Photo{
			ID:           photo.ID,
			ThumbnailUrl: photo.ThumbnailUrl,
			MediumUrl:    photo.MediumUrl,
			Owner:        photo.Owner,
			Title:        photo.Title,
		}
	}

	response := &apis.PhotoSearchResponse{
		PageCount: int32(pageCount),
		Total:     total,
		Photos:    responsePhotos,
	}

	return response, nil
}
