package handlers

import (
	"fmt"
	"context"
	"time"
	"log"

  "github.com/google/uuid"

	"github.com/bulkashmak/gator-cli/internal"
	"github.com/bulkashmak/gator-cli/internal/commands"	
	"github.com/bulkashmak/gator-cli/internal/database"
)

func HandleLogin(s *internal.State, cmd commands.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("failed to get a user: %v", err)
	}

	err = s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}


func HandleRegister(s *internal.State, cmd commands.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

  user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username,
	})
	if err != nil {
		return fmt.Errorf("failed to create a user: %v", err)
	}

	s.Cfg.SetUser(user.Name)
	fmt.Printf("User '%s' registered successfully\n", user.Name)
	log.Println(user)

	return nil
}
