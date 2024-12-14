package main

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/aaronland/go-picturebook-flickr/caption"
	_ "github.com/aaronland/go-picturebook-flickr/filter"
	"github.com/aaronland/go-picturebook/application/commandline"
	_ "gocloud.dev/blob/fileblob"
)

func main() {

	ctx := context.Background()
	logger := slog.Default()

	err := commandline.Run(ctx, logger)

	if err != nil {
		logger.Error("Failed to run picturebook application", "error", err)
		os.Exit(1)
	}

	os.Exit(0)
}
