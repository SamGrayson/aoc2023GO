package main

import (
	"fmt"
	"os"
	"strings"
	"util"
)

func main() {
	// fmt.Println(Part01())
	fmt.Println(Part02())
}

func Part01() int {
	dataInput, err := util.GetInput("09")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	var recur func(prev []int) int
	recur = func(prev []int) int {
		tracker := []int{}
		for i := 1; i < len(prev); i++ {
			diff := prev[i] - prev[i-1]
			tracker = append(tracker, diff)
		}

		allZeros := true
		for i := 0; i < len(tracker); i++ {
			if tracker[i] != 0 {
				allZeros = false
				break
			}
		}

		if allZeros {
			return prev[len(tracker)-1]
		}

		return prev[len(prev)-1] + recur(tracker)
	}

	// Create initial nodes
	var finalValue = 0
	for _, input := range inputArr {
		split := strings.Split(input, " ")
		start := []int{}
		for i := range split {
			start = append(start, util.ToInt(split[i]))
		}
		finalValue += recur(start)
	}

	return finalValue
}

func Part02() int {
	dataInput, err := util.GetInput("09")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	var recur func(prev []int) int
	recur = func(prev []int) int {
		tracker := []int{}
		for i := 1; i < len(prev); i++ {
			diff := prev[i] - prev[i-1]
			tracker = append(tracker, diff)
		}

		allZeros := true
		for i := 0; i < len(tracker); i++ {
			if tracker[i] != 0 {
				allZeros = false
				break
			}
		}

		if allZeros {
			return prev[0] - 0
		}

		return prev[0] - recur(tracker)
	}

	// Create initial nodes
	var finalValue = 0
	for _, input := range inputArr {
		split := strings.Split(input, " ")
		start := []int{}
		for i := range split {
			start = append(start, util.ToInt(split[i]))
		}
		finalValue += recur(start)
	}

	return finalValue
}
