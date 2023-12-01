package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"util"
)

func main() {
	p1, _ := Part01()
	p2, _ := Part02()
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part1 Answer:", p2)
}

func Part01() (int, error) {
	dataInput, err := util.GetInput("01", false)
	if err != nil { os.Exit(1) }

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

	return val, nil
}

func Part02() (int, error) {
	dataInput, err := util.GetInput("01", false)
	if err != nil { os.Exit(1) }

	// Create number mapping
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

	// Replace words with numbers ex: on1e tw2o
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

	return val, nil
}
