package main

import (
	"os"

	"github.com/0xirvan/roadmap-projects/task-tracker/cmd"
	"github.com/0xirvan/roadmap-projects/task-tracker/internal/storage"
	"github.com/0xirvan/roadmap-projects/task-tracker/internal/task"
)

func main() {
	store := storage.NewJSONStorage("data/tasks.json")
	repo := task.NewTaskRepository(store)
	svc := task.NewService(repo)

	if len(os.Args) < 2 {
		cmd.ShowHelp()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		c := &cmd.AddCmd{Svc: svc}
		if err := c.Run(args); err != nil {
			panic(err)
		}
	case "list":
		c := &cmd.ListCmd{Svc: svc}
		if err := c.Run(args); err != nil {
			panic(err)
		}
	case "update":
		c := &cmd.UpdateCmd{Svc: svc}
		if err := c.Run(args); err != nil {
			panic(err)
		}
	case "delete":
		c := &cmd.DeleteCmd{Svc: svc}
		if err := c.Run(args); err != nil {
			panic(err)
		}
	case "mark-done":
		c := &cmd.MarkCmd{Svc: svc}
		if err := c.Run(args, task.StatusDone); err != nil {
			panic(err)
		}
	case "mark-in-progress":
		c := &cmd.MarkCmd{Svc: svc}
		if err := c.Run(args, task.StatusInProgress); err != nil {
			panic(err)
		}
	}
}
