package bucket

import (
	"context"
	"flag"
	"fmt"
	"testing"
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

	b, err := NewFlickrAPIBucket(ctx, *client_uri)

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
