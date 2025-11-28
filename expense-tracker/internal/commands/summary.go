package commands

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/0xirvan/roadmap-projects/expense-tracker/internal/expense"
)

type SummaryCmd struct {
	Svc *expense.Service
}

func (c *SummaryCmd) Run() error {
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)

	filterMonth := summaryCmd.Int("month", 0, "summary by month")

	if len(os.Args) > 2 {
		summaryCmd.Parse(os.Args[2:])
	}

	total, err := c.Svc.Summary(*filterMonth)
	if err != nil {
		return err
	}

	var message string
	if *filterMonth > 0 {
		message = fmt.Sprintf("total expenses for %s: %.2f", time.Month(*filterMonth).String(), total)
	} else {
		message = fmt.Sprintf("total expenses: $%.2f", total)
	}

	fmt.Println(message)
	return nil
}
