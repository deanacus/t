package commands

import (
	"fmt"
	"strings"
	"t/utils"
)

// AddItem reads the list of tasks from the configured
// file, then checks to see if the new task already exists
// and pushes it to the slice if not. It then writes the
// resulting slice to the configured file.
func AddItem(input []string) {

	tasks := utils.CheckFile("./.todo")

	task := strings.Join(input, " ")

	// Check that the task doesn't already exist
	found := utils.Find(tasks, task)

	if found {
		fmt.Println("Task already exists")
		return
	}
	// Add the task to the end of the slice
	tasks = append(tasks, task)

	if utils.WriteTasks("./.todo", tasks) {
		fmt.Println("\033[1;35mTask added\033[0m")
	}

}
