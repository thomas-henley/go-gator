package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	command_func, exists := c.registeredCommands[cmd.Name]
	if exists {
		return command_func(s, cmd)
	} else {
		return errors.New("command not found")
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
