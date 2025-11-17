package cmd

import (
	"github.com/0xirvan/roadmap-projects/task-tracker/internal/task"
	"github.com/0xirvan/roadmap-projects/task-tracker/utils"
)

type ListCmd struct {
	Svc *task.TaskService
}

func (c *ListCmd) Run(args []string) error {
	var filter task.Status

	if len(args) > 0 {
		switch args[0] {
		case "done":
			filter = task.StatusDone
		case "todo":
			filter = task.StatusTodo
		case "in-progress":
			filter = task.StatusInProgress
		}
	}

	tasks, err := c.Svc.List(filter)
	if err != nil {
		return err
	}

	utils.PrintTasks(tasks)
	return nil
}
