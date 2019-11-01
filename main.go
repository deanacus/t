package main

import (
	"fmt"
	"os"
	"strings"

	"t/colour"
	"t/commands"
	"t/config"
)

var versionString string = colour.Blue("0.0.1")

const helpString string = `t - The minimalist command line todo manager

Usage: t [option] <Input>

Options:
  -l, --list                    List the tasks in the todo list
  -a, --add <Task>              Add a task to the todo list
  -d, --delete <TaskID>         Delete a task from the todo list
  -v, --version                 Show the version
  -h, --help                    Show this help information`

func main() {
	config.InitialiseConfig()

	if len(os.Args) <= 1 {
		os.Args = append(os.Args, "l")
	}

	var input = os.Args[2:]

	var flag string = strings.TrimPrefix(os.Args[1], "-")

	// Just to trim for long form flags
	flag = strings.TrimPrefix(flag, "-")

	switch flag {
	case "l", "list":
		commands.ListItems()
	case "a", "add":
		commands.AddItem(input)
	case "d", "delete":
		commands.DeleteItem(input)
	case "h", "help":
		showHelp()
	case "v", "version":
		showVersion()
	}
}

func showHelp() {
	fmt.Println(helpString)
}

func showVersion() {
	fmt.Println("Version:", versionString)
}
