package main

import (
	"context"
	_ "github.com/aaronland/go-picturebook-flickr/caption"
	_ "github.com/aaronland/go-picturebook-flickr/filter"
	"github.com/aaronland/go-picturebook/application/commandline"
	_ "gocloud.dev/blob/fileblob"
	"log"
	"os"
)

func main() {

	ctx := context.Background()

	fs, err := commandline.DefaultFlagSet(ctx)

	if err != nil {
		log.Fatal(err)
	}

	app, err := commandline.NewApplication(ctx, fs)

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
