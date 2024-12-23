package bucket

import (
	"iter"
)

type PhotosSearchResponse struct {
	FlickrPhotosResponse
	Photos *StandardPhotosResponse `json:"photos"`
}

type StandardPhotosResponse struct {
	// For reasons I don't understand yet this doesn't jibe with
	// what the Flickr API actually returns...
	// Page    int      `json:"page"`
	// Pages   int      `json:"pages"`
	// PerPage int      `json:"perpage"`
	// Total   int      `json:"total"`
	Photo []*Photo `json:"photo"`
}

func (r *PhotosSearchResponse) Iterate() iter.Seq[*Photo] {

	return func(yield func(*Photo) bool) {

		for _, ph := range r.Photos.Photo {
			if !yield(ph) {
				return
			}
		}

	}
}
