package task

import (
	"time"
)

type TaskService struct {
	repo *TaskRepository
}

func NewService(repo *TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Add(description string) (uint64, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return 0, err
	}

	newTask := Task{
		ID:          s.repo.GetNextID(),
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	tasks = append(tasks, newTask)

	if err := s.repo.Save(tasks); err != nil {
		return 0, err
	}

	return newTask.ID, nil
}

func (s *TaskService) List(filter Status) ([]Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	if filter == "" {
		return tasks, nil
	}

	var filtered []Task
	for _, task := range tasks {
		if task.Status == filter {
			filtered = append(filtered, task)
		}
	}

	return filtered, nil
}

func (s *TaskService) Update(id uint64, description string) error {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	task, idx := s.repo.FindByID(tasks, id)
	if task == nil {
		return ErrNotFound
	}

	task.Description = description
	task.UpdatedAt = time.Now().Format(time.RFC3339)

	tasks[idx] = *task

	return s.repo.Save(tasks)
}

func (s *TaskService) Delete(id uint64) error {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	task, idx := s.repo.FindByID(tasks, id)
	if task == nil {
		return ErrNotFound
	}

	tasks = append(tasks[:idx], tasks[idx+1:]...)

	return s.repo.Save(tasks)
}

func (s *TaskService) Mark(id uint64, status Status) error {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	task, idx := s.repo.FindByID(tasks, id)
	if task == nil {
		return ErrNotFound
	}

	task.Status = status
	task.UpdatedAt = time.Now().Format(time.RFC3339)

	tasks[idx] = *task

	return s.repo.Save(tasks)
}
