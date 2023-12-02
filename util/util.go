package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetInput(day string) (string, error) {
		// Get current working directory
		wd, _ := os.Getwd()

		// Are we debuggin?
		debug := os.Getenv("DEBUG")

		// Set debug path
		var path string
		if (!(debug == "true")) {
			path = fmt.Sprintf("./days/day%s/input.txt", day)
		} else {
			path = "input.txt"
		}

		// Read the text file
		pathToInput := filepath.Join(wd, path)
		fileContent, err := os.ReadFile(pathToInput)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return "", err
		}
	
		return string(fileContent), nil
}