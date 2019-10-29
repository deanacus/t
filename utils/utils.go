package utils

import (
	"bufio"
	"fmt"
	"os"
)

func handleMissingFile(filePath string) bool {
	fmt.Println(filePath, "does not exist. Should we create it for you? y/n")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	switch char {
	case 'y':
		os.Create(filePath)
		fmt.Println("File created")
		return true
	case 'n':
		fmt.Println("No file created")
		os.Exit(0)
	default:
		handleMissingFile(filePath)
	}

	return false
}

// CheckFile Tries to open a file for reading, and returns each line
// as a value in a slice, or asks the use to create the file
func CheckFile(filePath string) []string {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		handleMissingFile(filePath)
	}

	scanner := bufio.NewScanner(file)

	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result

}

// WriteTasks takes in a filePath and a slice of tasks,
// tries to open the supplied filePath, calls handleMissingFile
// if unsuccessful asks the user to crate the file, otherwise it
// opens the supplied filePath, empties it's contents
// and writes each task in the slice to the file before
// returning true on success.
func WriteTasks(filePath string, tasks []string) bool {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0700)
	defer file.Close()
	file.Truncate(0)

	if err != nil {
		handleMissingFile(filePath)
		return false
	}

	for _, line := range tasks {
		file.WriteString(line + "\n")
	}
	file.Sync()
	return true
}

// Find taks a slice and a string, and iterates
// over the slice, and returns true if it finds
// a match and false if it doesn't
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Filter taks a slice and a function that returns a boolean,
// then iterates over the slice running the function on each
// item, and pushes any that return true onto a new slice,
// returning the new slice when finished
func Filter(slice []string, fun func(int) bool) []string {
	filtered := make([]string, 0)

	for ind, item := range slice {
		if fun(ind) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
