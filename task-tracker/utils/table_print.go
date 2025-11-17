package utils

import (
	"fmt"
	"time"

	"github.com/0xirvan/roadmap-projects/task-tracker/internal/task"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
)

func colorStatus(status string) string {
	switch status {
	case "done":
		return ColorGreen + "✓ " + status + ColorReset
	case "in-progress":
		return ColorYellow + "⟳ " + status + ColorReset
	case "todo":
		return ColorRed + "○ " + status + ColorReset
	default:
		return status
	}
}

func PrintTasks(tasks []task.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for i, t := range tasks {

		created, _ := time.Parse(time.RFC3339, t.CreatedAt)
		updated, _ := time.Parse(time.RFC3339, t.UpdatedAt)

		fmt.Println("┌────────────────────────────────────────────────────────────┐")
		fmt.Printf("│ ID          : %-44d │\n", t.ID)
		fmt.Printf("│ Description : %-44s │\n", truncate(t.Description, 44))
		fmt.Printf("│ Status      : %-53s │\n", colorStatus(string(t.Status)))
		fmt.Printf("│ Created     : %-44s │\n", humanize(created))
		fmt.Printf("│ Updated     : %-44s │\n", humanize(updated))
		fmt.Println("└────────────────────────────────────────────────────────────┘")
		if i < len(tasks)-1 {
			fmt.Println()
		}
	}
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
