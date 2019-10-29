package commands

import (
	"fmt"
	"strconv"
	"t/utils"
)

// DeleteItem reads the tasks out of the configured file
// then filters the resulting slice base on index and taskID
// and finally writes the resulting slice to the file before
// letting the user know it has been successful
func DeleteItem(input []string) {

	if len(input) >= 2 {
		fmt.Println("\033[31mToo many values supplied\033[0m")
		return
	}

	// convert the taskID to int64, because I can't
	// work out how to convert to int
	taskID, _ := strconv.ParseInt(input[0], 10, 0)

	// Read file into array/slice
	tasks := utils.CheckFile("./.todo")

	// Check that the task exists
	if len(tasks) < int(taskID) {
		fmt.Println("\033[31mTask does not exist\033[0m")
		return
	}

	// Filter the task out of the slice
	newTasks := utils.Filter(tasks, func(ind int) bool {
		return ind != int(taskID)-1
	})

	// Write the slice back to the file
	success := utils.WriteTasks("./.todo", newTasks)

	if success {
		fmt.Println("\033[31mTask has been yeeted\033[0m")
	}

}
