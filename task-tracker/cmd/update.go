package cmd

import (
	"fmt"
	"strconv"

	"github.com/0xirvan/roadmap-projects/task-tracker/internal/task"
)

type UpdateCmd struct {
	Svc *task.TaskService
}

func (c *UpdateCmd) Run(args []string) error {
	if len(args) < 2 {
		fmt.Println("usage: task-tracker update <id> <description>")
		return nil
	}

	id, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return err
	}

	description := args[1]

	if err := c.Svc.Update(id, description); err != nil {
		return err
	}

	return nil
}
