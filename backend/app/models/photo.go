package models

// Photo is a photo from Flickr
type Photo struct {
	ID           string `jsonapi:"primary,photo"`
	ThumbnailUrl string `jsonapi:"attr,thumbnailUrl"`
	MediumUrl    string `jsonapi:"attr,mediumUrl"`
	LargeUrl     string `jsonapi:"attr,largeUrl"`
	Owner        string `jsonapi:"attr,owner"`
	Title        string `jsonapi:"attr,title"`
}
