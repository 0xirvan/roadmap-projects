package cmd

import (
	"fmt"
	"strconv"

	"github.com/0xirvan/roadmap-projects/task-tracker/internal/task"
)

type MarkCmd struct {
	Svc *task.TaskService
}

func (c *MarkCmd) Run(args []string, status task.Status) error {
	if len(args) < 1 {
		fmt.Println("usage: task-tracker mark-<status> <id>")
		return nil
	}

	id, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return err
	}

	if err := c.Svc.Mark(id, status); err != nil {
		return err
	}
	return nil
}
