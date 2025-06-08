package commands

import (
	"github.com/bulkashmak/gator-cli/internal"
  "errors"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Store map[string]func(*internal.State, Command) error
}

func (c *Commands) Run(s *internal.State, cmd Command) error {
	fun, ok := c.Store[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}

	return fun(s, cmd)
}

func (c *Commands) Register(name string, f func(*internal.State, Command) error) {
	c.Store[name] = f
}
