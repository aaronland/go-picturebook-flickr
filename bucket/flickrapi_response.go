package bucket

import (
	"encoding/json"
	"io"
	"iter"
	"log/slog"
)

type FlickrPhotosResponse interface {
	Iterate() iter.Seq[*Photo]
}

type Photo struct {
	Id       string `json:"id"`
	Owner    string `json:"owner"`
	Secret   string `json:"secret"`
	Farm     int    `json:"farm"`
	Title    string `json:"title"`
	IsPublic int    `json:"ispublic"`
	IsFriend int    `json:"isfriend"`
	IsFamily int    `json:"isfamily"`
	UrlN     string `json:"url_n"`
	UrlZ     string `json:"url_z"`
	UrlC     string `json:"url_c"`
	UrlL     string `json:"url_l"`
	UrlO     string `json:"url_o"`
}

func NewFlickrPhotosResponse(method string, r io.ReadSeeker) (FlickrPhotosResponse, error) {

	switch method {
	case "flickr.photosets.getPhotos":

		var rsp *PhotoSetResponse

		dec := json.NewDecoder(r)
		err := dec.Decode(&rsp)

		if err != nil {
			slog.Error("Failed to decode response for method", "method", method, "error", err)
			return nil, err
		}

		return rsp, nil

	default:

		var rsp *PhotosSearchResponse

		dec := json.NewDecoder(r)
		err := dec.Decode(&rsp)

		if err != nil {
			slog.Error("Failed to decode response for method", "method", method, "error", err)
			return nil, err
		}

		return rsp, nil
	}
}
