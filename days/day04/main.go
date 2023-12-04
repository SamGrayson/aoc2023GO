package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

func main() {
	// fmt.Println(Part01())
	fmt.Println(Part02())
}

func Part01() float64 {
	dataInput, err := util.GetInput("02")
	if err != nil {
		os.Exit(1)
	}

	// Remove game:
	m := regexp.MustCompile("Card \\d+:")
	dataInput = m.ReplaceAllString(dataInput, "")

	inputArr := strings.Split(dataInput, "\n")

	var scratchPoints []int

	for _, v := range inputArr {
		val := 0

		split := strings.Split(v, "|")
		trimmedW := strings.TrimSpace(split[0])
		trimmedMy := strings.TrimSpace(split[1])
		winningValues := util.SliceToMap(strings.Split(trimmedW, " "))
		dirtyMyValues := strings.Split(trimmedMy, " ")
		myValues := util.SliceToMap(util.RemoveEmptyChar(dirtyMyValues))

		for k := range winningValues {
			if myValues[k] {
				if val == 0 {
					val = 1
				} else {
					val *= 2
				}
			}
		}
		scratchPoints = append(scratchPoints, val)
	}

	return funk.Sum(scratchPoints)
}

type game struct {
	values        map[string]bool
	winningValues map[string]bool
	gameNum       int
	wins          int
	copies        int
}

func Part02() int {
	dataInput, err := util.GetInput("02")
	if err != nil {
		os.Exit(1)
	}

	// recursive function to track the card values used per game
	// var scratchFinder = func(arr []string, scratches []int) int {
	// 	return 1
	// }

	// Remove game:
	m := regexp.MustCompile("Card \\d+:")
	dataInput = m.ReplaceAllString(dataInput, "")

	inputArr := strings.Split(dataInput, "\n")

	var games []game

	// Generate values
	for i, v := range inputArr {
		split := strings.Split(v, "|")
		trimmedW := strings.TrimSpace(split[0])
		trimmedMy := strings.TrimSpace(split[1])
		winningValues := util.SliceToMap(strings.Split(trimmedW, " "))
		dirtyMyValues := strings.Split(trimmedMy, " ")
		myValues := util.SliceToMap(util.RemoveEmptyChar(dirtyMyValues))

		wins := 0
		for k := range winningValues {
			if myValues[k] {
				wins += 1
			}
		}

		games = append(games, game{
			values:        winningValues,
			winningValues: myValues,
			gameNum:       i + 1,
			wins:          wins,
			copies:        1,
		})
	}

	i := 0
	for i != len(games) {
		currentGame := games[i]
		// Apply wins
		for o := 0; o < currentGame.wins; o++ {
			games[currentGame.gameNum+o].copies += currentGame.copies
		}
		i += 1
	}

	cardSum := 0

	for _, v := range games {
		cardSum += v.copies
	}
	return cardSum
}
