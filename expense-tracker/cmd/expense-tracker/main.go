package main

import (
	"fmt"
	"log"
	"os"

	"github.com/0xirvan/roadmap-projects/expense-tracker/internal/commands"
	"github.com/0xirvan/roadmap-projects/expense-tracker/internal/expense"
	"github.com/0xirvan/roadmap-projects/expense-tracker/internal/storage"
)

func main() {
	store := storage.NewJSONStorage("data/expenses.json")
	repo := expense.NewRepository(store)
	svc := expense.NewService(repo)

	if len(os.Args) < 2 {
		commands.ShowHelp()
		return
	}
	command := os.Args[1]

	switch command {
	case "add":
		c := &commands.AddCmd{Svc: svc}
		if err := c.Run(); err != nil {
			log.Fatal(err)
		}
	case "list":
		c := &commands.ListCmd{Svc: svc}
		if err := c.Run(); err != nil {
			log.Fatal(err)
		}
	case "summary":
		c := &commands.SummaryCmd{Svc: svc}
		if err := c.Run(); err != nil {
			log.Fatal(err)
		}
	case "help":
		commands.ShowHelp()
	default:
		fmt.Println("unknown command, please see help")
	}
}
