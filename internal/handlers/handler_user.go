package handlers

import (
	"context"
	"fmt"
	"log"
	"time"
	"errors"

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

func HandleGetUsers(s *internal.State, cmd commands.Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("arguments are not allowed")
	}

	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get all users: %v", err)
	}

	for _, user := range users {
		if s.Cfg.CurrUserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}

func HandleDeleteUsers(s *internal.State, cmd commands.Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("arguments are not allowed")
	}

	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete all users: %v", err)
	}

	return nil
}

