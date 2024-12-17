package main

import (
	"context"
	"log"

	_ "github.com/aaronland/go-picturebook-flickr/caption"
	_ "github.com/aaronland/go-picturebook-flickr/filter"
	"github.com/aaronland/go-picturebook/app/picturebook"
	_ "gocloud.dev/blob/fileblob"
)

func main() {

	ctx := context.Background()
	err := picturebook.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
