package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/0xirvan/roadmap-projects/expense-tracker/internal/expense"
)

type AddCmd struct {
	Svc *expense.Service
}

func (c *AddCmd) Run() error {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	description := addCmd.String("description", "", "description expense to add")
	amount := addCmd.Float64("amount", 0, "amount of the expense")

	if len(os.Args) >= 2 {
		addCmd.Parse(os.Args[2:])
	}

	id, err := c.Svc.Add(*description, *amount)
	if err != nil {
		return err
	}
	fmt.Printf("expense added succesfully (ID: %d)", id)
	return nil
}
