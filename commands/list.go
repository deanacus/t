package commands

import (
	"fmt"
	"strconv"
	"t/colour"
	"t/utils"
)

// ListItems reads the list of tasks from the configured
// file, then prints them out to the terminal with a header
// # TODO
func ListItems() {

	tasks := utils.CheckFile("./.todo")

	length := len(tasks)

	newLine := "\n"

	fmt.Println(newLine + colour.BoldBlue("# TODO") + newLine)

	if length > 0 {
		for i, task := range tasks {
			id := strconv.Itoa(i+1) + "."
			fmt.Println(colour.Faint(id), task)
		}
	} else {
		fmt.Println(colour.BoldGreen("Nothing to do"))
	}

}
