package bucket

import (
	"context"
	"flag"
	"fmt"
	"testing"
	"net/url"
)

var client_uri = flag.String("flickr-client-uri", "", "...")
var query_uri = flag.String("flickr-query-uri", "", "...")

func TestFlickrAPIBucket(t *testing.T) {

	if *client_uri == "" {
		t.Skip()
	}

	if *query_uri == "" {
		t.Fatalf("Missing query-uri flag")
	}

	ctx := context.Background()

	bucket_q := url.Values{}
	bucket_q.Set("client_uri", *client_uri)

	bucket_u := url.URL{}
	bucket_u.Scheme = "flickrapi"
	bucket_u.RawQuery = bucket_q.Encode()
	
	b, err := NewFlickrAPIBucket(ctx, bucket_u.String())

	if err != nil {
		t.Fatalf("Failed to create bucket, %v", err)
	}

	for ph, err := range b.GatherPictures(ctx, *query_uri) {

		if err != nil {
			t.Fatalf("Failed to gather pictures, %v", err)
		}

		fmt.Println(ph)
	}
}
