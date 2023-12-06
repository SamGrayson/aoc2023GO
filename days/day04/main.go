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
	fmt.Println(Part01())
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

	// Regex for grabbing digits
	d := regexp.MustCompile("\\d+")

	for _, v := range inputArr {
		val := 0

		split := strings.Split(v, "|")
		winningValues := util.SliceToMap(d.FindAllString(split[0], -1))
		myValues := util.SliceToMap(d.FindAllString(split[1], -1))

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

	// Remove game:
	m := regexp.MustCompile(`Card \\d+:`)
	dataInput = m.ReplaceAllString(dataInput, "")

	inputArr := strings.Split(dataInput, "\n")

	var games []game

	// Regex for grabbing digits
	d := regexp.MustCompile(`\\d+`)

	// Generate values
	for i, v := range inputArr {
		split := strings.Split(v, "|")
		winningValues := util.SliceToMap(d.FindAllString(split[0], -1))
		myValues := util.SliceToMap(d.FindAllString(split[1], -1))

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

	cardSum := funk.Reduce(games, func(acc int, g game) int {
		return acc + g.copies
	}, 0)

	return cardSum.(int)
}
