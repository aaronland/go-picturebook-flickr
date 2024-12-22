package caption

import (
	"context"
	"fmt"
	"net/url"

	"github.com/aaronland/go-flickr-api/client"
	pb_caption "github.com/aaronland/go-picturebook/caption"
)

type FlickrAPICaption struct {
	pb_caption.Caption
	api_client client.Client
}

func init() {

	ctx := context.Background()
	err := pb_caption.RegisterCaption(ctx, "flickrapi", NewFlickrAPICaption)

	if err != nil {
		panic(err)
	}
}

func NewFlickrAPICaption(ctx context.Context, uri string) (pb_caption.Caption, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()
	client_uri := q.Get("client_uri")

	api_client, err := client.NewClient(ctx, client_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new client, %w", err)
	}

	b := &FlickrAPICaption{
		api_client: api_client,
	}

	return b, nil
}
