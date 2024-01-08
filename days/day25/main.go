package main

import (
	"fmt"
	"os"
	"strings"
	"util"
)

func main() {
	Part01()
}

// JUST USED AN ONLINE GRAPH TOOL - :D
func Part01() {
	dataInput, err := util.GetInput("25")
	if err != nil {
		os.Exit(1)
	}

	relativePath, _ := util.GetRelativePath("25", "output.dot")

	file, err := os.Create(relativePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write top
	file.WriteString("graph { \n")

	inputArr := strings.Split(dataInput, "\n")

	for _, line := range inputArr {
		line = strings.Replace(line, ": ", "--", -1)
		line = strings.Join(strings.Split(line, " "), ",")
		file.WriteString(line + "\n")
	}

	// Write bot
	file.WriteString("}")
}
