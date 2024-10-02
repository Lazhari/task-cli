package domain

import (
	"fmt"
	"testing"
	"time"
)

func TestTask_NewTask(t *testing.T) {
	description := "Test Task"
	task := NewTask(description)

	if task.Description != description {
		t.Errorf("expected %v, got %v", description, task.Description)
	}

	if task.Status != TODO {
		t.Errorf("expect status to be TODO, got %v", task.Status)
	}

	if task.CreatedAt.IsZero() || task.UpdatedAt.IsZero() {
		t.Error("expect CreatedAt and UpdatedAt to be set")
	}
}

func TestTask_SetStatus(t *testing.T) {
	task := NewTask("Test Task")
	task.SetStatus(IN_PROGRESS)

	if task.Status != IN_PROGRESS {
		t.Errorf("expected status to be IN_PROGRESS, got %v", task.Status)
	}

	if task.UpdatedAt.Before(task.CreatedAt) {
		t.Error("expected updatedAt to be after createdAt")
	}
}

func TestTask_Update(t *testing.T) {
	description := "Updated Test Task"
	task := NewTask("Test Task")

	task.Update(description)

	if task.Description != description {
		t.Errorf("expected description to be %v, got %v", description, task.Description)
	}

	if task.UpdatedAt.Before(task.CreatedAt) {
		t.Error("expected updatedAt to be after createdAt")
	}
}

func TestTask_String(t *testing.T) {
	// Create a sample task with known values
	createdAt := time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2023, 10, 2, 14, 30, 0, 0, time.UTC)

	task := Task{
		ID:          1,
		Description: "Test Task",
		Status:      TODO,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	// Call the String() method
	result := task.String()

	// Expected string format
	expected := fmt.Sprintf(
		"ID: %d, Description: %s, Status: %s, Created At: %s, Updated At: %s\n",
		task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339),
	)

	// Compare the result with the expected string
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
