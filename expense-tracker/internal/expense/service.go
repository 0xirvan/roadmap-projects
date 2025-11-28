package expense

import (
	"errors"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Add(description string, amount float64) (uint, error) {
	if description == "" || amount < 0 {
		return 0, errors.New("validation error: description cannot be empty and amount must be greater than 0")
	}
	expenses, err := s.repo.GetAll()
	if err != nil {
		return 0, err
	}

	newExpense := Expense{
		ID:          s.repo.GetNextID(),
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}
	expenses = append(expenses, newExpense)

	err = s.repo.SaveAll(expenses)
	if err != nil {
		return 0, err
	}
	return newExpense.ID, err
}

func (s *Service) List() ([]Expense, error) {
	return s.repo.GetAll()
}

func (s *Service) Summary(filterMonth int) (float64, error) {
	expenses, err := s.repo.GetAll()
	if err != nil {
		return 0, err
	}

	var total float64
	for _, expense := range expenses {
		if filterMonth > 0 && int(expense.Date.Month()) != filterMonth {
			continue
		}
		total += expense.Amount
	}

	return total, nil
}

func (s *Service) Delete(id uint) error {
	expenses, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	expense, idx := s.repo.FindByID(expenses, id)
	if expense == nil {
		return errors.New("expense not found")
	}

	expenses = append(expenses[:idx], expenses[idx+1:]...)

	return s.repo.SaveAll(expenses)
}
