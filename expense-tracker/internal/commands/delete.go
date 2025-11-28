package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/0xirvan/roadmap-projects/expense-tracker/internal/expense"
)

type DeleteCmd struct {
	Svc *expense.Service
}

func (c *DeleteCmd) Run() error {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	id := deleteCmd.Uint("id", 0, "the expense id to delete")

	if len(os.Args) > 2 {
		deleteCmd.Parse(os.Args[2:])
	}
	err := c.Svc.Delete(*id)
	if err != nil {
		return err
	}
	fmt.Println("expense delete successfully")
	return nil
}
