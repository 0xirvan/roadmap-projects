package cmd

import "fmt"

func ShowHelp() {
	fmt.Print(`Usage:
  task-tracker <command> [arguments]

Commands:
  add <description>           Add a new task
  list                        List all tasks
  list done                   List completed tasks
  list todo                   List pending tasks
  list in-progress            List tasks in progress
  mark-in-progress <id>       Mark a task as in-progress
  mark-done <id>              Mark a task as done
  update <id> <description>   Update an existing task
  delete <id>                 Delete a task

Examples:
  task-tracker add "learn Go"
  task-tracker update 1 "learn Go programming"
  task-tracker delete 2
  task-tracker mark-in-progress 1
  task-tracker mark-done 1
  task-tracker list done
  task-tracker list todo
  task-tracker list in-progress
`)
}
