package filter

import (
	"context"
	"net/url"
	"path/filepath"

	pb_filter "github.com/aaronland/go-picturebook/filter"
	"github.com/aaronland/go-picturebook/bucket"	
)

func init() {

	ctx := context.Background()
	err := pb_filter.RegisterFilter(ctx, "flickr", NewFlickrFilter)

	if err != nil {
		panic(err)
	}
}

type FlickrFilter struct {
	pb_filter.Filter
}

func NewFlickrFilter(ctx context.Context, uri string) (pb_filter.Filter, error) {

	_, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	f := &FlickrFilter{}

	return f, nil
}

func (f *FlickrFilter) Continue(ctx context.Context, source_bucket bucket.Bucket, path string) (bool, error) {

	fname := filepath.Base(path)

	if !flickr_re.MatchString(fname) {
		return false, nil
	}

	return true, nil
}
