package commands

import (
	"fmt"
	"strconv"
	"t/utils"
)

// ListItems reads the list of tasks from the configured
// file, then prints them out to the terminal with a header
// # TODO
func ListItems() {

	tasks := utils.CheckFile("./.todo")

	length := len(tasks)

	newLine := "\n"

	fmt.Println(newLine + "\033[1;34m# TODO\033[0m" + newLine)

	if length > 0 {
		for i, task := range tasks {
			id := strconv.Itoa(i+1) + "."
			fmt.Println("\033[2m"+id+"\033[0m", task)
		}
	} else {
		fmt.Println("Nothing to do")
	}

}
