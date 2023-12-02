package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"util"
)

func main() {
	// fmt.Println("Sum of IDs p1: ", Part01())
	fmt.Println("Game Power p2: ", Part02())
}

func isColor(c string) bool {
	switch c {
	case
		"red",
		"green",
		"blue":
		return true
	default:
		return false
	}
}

func Part01() int {
	dataInput, err := util.GetInput("02")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)
	badGames := map[string]bool{}
	goodGames := map[string]bool{}

	LIMITS := map[string]int{
		"red": 12, "green": 13, "blue": 14,
	}

	var currentGame string
	for i, v := range inputArr {
		// Set current game if found
		if strings.Contains(v, ":") {
			currentGame = strings.Split(v, ":")[0]
			goodGames[currentGame] = true
		}

		cleanV := strings.Replace(v, ";", "", -1)
		cleanV = strings.Replace(cleanV, ",", "", -1)

		var prevNum int
		if isColor(cleanV) {
			prevNum, _ = strconv.Atoi(inputArr[i-1])
		}
		if LIMITS[cleanV] < prevNum {
			badGames[currentGame] = true
		}
	}

	for k := range badGames {
		delete(goodGames, k)
	}

	idSum := 0
	for k := range goodGames {
		var intV, _ = strconv.Atoi(k)
		idSum += intV
	}

	return idSum
}

func Part02() int {
	dataInput, err := util.GetInput("02")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)
	var maxMap = map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	gamePower := 0

	for i, v := range inputArr {
		// Set current game if found
		if strings.Contains(v, ":") {
			gamePower = gamePower + (maxMap["red"] * maxMap["blue"] * maxMap["green"])
			maxMap["red"] = 0
			maxMap["blue"] = 0
			maxMap["green"] = 0
		}

		cleanV := strings.Replace(v, ";", "", -1)
		cleanV = strings.Replace(cleanV, ",", "", -1)

		var prevNum int
		if isColor(cleanV) {
			prevNum, _ = strconv.Atoi(inputArr[i-1])
			if maxMap[cleanV] < prevNum {
				maxMap[cleanV] = prevNum
			}
		}
	}

	// Add the last amount
	gamePower = gamePower + (maxMap["red"] * maxMap["blue"] * maxMap["green"])

	return gamePower
}
