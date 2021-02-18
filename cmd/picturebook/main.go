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
		log.Fatalf("Failed to create default flag set, %v", err)
	}

	// In anticipation of 0.5.0 release
	
	/*

	err = commandline.AppendDeprecatedFlags(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to append deprecated flags, %v", err)
	}

	*/
	
	app, err := commandline.NewApplication(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to create new picturebook application, %v", err)
	}

	err = app.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run picturebook application, %v", err)
	}

	os.Exit(0)
}
