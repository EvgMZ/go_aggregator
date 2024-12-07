package main

import (
	"aggregator/internal/config"
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
)
type State struct {
	Cfg *config.Config
	Db  *database.Queries
}
func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	com := Commands{
		Handlers: make(map[string]func(*State, Command) error),
	}
	state := &State{
		Cfg: &cfg,
	}
	command := Command{}
	com.Register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: not enough arguments provided")
		os.Exit(1)
	}
	command.Name = os.Args[1]
	command.Args = os.Args[2:]
	// cmd := command{name: cmdName, args: cmdArgs}

	if err := com.Run(state, command); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)
	state.Db = dbQueries
}
