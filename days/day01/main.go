package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part1 Answer:", Part01())
	fmt.Println("Part1 Answer:", Part02())
}

func Part01() int {
	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the directory:", err)
	}
	// Read the text file
	pathToInput := filepath.Join(wd, "./days/day01/input.txt")
	fileContent, err := os.ReadFile(pathToInput)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	dataInput := string(fileContent)

	// Remove any characters from the text field with regex
	pattern := "[a-zA-Z]"
	re := regexp.MustCompile(pattern)
	cleanInput := re.ReplaceAllString(dataInput, "")

	inputArr := strings.Fields(cleanInput)

	val := 0
	for _, v := range inputArr {
		smush, _ := strconv.Atoi(string(v[0]) + string(v[len(v)-1]))
		val = val + (smush)
	}

	return val
}

func Part02() int {
	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the directory:", err)
	}
	// Read the text file
	pathToInput := filepath.Join(wd, "./days/day01/input.txt")
	fileContent, err := os.ReadFile(pathToInput)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	dataInput := string(fileContent)

	// Replace number words with actual numbers
	numPair := make(map[string]string)
	numPair["one"] = "1"
	numPair["two"] = "2"
	numPair["three"] = "3"
	numPair["four"] = "4"
	numPair["five"] = "5"
	numPair["six"] = "6"
	numPair["seven"] = "7"
	numPair["eight"] = "8"
	numPair["nine"] = "9"

	// Replace words with numbers
	cleanInput := dataInput
	for k, v := range numPair {
		safeInsertIdx := 2
		// Silly way to do it, but shouldn't have any overlaps in the middle of a number
		rep := k[:safeInsertIdx] + v + k[safeInsertIdx:]
		cleanInput = strings.ReplaceAll(cleanInput, k, rep)
	}

	// Remove any characters from the text field with regex
	pattern := "[a-zA-Z]"
	re := regexp.MustCompile(pattern)
	cleanInput = re.ReplaceAllString(cleanInput, "")

	inputArr := strings.Fields(cleanInput)

	val := 0
	for _, v := range inputArr {
		smush, _ := strconv.Atoi(string(v[0]) + string(v[len(v)-1]))
		val = val + (smush)
	}

	return val
}
