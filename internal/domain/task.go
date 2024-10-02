package domain

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	TODO        TaskStatus = "todo"
	IN_PROGRESS TaskStatus = "in-progress"
	DONE        TaskStatus = "done"
)

type Task struct {
	ID          int
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewTask creates a new task
func NewTask(description string) *Task {
	return &Task{
		Description: description,
		Status:      TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// Update updates a given task
func (t *Task) Update(description string) {
	t.Description = description
	t.UpdatedAt = time.Now()
}

// SetStatus changes the task status and set updatedAt to time.Now()
func (t *Task) SetStatus(status TaskStatus) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

// String stringifies the task
func (t *Task) String() string {
	return fmt.Sprintf("ID: %d, Description: %s, Status: %s, Created At: %s, Updated At: %s\n",
		t.ID, t.Description, t.Status, t.CreatedAt.Format(time.RFC3339), t.UpdatedAt.Format(time.RFC3339))
}
