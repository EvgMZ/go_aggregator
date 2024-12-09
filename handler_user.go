package main

import (
	"aggregator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]
	user, err := s.Db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user not found %w", err)
	}
	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Println("User set to ", username)
	return nil
}

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	name := cmd.Args[0]
	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	return nil
}

func handlerReset(s *State, cmd Command) error {
	_, err := s.Db.DeleteAllUser(context.Background())
	if err != nil {
		return fmt.Errorf("can't delete all users %w", err)
	}
	return nil
}

func handlerGetUsers(s *State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("don't get users %w", err)
	}
	cur := s.Cfg.Current_user_name
	printUsers(users, cur)
	return nil
}

func printUsers(users []string, current string) {
	for _, item := range users {
		if item == current {
			fmt.Println("* ", item, "(current)")
		} else {
			fmt.Println("* ", item)
		}
	}
}
