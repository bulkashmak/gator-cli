package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/bulkashmak/gator-cli/internal"
	"github.com/bulkashmak/gator-cli/internal/commands"
	"github.com/bulkashmak/gator-cli/internal/database"
	"github.com/google/uuid"
)

func HandleFollow(s *internal.State, cmd commands.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	currUserName := s.Cfg.CurrUserName

	currUser, err := s.DB.GetUser(context.Background(), currUserName)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	feed, err := s.DB.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
	}

  feedRow, err := s.DB.CreateFeedFollower(context.Background(), database.CreateFeedFollowerParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed_follower: %w", err)
	}

	fmt.Printf("User '%s' is now following feed '%s'\n", feedRow.UserName, feedRow.FeedName)

	return nil
}
