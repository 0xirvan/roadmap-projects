package cmd

import (
	"fmt"
	"strconv"

	"github.com/0xirvan/roadmap-projects/task-tracker/internal/task"
)

type DeleteCmd struct {
	Svc *task.TaskService
}

func (c *DeleteCmd) Run(args []string) error {
	if len(args) < 1 {
		fmt.Println("usage: task-tracker delete <id>")
		return nil
	}

	id, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return err
	}

	if err := c.Svc.Delete(id); err != nil {
		return err
	}

	return nil
}
