package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/bulkashmak/gator-cli/internal"
	"github.com/bulkashmak/gator-cli/internal/commands"
	"github.com/bulkashmak/gator-cli/internal/rss"
)

func HandleAggregate(s *internal.State, cmd commands.Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("command arguments are not allowed")
	}

	url := "https://www.wagslane.dev/index.xml"

	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	fmt.Printf("%v+\n", feed)

	return nil
}

