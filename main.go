package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/thomas-henley/go-gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	names map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	command_func, exists := c.names[cmd.name]
	if exists {
		return command_func(s, cmd)
	} else {
		return errors.New("command does not exist")
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.names[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login requires a username argument")
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("logged in as user: %s\n", cmd.args[0])
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

	s := state{config: &cfg}

	cmds := commands{}
	cmds.names = map[string]func(*state, command) error{}
	cmds.register("login", handlerLogin)

	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Printf("command argument missing\n")
		os.Exit(1)
	}

	cmd := command{
		name: arguments[1],
		args: arguments[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// cfg, err := config.Read()
	// if err != nil {
	// 	log.Fatalf("error reading config: %v", err)
	// }
	// fmt.Printf("Read config: %+v\n", cfg)
	//
	// err = cfg.SetUser("thomas")
	// if err != nil {
	// 	log.Fatalf("couldn't set current user: %v", err)
	// }
	//
	// cfg, err = config.Read()
	// if err != nil {
	// 	log.Fatalf("error reading config: %v", err)
	// }
	// fmt.Printf("Read config again: %+v\n", cfg)
}
