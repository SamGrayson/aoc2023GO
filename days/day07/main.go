package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

func main() {
	// Part01()
	Part02()
}

type hand struct {
	numHand  []string
	origHand string
	cardRank int
	wager    int
}

var CardValues = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"1": 1,
}

var CardValues2 = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"1": 1,
	"J": 0,
}
var CardRanks = map[string]int{
	"5":  6,
	"4":  5,
	"32": 4,
	"23": 4,
	"3":  3,
	"22": 2,
	"2":  1,
	"1":  0,
}

func isLeftLowerThanRight(leftHand hand, rightHand hand) bool {
	// If the cards equal each other - need to check the characters
	if leftHand.cardRank < rightHand.cardRank {
		return true
	}

	if leftHand.cardRank == rightHand.cardRank {
		for i := 0; i < 5; i++ {
			leftNum, _ := strconv.Atoi(leftHand.numHand[i])
			rightNum, _ := strconv.Atoi(rightHand.numHand[i])
			if leftNum == rightNum {
				continue
			}
			return leftNum < rightNum
		}
	}

	return false
}

func Part01() int {
	dataInput, err := util.GetInput("07")
	if err != nil {
		os.Exit(1)
	}

	var setupHand = func(incHand []string) hand {
		numHand := []string{}
		matches := make(map[string][]string)
		wager, _ := strconv.Atoi(incHand[1])
		for i := 0; i < len(incHand[0]); i++ {
			numStr := fmt.Sprintf("%d", CardValues[string(incHand[0][i])])
			numHand = append(numHand, numStr)
			matches[numStr] = append(matches[numStr], numStr)
		}

		cardRankKey := ""
		for _, v := range matches {
			if len(v) > 1 && strings.Count(cardRankKey, fmt.Sprintf("%d", len(v))) < 2 {
				cardRankKey += fmt.Sprintf("%d", len(v))
			}
		}
		if len(cardRankKey) == 0 {
			cardRankKey = "1"
		}

		return hand{
			numHand:  numHand,
			cardRank: CardRanks[cardRankKey],
			origHand: incHand[0],
			wager:    wager,
		}
	}

	inputArr := strings.Split(dataInput, "\n")
	hands := funk.Map(inputArr, func(s string) hand {
		hand := strings.Split(s, " ")
		fmtHand := setupHand(hand)
		return fmtHand
	}).([]hand)

	handRankingTracker := hands
	// Sort left -> right
	for i := 1; i < len(hands); i++ {
		key := hands[i]
		j := i - 1

		for j >= 0 && isLeftLowerThanRight(hands[j], key) {
			handRankingTracker[j+1] = handRankingTracker[j]
			j = j - 1
		}
		handRankingTracker[j+1] = key
	}

	i := len(hands)
	result := funk.Reduce(hands, func(acc int, h hand) int {
		winnings := h.wager * i
		i--
		return acc + winnings
	}, 0)

	funk.All(true)

	fmt.Println(result)
	return result.(int)
}

func Part02() int {
	dataInput, err := util.GetInput("07")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	var getMaxNumHand = func(jokerHand []string) []string {
		var maxVal []string
		nonJokers := funk.Reduce(jokerHand, func(acc []string, s string) []string {
			if s != "J" {
				acc = append(acc, s)
			}
			return acc
		}, make([]string, 0)).([]string)
		for _, v := range nonJokers {
			fmt.Println(v)
		}
		return maxVal
	}

	var setupHand = func(incHand []string) hand {
		numHand := []string{}
		matches := make(map[string][]string)
		wager, _ := strconv.Atoi(incHand[1])
		for i := 0; i < len(incHand[0]); i++ {
			numStr := fmt.Sprintf("%d", CardValues2[string(incHand[0][i])])
			numHand = append(numHand, numStr)
			matches[numStr] = append(matches[numStr], numStr)
		}

		if funk.Contains(numHand, "0") {
			maxHand := getMaxNumHand(numHand)
			fmt.Println(maxHand)
		}

		cardRankKey := ""
		for _, v := range matches {
			if len(v) > 1 && strings.Count(cardRankKey, fmt.Sprintf("%d", len(v))) < 2 {
				cardRankKey += fmt.Sprintf("%d", len(v))
			}
		}
		if len(cardRankKey) == 0 {
			cardRankKey = "1"
		}

		return hand{
			numHand:  numHand,
			cardRank: CardRanks[cardRankKey],
			origHand: incHand[0],
			wager:    wager,
		}
	}

	hands := funk.Map(inputArr, func(s string) hand {
		hand := strings.Split(s, " ")
		fmtHand := setupHand(hand)
		return fmtHand
	}).([]hand)

	fmt.Println(hands)

	return 1
}
