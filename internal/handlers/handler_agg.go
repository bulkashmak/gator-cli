package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bulkashmak/gator-cli/internal"
	"github.com/bulkashmak/gator-cli/internal/commands"
	"github.com/bulkashmak/gator-cli/internal/database"
	"github.com/bulkashmak/gator-cli/internal/rss"
	"github.com/google/uuid"
)

func HandleAggregate(s *internal.State, cmd commands.Command) error {
  if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <fetch_interval>", cmd.Name)
	}

	fetchInternal, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid fetch interval: %w", err)
	}

	log.Printf("Fetching feed every %s...", fetchInternal)

	ticker := time.NewTicker(fetchInternal)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *internal.State) {
  feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feed to fetch", err)
		return
	}
	log.Println("Found a feed to fetch")
  scrapeFeed(s.DB, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
  _, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s as fetched: %v", feed.Name, err)
		return
	}

  feedData, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't fetch feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Items {
	  log.Printf("Found post: %s", item.Title)

		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err = db.CreatePost(context.TODO(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{
				String: item.Description,
			  Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("failed to create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %d posts found", feed.Name, len(feedData.Channel.Items))
}

