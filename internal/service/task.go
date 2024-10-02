package service

import (
	"github.com/lazhari/task-cli/internal/domain"
	"github.com/lazhari/task-cli/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new TaskService instance
func NewTaskService(repo repository.TaskRepository) domain.TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) AddTask(description string) (*domain.Task, error) {
	task := domain.NewTask(description)
	err := s.repo.AddTask(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) UpdateTask(id int, description string) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	task.Update(description)
	return s.repo.Save()
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.DeleteByID(id)
}

func (s *TaskService) SetStatus(id int, status domain.TaskStatus) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil
	}
	task.SetStatus(domain.IN_PROGRESS)
	return s.repo.Save()
}

func (s *TaskService) ListAllTasks() ([]*domain.Task, error) {
	return s.repo.LoadTasks()
}

func (s *TaskService) ListTasksByStatus(status domain.TaskStatus) ([]*domain.Task, error) {
	tasks, err := s.repo.LoadTasks()
	if err != nil {
		return nil, err
	}
	var filteredTasks []*domain.Task
	for _, t := range tasks {
		if t.Status == status {
			filteredTasks = append(filteredTasks, t)
		}
	}

	return filteredTasks, nil
}
