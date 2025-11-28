package utils

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/0xirvan/roadmap-projects/expense-tracker/internal/expense"
)

func TablePrint(w io.Writer, data []expense.Expense) {
	writer := tabwriter.NewWriter(w, 0, 0, 3, ' ', 0)
	fmt.Fprintln(writer, "ID\tDate\tDescription\tAmount")

	for _, e := range data {
		fmt.Fprintf(writer, "%d\t%s\t%s\t$%.2f\n", e.ID, e.Date.Format("2006-01-02"), e.Description, e.Amount)
	}

	writer.Flush()
}
