package caption

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/aaronland/go-picturebook/bucket"
	pb_caption "github.com/aaronland/go-picturebook/caption"
	"github.com/tidwall/gjson"
)

func init() {

	ctx := context.Background()
	err := pb_caption.RegisterCaption(ctx, "flickr", NewFlickrCaption)

	if err != nil {
		panic(err)
	}
}

type FlickrCaption struct {
	pb_caption.Caption
}

func NewFlickrCaption(ctx context.Context, uri string) (pb_caption.Caption, error) {

	_, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	c := &FlickrCaption{}

	return c, nil
}

func (c *FlickrCaption) Text(ctx context.Context, source_bucket bucket.Bucket, path string) (string, error) {

	ext := filepath.Ext(path)

	img_ext := fmt.Sprintf("_o%s", ext)
	info_ext := "_i.json"

	info := strings.Replace(path, img_ext, info_ext, -1)

	/*
		exists, err := source_bucket.Exists(ctx, info)

		if err != nil {
			return "", err
		}

		if !exists {
			return "", errors.New("Missing _i.json file")
		}
	*/

	fh, err := source_bucket.NewReader(ctx, info, nil)

	if err != nil {
		return "", err
	}

	defer fh.Close()

	body, err := io.ReadAll(fh)

	var item interface{}
	err = json.Unmarshal(body, &item)

	if err != nil {
		return "", err
	}

	var rsp gjson.Result
	var photo_id int64
	var title string
	var taken string

	rsp = gjson.GetBytes(body, "photo.id")

	if !rsp.Exists() {
		return "", errors.New("Missing photo ID")
	}

	photo_id = rsp.Int()

	rsp = gjson.GetBytes(body, "photo.title._content")

	if !rsp.Exists() {
		return "", errors.New("Missing title")
	}

	title = rsp.String()

	rsp = gjson.GetBytes(body, "photo.dates.taken")

	if !rsp.Exists() {
		return "", errors.New("Missing date")
	}

	taken = rsp.String()

	// go... Y U SO WEIRD...
	// https://golang.org/src/time/format.go

	tm, err := time.Parse("2006-01-02 15:04:05", taken)

	if err != nil {
		return "", nil
	}

	dt := tm.Format("Jan 02, 2006")

	caption := fmt.Sprintf("<b>%s</b><br />%s / %d", title, dt, photo_id)
	return caption, nil
}
