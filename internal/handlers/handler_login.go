package handlers

import (
	"fmt"

	"github.com/bulkashmak/gator-cli/internal/commands"
	"github.com/bulkashmak/gator-cli/internal"
)

func HandleLogin(s *internal.State, cmd commands.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	err := s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

