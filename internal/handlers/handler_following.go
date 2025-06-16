package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/bulkashmak/gator-cli/internal"
	"github.com/bulkashmak/gator-cli/internal/commands"
)

func HandleFollowing(s *internal.State, cmd commands.Command) error {
  if len(cmd.Args) != 0 {
		return errors.New("command arguments are not allowed")
	}

  currUserName := s.Cfg.CurrUserName
	currUser, err := s.DB.GetUser(context.Background(), currUserName)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	feedFollows, err := s.DB.GetFeedFollowsForUser(context.Background(), currUser.ID)
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
	}

  for _, follow := range feedFollows {
		fmt.Printf("- %s\n", follow.FeedName)
	}
 
	return nil
}
