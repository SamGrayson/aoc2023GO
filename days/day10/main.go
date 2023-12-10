package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

func main() {
	Part01()
	// Part02()
}

func Part01() {
	dataInput, err := util.GetInput("02")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	matrix := [][]string{}
	// row, col
	startingPoint := [2]int{}

	for i := 0; i < len(inputArr); i++ {
		matrix[i] = []string{}
		for j := 0; j < len(inputArr[i]); j++ {
			if slices.Contains(inputArr, "S") {
				startingPoint = [2]int{i, j}
			}
			matrix[i][j] = inputArr[i]
		}
	}

	funk.All(true)

	fmt.Println(inputArr)
}

func Part02() {
	dataInput, err := util.GetInput("02")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)

	fmt.Println(inputArr)
}
