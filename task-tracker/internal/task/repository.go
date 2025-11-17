package task

import (
	"sync"

	"github.com/0xirvan/roadmap-projects/task-tracker/internal/storage"
)

type TaskRepository struct {
	store  *storage.JSONStorage
	lastID uint64
	mu     sync.RWMutex
}

func NewTaskRepository(store *storage.JSONStorage) *TaskRepository {
	r := &TaskRepository{store: store}

	var tasks []Task
	if err := store.LoadJSON(&tasks); err == nil {
		var max uint64
		for _, t := range tasks {
			if t.ID > max {
				max = t.ID
			}
		}
		r.lastID = max
	}

	return r
}

func (r *TaskRepository) GetAll() ([]Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var tasks []Task

	err := r.store.LoadJSON(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) Save(tasks []Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.store.SaveJSON(tasks)
}

func (r *TaskRepository) FindByID(tasks []Task, id uint64) (*Task, uint64) {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], uint64(i)
		}
	}
	return nil, 0
}

func (r *TaskRepository) GetNextID() uint64 {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++
	return r.lastID
}
