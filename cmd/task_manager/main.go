package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Daniel-Q-Reis/go-task-manager-cli/internal/manager"
)

func main() {
	tm := manager.NewTaskManager()      // Create a new instance of TaskManager
	reader := bufio.NewReader(os.Stdin) // To read the full line, incluinding spaces

	for {
		fmt.Println("----------------------- MENU ------------------------")
		fmt.Println()
		fmt.Println(`Enter:
    1 to Add a new task;
    2 to List/Manage tasks;
    0 to Exit`)

		fmt.Print("\nEnter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil || (choice != 0 && choice != 1 && choice != 2) {
			fmt.Println("Error: Please enter 1, 2, or 0.")
			continue
		}

		switch choice {
		case 0:
			fmt.Println("Exiting program...")
			return
		case 1:
			tm.AddTask()
		case 2:
			tm.ManageTasks()
		default:
			// Should not be reached due to input validation
		}
	}
}
