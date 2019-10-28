package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readfile() {
	HOME, err := os.UserHomeDir()
	check(err)
	if fileExists(HOME + "/.todo") {
		dat, err := ioutil.ReadFile(HOME + "/.todo")
		fmt.Println(HOME + "/.todo")
		check(err)
		fmt.Print(string(dat))
	} else {
		dat, err := os.Create("./.todo")
		check(err)
		fmt.Println("File did not exist. It has been created")
		fmt.Println(dat)
	}
}
