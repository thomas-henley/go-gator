package main

import (
	"context"
	"fmt"
)

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
