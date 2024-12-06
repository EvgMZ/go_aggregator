package main

import (
	"aggregator/internal/config"
	"aggregator/internal/gator"
	"fmt"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	com := gator.Commands{
		Handlers: make(map[string]func(*gator.State, gator.Command) error),
	}
	state := &gator.State{
		Cfg: &cfg,
	}
	command := gator.Command{}
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

}
