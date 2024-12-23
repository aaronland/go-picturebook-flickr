package caption

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/aaronland/go-flickr-api/client"
	pb_bucket "github.com/aaronland/go-picturebook/bucket"
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

func (c *FlickrAPICaption) Text(ctx context.Context, b pb_bucket.Bucket, key string) (string, error) {

	base := filepath.Base(key)

	// Please use a regexp...
	fname := strings.Split(base, "_")
	photo_id := fname[0]

	args := &url.Values{}
	args.Set("method", "flickr.photos.getInfo")
	args.Set("photo_id", photo_id)

	r, err := c.api_client.ExecuteMethod(ctx, args)

	if err != nil {
		return "", err
	}

	defer r.Close()

	var rsp *GetInfoResponse

	dec := json.NewDecoder(r)
	err = dec.Decode(&rsp)

	parts := make([]string, 0)

	if rsp.Photo.Title.Content == "" {
		parts = append(parts, "(untitled)")
	} else {
		parts = append(parts, fmt.Sprintf(`"%s"`, rsp.Photo.Title.Content))
	}

	parts = append(parts, fmt.Sprintf("%s (%s)", rsp.Photo.Owner.RealName, rsp.Photo.Owner.UserName))
	parts = append(parts, strings.Split(rsp.Photo.Dates.Taken, " ")[0])

	return strings.Join(parts, "\n"), nil
}
