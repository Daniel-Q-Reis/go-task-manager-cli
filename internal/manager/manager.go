package manager

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Daniel-Q-Reis/go-task-manager-cli/internal/task"
)

// TaskManager handles operations related to tasks.
type TaskManager struct {
	Tasks []task.Task
}

// NewTaskManager creates and returns a new TaskManager.
func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make([]task.Task, 0),
	}
}

// AddTask prompts the user for a task title and adds it to the list.
func (tm *TaskManager) AddTask() {
	var title string
	fmt.Print("Enter the title of the task you want to add: ")
	fmt.Scanln(&title)
	title = strings.TrimSpace(title)

	if title == "" {
		fmt.Println("Error: Task title cannot be empty.")
		return
	}

	tm.Tasks = append(tm.Tasks, task.Task{Title: title, Status: false})
	fmt.Printf("Task '%s' added successfully!\n", title)
	// Optionally, show menu again or guide user
}

// ManageTasks allows the user to view and change task statuses.
func (tm *TaskManager) ManageTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("No tasks registered yet.")
		return
	}

	for {
		fmt.Println("\n--- Current Tasks ---")
		for i, t := range tm.Tasks {
			statusStr := "Pending"
			if t.Status {
				statusStr = "Completed"
			}
			fmt.Printf("%d. %s [%s]\n", i+1, t.Title, statusStr)
		}
		fmt.Println("---------------------\n")

		fmt.Println("Enter the task number to change its status (or 0 to return to main menu):")
		var input string
		fmt.Scanln(&input)

		if input == "0" {
			return
		}

		taskNum, err := strconv.Atoi(input)
		if err != nil || taskNum < 1 || taskNum > len(tm.Tasks) {
			fmt.Println("Error: Please enter a valid task number from the list.")
			continue
		}

		// Toggle task status
		tm.Tasks[taskNum-1].Status = !tm.Tasks[taskNum-1].Status

		statusStr := "Pending"
		if tm.Tasks[taskNum-1].Status {
			statusStr = "Completed"
		}
		fmt.Printf("Task '%s' status changed successfully to: %s\n", tm.Tasks[taskNum-1].Title, statusStr)
	}
}
