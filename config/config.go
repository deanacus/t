package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config is the main config shape for our app
type Config struct {
	tFilePath string
}

var home, _ = os.UserHomeDir()

// DetectConfig will check for the existance of a config
// file in the users home directory and override the values
// on the defaultConfig variable if one is found
func DetectConfig() {
	// file, err := os.Open(home + ".t.yml")
	file, err := os.Open(".t.json")
	// file, err := os.Open(home + ".trc")
	// defer file.Close()

	jsn, _ := ioutil.ReadFile(".t.json")
	s := string(jsn)
	fmt.Println(s)

	// defaultConfig := Config{tFilePath: home + "/.todo"}
	newConfig := Config{tFilePath: "o"}

	if err != nil {
		fmt.Println("no config file found")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&newConfig)

	fmt.Println(&newConfig)

	// return defaultConfig
}
