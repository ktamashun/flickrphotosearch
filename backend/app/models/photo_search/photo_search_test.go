package photo_search

import (
	"github.com/ktamashun/flickrphotosearch/backend/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertApiPhoto(t *testing.T) {
	apiPhoto := ApiPhoto{
		Farm:     5,
		Id:       "123456",
		Isfamily: 1,
		Isfriend: 1,
		Ispublic: 1,
		Owner:    "testUser",
		Secret:   "xyzabc",
		Server:   "12345",
		Title:    "Test Photo",
	}

	expectedPhoto := models.Photo{
		ID:           "123456",
		ThumbnailUrl: "https://farm5.staticflickr.com/12345/123456_xyzabc_m.jpg",
		MediumUrl:    "https://farm5.staticflickr.com/12345/123456_xyzabc_b.jpg",
		LargeUrl:     "https://farm5.staticflickr.com/12345/123456_xyzabc_h.jpg",
		Owner:        "testUser",
		Title:        "Test Photo",
	}

	assert.Equal(t, &expectedPhoto, ConvertApiPhoto(&apiPhoto))
}
