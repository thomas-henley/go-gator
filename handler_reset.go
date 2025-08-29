package main

import (
	"context"
	"fmt"
)

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
