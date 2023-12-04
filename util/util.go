package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func GetInput(day string) (string, error) {
	// Get current working directory
	wd, _ := os.Getwd()

	// Are we debuggin?
	debug := os.Getenv("DEBUG")

	// Set debug path
	var path string
	if !(debug == "true") {
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

func IsNum(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}
	return false
}

func RemoveSpaceChar(arr []string) []string {
	newArr := []string{}
	for _, v := range arr {
		if v != " " {
			newArr = append(newArr, (v))
		}
	}
	return newArr
}

func RemoveEmptyChar(arr []string) []string {
	newArr := []string{}
	for _, v := range arr {
		if v != "" {
			newArr = append(newArr, (v))
		}
	}
	return newArr
}

func SliceToMap(arr []string) map[string]bool {
	ret := make(map[string]bool)
	for i := 0; i < len(arr); i += 1 {
		ret[arr[i]] = true
	}
	return ret
}
