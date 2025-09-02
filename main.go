package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/thomas-henley/go-gator/internal/config"
	"github.com/thomas-henley/go-gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	dbQueries := database.New(db)

	programState := state{
		db: dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)

	arguments := os.Args

	if len(arguments) < 2 {
		log.Fatalf("Usage: cli <command> [args...]\n")
	}

	cmd := command{
		Name: arguments[1],
		Args: arguments[2:],
	}

	err = cmds.run(&programState, cmd)
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
}
