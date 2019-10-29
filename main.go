package main

import (
	"fmt"
	"os"
	"strings"

	"t/commands"
)

const versionString string = "0.0.1"

const helpString string = `t - The minimalist command line todo manager

Usage: t [option] <Input>

Options:
  -l, --list                    List the tasks in the todo list
  -a, --add <Task>              Add a task to the todo list
  -d, --delete <TaskID>         Delete a task from the todo list
  -v, --version                 Show the version
  -h, --help                    Show this help information`

func main() {
	if len(os.Args) <= 1 {
		os.Args = append(os.Args, "l")
	}

	var input = os.Args[2:]

	var flag string = strings.TrimPrefix(os.Args[1], "-")

	switch flag {
	case "l":
		commands.ListItems()
	case "a":
		commands.AddItem(input)
	case "d":
		commands.DeleteItem(input)
	case "h":
		showHelp()
	case "v":
		showVersion()
	}

}

func showHelp() {
	fmt.Println(helpString)
}

func showVersion() {
	fmt.Println("Version:", versionString)
}
