package main

import (
	"database/sql"
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
	cmds.Register("addfeed", handlers.HandleAddFeed)
	cmds.Register("feeds", handlers.HandleFeeds)
	cmds.Register("follow", handlers.HandleFollow)
	cmds.Register("following", handlers.HandleFollowing)

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
