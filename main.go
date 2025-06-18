package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"

	"github.com/bulkashmak/gator-cli/internal"
	"github.com/bulkashmak/gator-cli/internal/commands"
	"github.com/bulkashmak/gator-cli/internal/config"
	"github.com/bulkashmak/gator-cli/internal/database"
	"github.com/bulkashmak/gator-cli/internal/handlers"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("failed to open db connection: %v", err)
	}
	dbQueries := database.New(db)

	appState := &internal.State{
		Cfg: &cfg,
		DB:  dbQueries,
	}

	cmds := commands.Commands{
		Store: make(map[string]func(*internal.State, commands.Command) error),
	}
	cmds.Register("login", handlers.HandleLogin)
	cmds.Register("register", handlers.HandleRegister)
	cmds.Register("users", handlers.HandleGetUsers)
	cmds.Register("reset", handlers.HandleDeleteUsers)
	cmds.Register("agg", handlers.HandleAggregate)
	cmds.Register("addfeed", middlewareLoggedIn(handlers.HandleAddFeed))
	cmds.Register("feeds", handlers.HandleFeeds)
	cmds.Register("follow", middlewareLoggedIn(handlers.HandleFollow))
	cmds.Register("following", middlewareLoggedIn(handlers.HandleFollowing))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.Run(appState, commands.Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}

func middlewareLoggedIn(handler func(s *internal.State, cmd commands.Command, user database.User) error) func(*internal.State, commands.Command) error {
	return func(s *internal.State, cmd commands.Command) error {
		currUserName := s.Cfg.CurrUserName
		if currUserName == "" {
			return errors.New("no logged in user in config")
		}

		currUser, err := s.DB.GetUser(context.Background(), currUserName)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}

		return handler(s, cmd, currUser)
	}
}
