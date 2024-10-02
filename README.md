# Task Tracker
The Task Tracker is a simple task management application built using Go and designed based on Hexagonal architecture principles.

## Project URL

[https://roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/task-tracker)

## Prerequisites

Make sure you have Go installed. You can download and install it from [here](https://go.dev/dl/).

Check your Go installation:

```bash
go version
```

## Installation

### Clone the Repository

```bash
git clone https://github.com/lazhari/task-cli.git
cd task-cli
```

### Running the Project

You have two options to run the project: using `go run` or `go install`.

#### Option 1: Run Directly with go run

```bash
go run ./cmd/task_cli.go <command> [arguments]
```
##### Example Usage:
- Add a task
  ```bash
  go run . add "Buy groceries"
  ```
- List all tasks:
  ```bash
  go run . list
  ```

#### Option 2: Install and Run

```bash
cd task-cli
go install .
```

##### Example Usage:
- Add a task
  ```bash
  task-cli add "Buy groceries"
  ```
- List all tasks:
  ```bash
  task-cli list
  ```

## Commands

| Command                   | Description                                    |
|----------------------------|------------------------------------------------|
| `add <description>`         | Add a new task                                |
| `update <id> <description>` | Update the description of an existing task     |
| `delete <id>`               | Delete a task by its ID                       |
| `mark-in-progress <id>`     | Mark a task as in progress                    |
| `mark-done <id>`            | Mark a task as done                           |
| `list`                      | List all tasks                                |
| `list todo`                 | List tasks that are not started (todo)        |
| `list in-progress`          | List tasks that are in progress               |
| `list done`                 | List tasks that are done                      |

## Task File Location
The tasks are stored in a JSON file in the user's home directory:

- **Linux/macOS**: `~/.task-tracker.json`
- **Windows**: `C:\Users\<YourUserName>\.task-tracker.json`