package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/bulkashmak/gator-cli/internal"
	"github.com/bulkashmak/gator-cli/internal/commands"
	"github.com/bulkashmak/gator-cli/internal/database"
)

func HandleFollowing(s *internal.State, cmd commands.Command, user database.User) error {
  if len(cmd.Args) != 0 {
		return errors.New("command arguments are not allowed")
	}

	feedFollows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
	}

  for _, follow := range feedFollows {
		fmt.Printf("- %s\n", follow.FeedName)
	}
 
	return nil
}
