package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lazhari/task-cli/internal/cli"
	"github.com/lazhari/task-cli/internal/service"
	"github.com/lazhari/task-cli/internal/storage"
)

func Execute() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: Unable to get the home directory")
		os.Exit(1)
	}
	// Define the path to the task file in the home directory
	tasksFilePath := filepath.Join(homeDir, ".task-tracker.json")
	// Initialize storage (JSON-based)
	repo := storage.NewJSONTaskRepository(tasksFilePath)

	// Initialize the service
	taskService := service.NewTaskService(repo)

	// Initialize the CLI handler
	cliHandler := cli.NewCLIHandler(taskService)

	// Execute the CLI handler with command-line arguments
	cliHandler.Run(os.Args)
}
