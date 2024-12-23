package caption

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"testing"

	"github.com/aaronland/go-picturebook-flickr/bucket"
)

var client_uri = flag.String("flickr-client-uri", "", "...")
var photo_uri = flag.String("flickr-photo-uri", "", "...")

func TestFlickrAPICaption(t *testing.T) {

	if *client_uri == "" {
		t.Skip()
	}

	if *photo_uri == "" {
		t.Fatalf("Missing photo-uri flag")
	}

	ctx := context.Background()

	bucket_q := url.Values{}
	bucket_q.Set("client_uri", *client_uri)

	bucket_u := url.URL{}
	bucket_u.Scheme = "flickrapi"
	bucket_u.RawQuery = bucket_q.Encode()

	b, err := bucket.NewFlickrAPIBucket(ctx, bucket_u.String())

	if err != nil {
		t.Fatalf("Failed to create bucket, %v", err)
	}

	caption_q := url.Values{}
	caption_q.Set("client_uri", *client_uri)

	caption_u := url.URL{}
	caption_u.Scheme = "flickrapi"
	caption_u.RawQuery = caption_q.Encode()

	c, err := NewFlickrAPICaption(ctx, caption_u.String())

	if err != nil {
		t.Fatalf("Failed to create caption, %v", err)
	}

	c_text, err := c.Text(ctx, b, *photo_uri)

	if err != nil {
		t.Fatalf("Failed to derive caption text, %v", err)
	}

	fmt.Println(c_text)
}
