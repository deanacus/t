package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Config is the main config shape for our app
type Config struct {
	Path     string
	FileName string
	FilePath string
}

// DefaultConfig is our config instance used in the app
var DefaultConfig = Config{
	Path:     home + "/.t/",
	FileName: "todo.txt",
}

var home, _ = os.UserHomeDir()

// InitialiseConfig will check for the existance of a config
// file in the users home directory and override the values
// on the defaultConfig variable if one is found
func InitialiseConfig() {
	// file, err := os.Open(home + ".t.yml")
	file, err := os.Open(".t.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&DefaultConfig)

	if err != nil {
		fmt.Println("Error loading config file")
		os.Exit(2)
	}

	if DefaultConfig.Path[1] == 47 {
		DefaultConfig.Path = strings.Replace(DefaultConfig.Path, "~", home, -1)
	}

	DefaultConfig.FilePath = DefaultConfig.Path + DefaultConfig.FileName
}
