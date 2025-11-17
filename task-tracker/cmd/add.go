package cmd

import (
	"fmt"

	"github.com/0xirvan/roadmap-projects/task-tracker/internal/task"
)

type AddCmd struct {
	Svc *task.TaskService
}

func (c *AddCmd) Run(args []string) error {
	if len(args) < 1 {
		fmt.Println("usage: task-tracker add <description>")
		return nil
	}

	description := args[0]

	id, err := c.Svc.Add(description)
	if err != nil {
		return err
	}

	fmt.Printf("Task added with ID %d\n", id)
	return nil
}
