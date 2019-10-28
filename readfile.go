package main

import (
	// "bufio"
	"fmt"
	// "io"
	"io/ioutil"
	"os"
)




func check(e error) {
	if e != nil {
			panic(e)
	}
}


func readfile() {
	HOME, err := os.UserHomeDir()
	dat, err := ioutil.ReadFile(HOME + "/.todo")
	fmt.Println(HOME + "/.todo");
	check(err)
	fmt.Print(string(dat))
}