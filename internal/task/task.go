package task

import (
	"fmt"
)

// Task represents a single task in the to-do list.
type Task struct {
	Title  string
	Status bool // false for pending, true for completed
}

// String returns a formatted string representation of the task.
func (t Task) String() string {
	status := "Pending"
	if t.Status {
		status = "Completed"
	}
	return fmt.Sprintf("Title: %s, Status: %s", t.Title, status)
}
