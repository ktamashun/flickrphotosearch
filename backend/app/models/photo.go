package models

import (
	"crypto/tls"
	"encoding/json"
	"flickrphotosearch/backend/app/common"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/bradfitz/gomemcache/memcache"
	"time"
	"log"
)

// Photo is a photo from Flickr
type Photo struct {
	ID           string `jsonapi:"primary,photo"`
	ThumbnailUrl string `jsonapi:"attr,thumbnailUrl"`
	MediumUrl    string `jsonapi:"attr,mediumUrl"`
	LargeUrl     string `jsonapi:"attr,largeUrl"`
	Owner        string `jsonapi:"attr,owner"`
	Title        string `jsonapi:"attr,title"`
}

type PhotoSearchRawPhoto struct {
	Farm     int    `json:"farm"`
	Id       string `json:"id"`
	Isfamily int    `json:"isfamily"`
	Isfriend int    `json:"isfriend"`
	Ispublic int    `json:"ispublic"`
	Owner    string `json:"owner"`
	Secret   string `json:"secret"`
	Server   string `json:"server"`
	Title    string `json:"title"`
}

type PhotoSearchBlock struct {
	Page    int                    `json:"page"`
	Pages   int                    `json:"pages"`
	Perpage int                    `json:"perpage"`
	Total   string                 `json:"total"`
	Photo   []*PhotoSearchRawPhoto `json:"photo"`
}

type PhotoSearchResponse struct {
	Stat   string           `json:"stat"`
	Photos PhotoSearchBlock `json:"photos"`
}

const FlickrApiUrl = "https://api.flickr.com/services/rest/?method=flickr.photos.search&format=json&api_key=%s&per_page=%d&text=%s&page=%d&nojsoncallback=1"
const FlickrPhotoUrl = "https://farm%d.staticflickr.com/%s/%s_%s_%s.jpg"

func SearchPhoto(query string, page int) ([]*Photo, int, string) {
	photos := []*Photo{}
	apiUrl := fmt.Sprintf(FlickrApiUrl, common.AppConfig.FlickrApiKey, 100, url.QueryEscape(query), page)
	body := []byte{}

	cachedData, err := common.Cache.Get(apiUrl)

	log.Println("Requested URL: ", apiUrl)

	if err != nil {
		log.Println("Cache miss, making request to the Flickr API")

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		res, err := client.Get(apiUrl)
		if err != nil {
			panic(err.Error())
		}

		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}

		common.Cache.Set(&memcache.Item{Key: apiUrl, Value: body, Expiration: int32(time.Minute.Seconds())})
	} else {
		body = cachedData.Value
		log.Println("Loading data from cache")
	}

	parsedResponse, err := ParsePhotoSearchResponse(body)
	if err != nil {
		panic(err)
	}

	for _, rawphoto := range parsedResponse.Photos.Photo {
		photos = append(photos, ConvertRawPhoto(rawphoto))
	}

	return photos, parsedResponse.Photos.Pages, parsedResponse.Photos.Total
}

func ParsePhotoSearchResponse(body []byte) (*PhotoSearchResponse, error) {
	var s = new(PhotoSearchResponse)
	err := json.Unmarshal(body, &s)

	if err != nil {
		fmt.Println("whoops:", err)
	}

	return s, err
}

func ConvertRawPhoto(rawphoto *PhotoSearchRawPhoto) *Photo {
	return &Photo{
		ID:           rawphoto.Id,
		ThumbnailUrl: fmt.Sprintf(FlickrPhotoUrl, rawphoto.Farm, rawphoto.Server, rawphoto.Id, rawphoto.Secret, "m"),
		MediumUrl:    fmt.Sprintf(FlickrPhotoUrl, rawphoto.Farm, rawphoto.Server, rawphoto.Id, rawphoto.Secret, "b"),
		LargeUrl:     fmt.Sprintf(FlickrPhotoUrl, rawphoto.Farm, rawphoto.Server, rawphoto.Id, rawphoto.Secret, "h"),
		Owner:        rawphoto.Owner,
		Title:        rawphoto.Title,
	}
}
