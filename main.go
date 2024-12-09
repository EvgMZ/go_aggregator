package main

import (
	"aggregator/internal/config"
	"aggregator/internal/database"
	"context"
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
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting db", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	state := &State{
		Cfg: &cfg,
		Db:  dbQueries,
	}
	com := Commands{
		Handlers: make(map[string]func(*State, Command) error),
	}
	command := Command{}
	com.Register("login", handlerLogin)
	com.Register("register", handlerRegister)
	com.Register("reset", handlerReset)
	com.Register("users", handlerGetUsers)
	com.Register("agg", handlerAgg)
	// com.Register("addfeed", handlerAddFedd)
	com.Register("feeds", handlerGetFeeds)
	// com.Register("follow", handlerFollow)
	// com.Register("following", handlerListFeedFollows)
	com.Register("following", middlewareLoggedIn(handlerListFeedFollows))
	com.Register("follow", middlewareLoggedIn(handlerFollow))
	com.Register("addfeed", middlewareLoggedIn(handlerAddFedd))
	com.Register("unfollow", middlewareLoggedIn(handlerUnfollow))
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

	// state.Db = dbQueries
}
func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.Current_user_name)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
