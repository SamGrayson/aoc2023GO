package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

func checkVilidity(test []string, hashList []int) bool {
	stringJoin := strings.Join(test, "")

	re := regexp.MustCompile(`#+`)
	matches := re.FindAllString(stringJoin, -1)

	groupsInt := funk.Map(matches, func(str string) int {
		return len(str)
	}).([]int)

	// Sort - actually don't, must be in order
	// slices.Sort(hashList)
	// slices.Sort(groupsInt)

	return reflect.DeepEqual(hashList, groupsInt)
}

func Part01() int {
	dataInput, err := util.GetInput("12")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	validCombinations := []string{}

	var recur func(s []string, start, end int, combos *map[string]bool, lockedIdx map[int]bool, check []int)
	recur = func(s []string, start, end int, combos *map[string]bool, lockedIdx map[int]bool, check []int) {
		stringJoin := strings.Join(s, "")
		if (*combos)[stringJoin] {
			return
		}
		if start == end {
			(*combos)[stringJoin] = true
			if checkVilidity(s, check) {
				validCombinations = append(validCombinations, stringJoin)
			}
			return
		}
		for i := range s {
			// Don't swap "locked" indexes
			if !lockedIdx[i] && !lockedIdx[start] {
				s[start], s[i] = s[i], s[start]
				recur(s, start+1, end, combos, lockedIdx, check)
				s[start], s[i] = s[i], s[start] // flip it back
			} else {
				// I want to make it all the way to the end
				recur(s, start+1, end, combos, lockedIdx, check)
			}
		}
	}

	for _, v := range inputArr {
		var combos = map[string]bool{}
		inputSplit := strings.Split(v, " ")
		testString := inputSplit[0]
		springs := inputSplit[1]

		// Get the springs to test. (convert to []int)
		springSplit := funk.Map(strings.Split(springs, ","), func(s string) int {
			return util.ToInt(s)
		}).([]int)
		// Get total springs and replace the string with the appropriate number of #s
		springTotal := funk.Reduce(springSplit, func(acc int, s int) int {
			acc += s
			return acc
		}, 0).(int)
		// How many springs do we want to replace where there are ?
		springAdd := springTotal - strings.Count(testString, "#")
		testStringSplit := strings.Split(testString, "")

		// Set "locked" indexes (those that were already set)
		var lockedIdx = map[int]bool{}
		for i, v := range testStringSplit {
			if v != "?" {
				lockedIdx[i] = true
			}
		}

		// Replace ? with required #
		for springAdd > 0 {
			for i, v := range testStringSplit {
				if v == "?" {
					testStringSplit[i] = "#"
					springAdd--
					break
				}
			}
		}

		// Replace remaining ? with "."
		for i, v := range testStringSplit {
			if v == "?" {
				testStringSplit[i] = "."
			}
		}

		// gather string permutations
		recur(testStringSplit, 0, len(testStringSplit)-1, &combos, lockedIdx, springSplit)
	}

	// Return the sum of valid combinations
	return len(validCombinations)
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
	fmt.Println("Valid combinations p1: ", Part01())
	// Part02()
}
