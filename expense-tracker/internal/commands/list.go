package commands

import (
	"os"

	"github.com/0xirvan/roadmap-projects/expense-tracker/internal/expense"
	"github.com/0xirvan/roadmap-projects/expense-tracker/internal/utils"
)

type ListCmd struct {
	Svc *expense.Service
}

func (c *ListCmd) Run() error {
	expenses, err := c.Svc.List()
	if err != nil {
		return err
	}

	utils.TablePrint(os.Stdout, expenses)
	return nil
}
