package main

import (
	"context"
	"log"

	_ "github.com/aaronland/go-picturebook-flickr/bucket"
	_ "github.com/aaronland/go-picturebook-flickr/caption"
	_ "github.com/aaronland/go-picturebook-flickr/filter"
	_ "gocloud.dev/blob/fileblob"

	"github.com/aaronland/go-picturebook/app/picturebook"
)

func main() {

	ctx := context.Background()
	err := picturebook.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
