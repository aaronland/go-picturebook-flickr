package bucket

import (
	"context"
	"flag"
	"fmt"
	"testing"
)

var client_uri = flag.String("flickr-client-uri", "", "...")

func TestFlickrAPIBucket(t *testing.T) {

	if *client_uri == "" {
		t.Skip()
	}

	ctx := context.Background()

	b, err := NewFlickrAPIBucket(ctx, *client_uri)

	if err != nil {
		t.Fatalf("Failed to create bucket, %v", err)
	}

	for ph, err := range b.GatherPictures(ctx) {

		if err != nil {
			t.Fatalf("Failed to gather pictures, %v", err)
		}

		fmt.Println(ph)
	}
}
