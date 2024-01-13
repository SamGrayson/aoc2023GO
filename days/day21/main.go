package main

import (
	"fmt"
	"os"
	"strings"
	"util"
)

func getNeighbors(start [2]int, matrix [][]string) [][2]int {
	neighbors := [][2]int{}
	for _, n := range util.GetNeighborsPlus() {
		row := n[0] + start[0]
		col := n[1] + start[1]
		if row >= 0 && row < len(matrix) && col < len(matrix[0]) && col >= 0 {
			if matrix[row][col] != "#" {
				neighbors = append(neighbors, [2]int{row, col})
			}
		}
	}
	return neighbors
}

func Part01(maxSteps int) {
	dataInput, err := util.GetInput("21")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	matrix := make([][]string, len(inputArr))
	startingPoint := [2]int{0, 0}

	for i := 0; i < len(inputArr); i++ {
		matrix[i] = []string{}
		split := strings.Split(inputArr[i], "")
		matrix[i] = make([]string, len(split))
		for j := 0; j < len(split); j++ {
			if split[j] == "S" && startingPoint == [2]int{0, 0} {
				startingPoint = [2]int{i, j}
			}
			matrix[i][j] = split[j]
		}
	}

	current_steps := util.Set{
		startingPoint: true,
	}

	stepCount := 1
	for stepCount <= maxSteps {
		next_steps := util.Set{}
		for step := range current_steps {
			validNeighbors := getNeighbors(step.([2]int), matrix)
			for _, n := range validNeighbors {
				next_steps.Add(n)
			}
		}
		stepCount++
		current_steps = next_steps
	}

	fmt.Println(len(current_steps))
}

func Part02() {
	dataInput, err := util.GetInput("02")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)

	fmt.Println(inputArr)
}

func main() {
	Part01(64)
	// Part02()
}
