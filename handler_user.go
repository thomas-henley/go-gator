package main

import (
	"context"
	"fmt"
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
