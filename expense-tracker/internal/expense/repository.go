package expense

import (
	"sync"

	"github.com/0xirvan/roadmap-projects/expense-tracker/internal/storage"
)

type Repository struct {
	store  *storage.JSONStorage
	lastID uint
	mu     sync.RWMutex
}

func NewRepository(store *storage.JSONStorage) *Repository {
	r := &Repository{store: store, lastID: 1}

	var expenses []Expense
	if err := r.store.LoadJSON(&expenses); err == nil {
		var maxID uint
		for _, e := range expenses {
			if e.ID > maxID {
				maxID = e.ID
			}
		}
		r.lastID = maxID
	}
	return r
}

func (r *Repository) GetAll() ([]Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var expenses []Expense
	err := r.store.LoadJSON(&expenses)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (r *Repository) SaveAll(expenses []Expense) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.store.SaveJSON(expenses)
}

func (r *Repository) FindByID(expenses []Expense, id uint) (*Expense, uint) {
	for i := range expenses {
		if expenses[i].ID == id {
			return &expenses[i], id
		}
	}
	return nil, 0
}

func (r *Repository) GetNextID() uint {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.lastID++
	return r.lastID
}
