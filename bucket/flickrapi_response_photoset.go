package bucket

import (
	"iter"
)

type PhotoSetResponse struct {
	FlickrPhotosResponse
	PhotoSet *StandardPhotoSetResponse `json:"photoset"`
}

type StandardPhotoSetResponse struct {
	Id              string   `json:"id"`
	Title           string   `json:"title"`
	SortingOptionId string   `json:"sorting_option_id"`
	Primary         string   `json:"primary"`
	Owner           string   `json:"owner"`
	OwnerName       string   `json:"ownername"`
	Page            int      `json:"page"`
	Pages           int      `json:"pages"`
	PerPage         int      `json:"perpage"`
	Total           int      `json:"total"`
	Photo           []*Photo `json:"photo"`
}

func (r *PhotoSetResponse) Iterate() iter.Seq[*Photo] {

	return func(yield func(*Photo) bool) {

		for _, ph := range r.PhotoSet.Photo {
			if !yield(ph) {
				return
			}
		}

	}
}
