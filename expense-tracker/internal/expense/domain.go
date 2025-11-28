package expense

import "time"

type Expense struct {
	ID          uint      `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
}
