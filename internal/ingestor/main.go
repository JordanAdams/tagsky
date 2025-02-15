package ingestor

import (
	"context"
	"fmt"

	"github.com/jordanadams/tagsky/internal/jetstream"
)

type IngestorOptions struct {
	Consumer jetstream.Consumer
}

func Start(ctx context.Context) error {
	js, err := jetstream.NewConsumer("wss://jetstream1.us-east.bsky.network/subscribe")
	if err != nil {
		return fmt.Errorf("failed to create jetstream consumer: %w", err)
	}

	js.Handler.HandleCommit(handleCommits)

	err = js.Start(ctx)
	if err != nil {
		return fmt.Errorf("failed to start jetstream consumer: %w", err)
	}

	return nil
}
