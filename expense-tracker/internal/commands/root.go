package commands

import "fmt"

func ShowHelp() {
	const helpText = `Expense Tracker CLI
Simple tool to manage your finances.

Usage: 
  expense-tracker <command> [flags]

Commands:
  add       Add a new expense
            Flags:
              --description <string>  Description of the expense (required)
              --amount      <int>     Amount of the expense (required)

  list      List all added expenses
            (No flags required)

  summary   Show total expenses
            Flags:
              --month       <int>     Filter summary by month (1-12) (optional)

  delete    Delete an expense by ID
            Flags:
              --id          <int>     ID of the expense to delete (required)

Examples:
  $ expense-tracker add --description "Lunch" --amount 20
  $ expense-tracker list
  $ expense-tracker summary --month 8
  $ expense-tracker delete --id 2
`
	fmt.Print(helpText)
}
