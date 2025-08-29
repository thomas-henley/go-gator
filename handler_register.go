package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thomas-henley/go-gator/internal/database"
)

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
