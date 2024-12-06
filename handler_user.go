package main

import (
	"aggregator/internal/gator"
	"fmt"
)

func handlerLogin(s *gator.State, cmd gator.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]
	err := s.Cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("User set to ", username)
	return nil
}
