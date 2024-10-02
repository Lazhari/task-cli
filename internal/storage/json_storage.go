package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/lazhari/task-cli/internal/domain"
	"github.com/lazhari/task-cli/internal/repository"
)

var TaskNotFound = errors.New("task not found")

type JSONTaskRepository struct {
	FilePath string
	tasks    []*domain.Task
}

func NewJSONTaskRepository(filePath string) repository.TaskRepository {
	return &JSONTaskRepository{
		FilePath: filePath,
		tasks:    []*domain.Task{},
	}
}

func (r *JSONTaskRepository) LoadTasks() ([]*domain.Task, error) {
	if _, err := os.Stat(r.FilePath); os.IsNotExist(err) {
		r.tasks = []*domain.Task{}
		return r.tasks, nil
	}

	data, err := os.ReadFile(r.FilePath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &r.tasks); err != nil {
		return nil, err
	}

	return r.tasks, nil
}

func (r *JSONTaskRepository) Save() error {
	bytes, err := json.MarshalIndent(r.tasks, "", "  ")
	if err != nil {
		return err
	}

	// Write data to the file, creating the file if it doesn't exist
	file, err := os.OpenFile(r.FilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	return err
}

func (r *JSONTaskRepository) AddTask(task *domain.Task) error {
	if _, err := r.LoadTasks(); err != nil {
		return err
	}
	task.ID = r.generateNewTaskID()
	r.tasks = append(r.tasks, task)
	return r.Save()
}

func (r *JSONTaskRepository) FindByID(id int) (*domain.Task, error) {
	if _, err := r.LoadTasks(); err != nil {
		return nil, err
	}
	for _, t := range r.tasks {
		if t.ID == id {
			return t, nil
		}
	}

	return nil, TaskNotFound
}

func (r *JSONTaskRepository) DeleteByID(id int) error {
	if _, err := r.LoadTasks(); err != nil {
		return err
	}
	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			// Store the new
			return r.Save()
		}
	}

	return TaskNotFound
}

// Generate a new unique ID by finding the highest existing ID
func (r *JSONTaskRepository) generateNewTaskID() int {
	maxID := 0
	for _, t := range r.tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	return maxID + 1
}
