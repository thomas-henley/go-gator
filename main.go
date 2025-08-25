package main

import (
	"fmt"
	"os"

	"github.com/thomas-henley/go-gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

	programState := state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)

	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Printf("Usage: cli <command> [args...]\n")
		os.Exit(1)
	}

	cmd := command{
		Name: arguments[1],
		Args: arguments[2:],
	}

	err = cmds.run(&programState, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
