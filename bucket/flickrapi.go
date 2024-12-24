package bucket

/*

./bin/picturebook \
	-source-uri 'flickrapi://?client_uri=oauth1%3A%2F%2F%3Fconsumer_key%3D...%26consumer_secret%3D...%26oauth_token%3D...%26oauth_token_secret%3D...'
	-target-uri /usr/local/src/go-picturebook-flickr \
	'method=flickr.photos.search&user_id=35034348999@N01&tags=flickrhq'

./bin/picturebook \
	-source-uri 'flickrapi://?client_uri=oauth1%3A%2F%2F%3Fconsumer_key%3D...%26consumer_secret%3D...%26oauth_token%3D...%26oauth_token_secret%3D...'
	-target-uri /usr/local/src/go-picturebook-flickr \
	-filename ca.pdf \
	'method=flickr.photosets.getPhotos&user_id=35034348999@N01&photoset_id=72177720319945125'

*/

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

// FlickrAPIBucket implements the `aaronland/go-picturebook/bucket.Bucket` interface for photos returned by the Flickr API.
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

// AllowedFlickrAPIMethods returns the list of Flickr API methods supported by this package.
func AllowedFlickrAPIMethods() []string {

	ok_methods := []string{
		"flickr.favorites.getList",
		"flickr.galleries.getPhotos",
		"flickr.groups.pools.getPhotos",
		"flickr.photos.getContactsPhotos",
		"flickr.photos.getContactsPublicPhotos",
		"flickr.photos.getWithGeoData",
		"flickr.photos.getWithoutGeoData",
		"flickr.people.getPhotos",
		"flickr.people.getPhotosOf",
		"flickr.photos.search",
		"flickr.photosets.getPhotos",
	}

	return ok_methods
}

// NewFlickrAPIBucket returns a new `FlickrAPIBucket` instance implementing the `aaronland/go-picturebook/bucket.Bucket` interface for use with photos returned by the Flickr API.
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

// GatherPictures returns a new `iter.Seq2[string, error]` instance containing the URIs for a collection of Flickr API results matching the queries in 'uris'.
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

	ok_methods := AllowedFlickrAPIMethods()

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
			yield("", fmt.Errorf("Failed to execute paginated method, %w", err))
			return
		}
	}
}

// NewReader returns a new `io.ReadSeekCloser` instance for an Flickr photo identified by 'key'.
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

// NewWriter returns an error because this package only implements non-destructive methods of the `aaronland/go-picturebook/bucket.Bucket` interface.
func (b *FlickrAPIBucket) NewWriter(ctx context.Context, key string, opts any) (io.WriteCloser, error) {
	return nil, fmt.Errorf("Not implemented")
}

// Delete returns an error because this package only implements non-destructive methods of the `aaronland/go-picturebook/bucket.Bucket` interface.
func (b *FlickrAPIBucket) Delete(ctx context.Context, key string) error {
	return fmt.Errorf("Not implemented")
}

// Attribute returns a new `aaronland/go-picturebook/bucket.Attributes` instance for an object image identified by 'key'.
func (b *FlickrAPIBucket) Attributes(ctx context.Context, key string) (*pb_bucket.Attributes, error) {
	return nil, fmt.Errorf("Not implemented")
}

// Close completes and terminates any underlying code used by 'b'.
func (b *FlickrAPIBucket) Close() error {
	return nil
}
