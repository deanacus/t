package main

import (
	"fmt"
	"os"
	"strings"
)

const VERSION string = "0.0.1"

const HELP string = `t - The minimalist command line todo manager

Usage: t [option] <Input>

Options:
  -a, --add <Task>              Add a task to the todo list
  -c, --complete <TaskID>       Move a task to the completed list
  -d, --delete <TaskID>         Delete a task from the todo list
  -e, --edit <TaskID> <Task>    Edit a task in the todo list
  -v, --version                 Show the version
  -h, --help                    Show this help information`

func main() {
	if len(os.Args) <= 1 {
		os.Args = append(os.Args, "l")
	}

	var flag string = strings.TrimPrefix(os.Args[1], "-")
	switch flag {
		case "h":
			showHelp()
		case "l":
			listItems()
		case "a":
			addItem()
		case "d":
			deleteItem()
		case "v":
			showVersion()
	}

	readfile()

}

func listItems() {

	fmt.Println("list the stuff")

}

func addItem() {

	fmt.Println("add the stuff")

}

func deleteItem() {

	fmt.Println("delete the stuff")

}

func showHelp() {

	fmt.Println(HELP)

}

func showVersion() {
	fmt.Println("Version:", VERSION)
}
