package gator

import (
	"aggregator/internal/config"
	"fmt"
)

type State struct {
	Cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}
type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.Handlers == nil {
		c.Handlers = make(map[string]func(*State, Command) error)
	}
	c.Handlers[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, exists := c.Handlers[cmd.Name]
	if !exists {
		fmt.Errorf("command does not exists")
	}
	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
