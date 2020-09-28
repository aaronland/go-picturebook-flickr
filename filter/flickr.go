package filter

import (
	"context"
	pb_filter "github.com/aaronland/go-picturebook/filter"
	"gocloud.dev/blob"
	"net/url"
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

func (f *FlickrFilter) Continue(ctx context.Context, bucket *blob.Bucket, path string) (bool, error) {

	if !flickr_re.MatchString(path) {
		return false, nil
	}

	return true, nil
}
