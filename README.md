# Go Task Manager CLI

This is a simple command-line interface (CLI) application developed in Go for managing your daily tasks. It allows you to add new tasks, list existing tasks, and toggle the status of a task between 'Pending' and 'Completed'.

---

## Features

* **Add Task:** Prompt to enter the title of a new task.
* **List/Manage Tasks:** View all tasks with their current statuses and change the status of a specific task.

---

## How to Run

1.  **Prerequisites:** Ensure you have [Go](https://golang.org/doc/install) installed on your machine (version 1.18+ recommended).

2.  **Clone the Repository:**
    ```bash
    git clone [https://github.com/YOUR_USERNAME/go-task-manager-cli.git](https://github.com/YOUR_USERNAME/go-task-manager-cli.git)
    cd go-task-manager-cli
    ```
    *(Replace `YOUR_USERNAME` with your GitHub username).*

3.  **Run the Application:**
    ```bash
    go run ./cmd/task_manager/main.go
    ```
    Alternatively, to build and run the executable:
    ```bash
    go build -o task_manager ./cmd/task_manager/main.go
    ./task_manager
    ```

---

## Project Structure

```bash
go-task-manager-cli/
├── cmd/                    # Contains main executable entry points
│   └── task_manager/       # CLI application for task management
│       └── main.go         # Application entry point
├── internal/               # Internal packages not intended for public import
│   ├── manager/            # Business logic for managing tasks (add, manage, etc.)
│   │   └── manager.go
│   └── task/               # Defines the Task struct and related methods
│       └── task.go
├── .gitignore              # Specifies intentionally untracked files to ignore
├── LICENSE                 # MIT License details
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
└── README.md               # Project documentation
```
---

## Technologies Used

* **Go**
* Standard Go packages: `fmt`, `strconv`, `strings`, `bufio`, `os`

---

## Future Enhancements (Possible Improvements)

* Data persistence (save tasks to a file or a database so they are not lost upon closing).
* Add functionality to remove tasks.
* Add functionality to edit task titles.
* Implement more robust error handling
* Add better UX