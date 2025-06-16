package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bulkashmak/gator-cli/internal"
	"github.com/bulkashmak/gator-cli/internal/commands"
	"github.com/bulkashmak/gator-cli/internal/database"
	"github.com/google/uuid"
)

func HandleAddFeed(s *internal.State, cmd commands.Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

  username := s.Cfg.CurrUserName
	user, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
    ID: uuid.New(),
    CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create a feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func HandleFeeds(s *internal.State, cmd commands.Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("command arguments are not allowed")
	}

	feeds, err := s.DB.ListFeedsWithUserNames(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds from db: %w", err)
	}

	for _, feed := range feeds {
    fmt.Printf("%s | %s | %s", feed.Name, feed.Url, feed.UserName)
	}
	
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
