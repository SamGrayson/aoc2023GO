package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

func main() {
	fmt.Println("Successes p1: ", Part01())
	fmt.Println("Successes p2: ", Part02())
}

func Part01() int {
	dataInput, err := util.GetInput("06")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")
	digitRe := regexp.MustCompile(`\d+`)

	times := digitRe.FindAllString(inputArr[0], -1)
	records := digitRe.FindAllString(inputArr[1], -1)

	var attempts [][2]int
	for i, v := range times {
		time, _ := strconv.Atoi(v)
		record, _ := strconv.Atoi(records[i])
		var joined = [2]int{time, record}
		attempts = append(attempts, joined)
	}

	var successes []int
	for _, attempt := range attempts {
		time := attempt[0]
		record := attempt[1]
		successCount := 0
		for i := 0; i < attempt[0]; i++ {
			currentDistance := i * (time - i)
			if currentDistance > record {
				successCount++
			}
		}
		successes = append(successes, successCount)
	}

	total := funk.Reduce(successes, func(acc int, i int) int {
		return acc * i
	}, 1)

	return total.(int)
}

func Part02() int {
	dataInput, err := util.GetInput("06")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")
	digitRe := regexp.MustCompile(`\d+`)

	time := strings.Join(digitRe.FindAllString(inputArr[0], -1), "")
	record := strings.Join(digitRe.FindAllString(inputArr[1], -1), "")

	attempt := [2]string{time, record}

	attemptTime, _ := strconv.Atoi(attempt[0])
	attemptRecord, _ := strconv.Atoi(attempt[1])
	successCount := 0

	for i := 0; i < attemptTime; i++ {
		currentDistance := i * (attemptTime - i)
		if currentDistance > attemptRecord {
			successCount++
		}
	}

	return successCount
}
