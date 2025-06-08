package main

import (
	"log"
	"os"

	"github.com/bulkashmak/gator-cli/internal"
	"github.com/bulkashmak/gator-cli/internal/commands"
	"github.com/bulkashmak/gator-cli/internal/config"
	"github.com/bulkashmak/gator-cli/internal/handlers"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	appState := &internal.State{
		Cfg: &cfg,
	}

	cmds := commands.Commands{
		Store: make(map[string]func(*internal.State, commands.Command) error),
	}
  cmds.Register("login", handlers.HandleLogin)

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
