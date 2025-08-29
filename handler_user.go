package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thomas-henley/go-gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	if _, err := s.db.GetUser(context.Background(), name); err != nil {
		return fmt.Errorf("error retrieving user from db: %w", err)
	}

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("logged in as user: %s\n", name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
    if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	currentTime := time.Now().UTC()

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name: name,
		})
	if err != nil {
        return fmt.Errorf("error creating user: %w", err)
	}

	s.cfg.SetUser(name)

	fmt.Printf("User \"%s\" has been created\n", user.Name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting the user database: %w", err)
	}

	fmt.Printf("database has been reset\n")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s\n", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching users: %w\n", err)
	}

    for _, user := range(users) {
		loginStatus := ""
		if (user.Name == s.cfg.CurrentUserName) {
			loginStatus = " (current)"
		}
		fmt.Printf("* %s%s\n", user.Name, loginStatus)
		
	}

	return nil
}
