package bucket

import (
	"context"
	"fmt"
	"io"
	"iter"
	"log/slog"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/aaronland/go-flickr-api/client"
	pb_bucket "github.com/aaronland/go-picturebook/bucket"
	"github.com/whosonfirst/go-ioutil"
)

type FlickrAPIBucket struct {
	pb_bucket.Bucket
	api_client client.Client
}

func init() {

	ctx := context.Background()
	err := pb_bucket.RegisterBucket(ctx, "flickrapi", NewFlickrAPIBucket)

	if err != nil {
		panic(err)
	}
}

func NewFlickrAPIBucket(ctx context.Context, uri string) (pb_bucket.Bucket, error) {

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

	b := &FlickrAPIBucket{
		api_client: api_client,
	}

	return b, nil
}

func (b *FlickrAPIBucket) GatherPictures(ctx context.Context, uris ...string) iter.Seq2[string, error] {

	return func(yield func(string, error) bool) {

		for _, uri := range uris {

			for path, err := range b.gatherPictures(ctx, uri) {
				if !yield(path, err) {
					return
				}
			}
		}
	}
}

func (b *FlickrAPIBucket) gatherPictures(ctx context.Context, uri string) iter.Seq2[string, error] {

	logger := slog.Default()
	logger = logger.With("uri", uri)

	ok_methods := []string{
		"flickr.galleries.getPhotos",
		"flickr.groups.pools.getPhotos",
		"flickr.photos.getContactsPhotos",
		"flickr.photos.getContactsPublicPhotos",
		"flickr.photos.getWithGeoData",
		"flickr.photos.getWithoutGeoData",
		"flickr.people.getPhotos",
		"flickr.photos.search",
		"flickr.photosets.getPhotos",
	}

	return func(yield func(string, error) bool) {

		args, err := url.ParseQuery(uri)

		if err != nil {
			yield("", err)
			return
		}

		if !slices.Contains(ok_methods, args.Get("method")) {
			yield("", fmt.Errorf("Unsupported method"))
			return
		}

		with_extras := []string{
			"url_n",
			"url_z",
			"url_c",
			"url_l",
			"url_o",
		}

		if args.Has("extras") {

			extras := strings.Split(args.Get("extras"), ",")

			for _, e := range with_extras {
				if !slices.Contains(extras, e) {
					extras = append(extras, e)
				}
			}

			args.Set("extras", strings.Join(extras, ","))

		} else {

			args.Set("extras", strings.Join(with_extras, ","))
		}

		cb := func(ctx context.Context, r io.ReadSeekCloser, err error) error {

			if err != nil {
				logger.Error("Callback yielded an error", "error", err)
				return err
			}

			rsp, err := NewFlickrPhotosResponse(args.Get("method"), r)

			if err != nil {
				yield("", err)
				return err
			}

			/*
				var rsp *PhotosSearchResponse

				dec := json.NewDecoder(r)
				err = dec.Decode(&rsp)

				if err != nil {
					yield("", err)
					return err
				}

				for _, ph := range rsp.Photos.Photo {
			*/

			for ph := range rsp.Iterate() {

				possible := []string{
					ph.UrlO,
					ph.UrlL,
					ph.UrlC,
					ph.UrlZ,
					ph.UrlN,
				}

				var ph_url string

				for _, u := range possible {
					if u != "" {
						ph_url = u
						break
					}
				}

				if !yield(ph_url, nil) {
					return fmt.Errorf("Yield did not return true")
				}
			}

			return nil
		}

		err = client.ExecuteMethodPaginatedWithClient(ctx, b.api_client, &args, cb)

		if err != nil {
			yield("", err)
			return
		}
	}
}

func (b *FlickrAPIBucket) NewReader(ctx context.Context, key string, opts any) (io.ReadSeekCloser, error) {

	// Validate key here...

	rsp, err := http.Get(key)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve %s, %w", key, err)
	}

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to retrieve %s, %d %s", key, rsp.StatusCode, rsp.Status)
	}

	return ioutil.NewReadSeekCloser(rsp.Body)
}

func (b *FlickrAPIBucket) NewWriter(ctx context.Context, key string, opts any) (io.WriteCloser, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (b *FlickrAPIBucket) Delete(ctx context.Context, key string) error {
	return fmt.Errorf("Not implemented")
}

func (b *FlickrAPIBucket) Attributes(ctx context.Context, key string) (*pb_bucket.Attributes, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (b *FlickrAPIBucket) Close() error {
	return nil
}
