package main

import (
	"fmt"
	"math"
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

type part struct {
	testArr           []string
	lockedIdx         map[int]bool
	springSplit       []int
	validCombinations *[]string
}

func Part02() int {
	dataInput, err := util.GetInput("12")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	var recur func(s []string, start, end int, combos *map[string]bool, lockedIdx map[int]bool, check []int, validCombinations *[]string)
	recur = func(s []string, start, end int, combos *map[string]bool, lockedIdx map[int]bool, check []int, validCombinations *[]string) {
		stringJoin := strings.Join(s, "")
		if (*combos)[stringJoin] {
			return
		}
		if start == end {
			(*combos)[stringJoin] = true
			if checkVilidity(s, check) {
				*validCombinations = append(*validCombinations, stringJoin)
			}
			return
		}
		for i := range s {
			// Don't swap "locked" indexes
			if !lockedIdx[i] && !lockedIdx[start] {
				s[start], s[i] = s[i], s[start]
				recur(s, start+1, end, combos, lockedIdx, check, validCombinations)
				s[start], s[i] = s[i], s[start] // flip it back
			} else {
				// I want to make it all the way to the end
				recur(s, start+1, end, combos, lockedIdx, check, validCombinations)
			}
		}
	}

	finalValue := 0
	finalValues := []int{}

	for i, v := range inputArr {
		fmt.Println("Currently Processing Input: ", i)
		var combos = map[string]bool{}
		inputSplit := strings.Split(v, " ")
		testString := inputSplit[0]
		testStringSplit := strings.Split(testString, "")

		// Test 1 then Test 2
		// Test 2 result / Test 1 result = future
		// Test 1 result * (future * 3)
		part1 := part{
			testArr:           testStringSplit,
			lockedIdx:         map[int]bool{},
			validCombinations: &[]string{},
		}

		part2End := append([]string{"?"}, testStringSplit...)
		part2 := part{
			testArr:           append(testStringSplit, part2End...),
			lockedIdx:         map[int]bool{},
			validCombinations: &[]string{},
		}
		var parts = []part{part1, part2}
		for i, part := range parts {
			// Set "locked" indexes (those that were already set)
			springs := inputSplit[1]

			// Get the springs to test. (convert to []int)
			springSplit := funk.Map(strings.Split(springs, ","), func(s string) int {
				return util.ToInt(s)
			}).([]int)

			if i == 0 {
				part.springSplit = springSplit
			} else if i == 1 {
				part.springSplit = append(springSplit, springSplit...)
			}

			for i, v := range part.testArr {
				if v != "?" {
					part.lockedIdx[i] = true
				}
			}

			// Get total springs and replace the string with the appropriate number of #s
			springTotal := funk.Reduce(part.springSplit, func(acc int, s int) int {
				acc += s
				return acc
			}, 0).(int)

			var springAdd int
			if i == 1 {
				springAdd = springTotal - (strings.Count(testString, "#") * 2)
			} else {
				// How many springs do we want to replace where there are ?
				springAdd = springTotal - strings.Count(testString, "#")
			}

			// Replace ? with required #
			for springAdd > 0 {
				for i, v := range part.testArr {
					if v == "?" {
						part.testArr[i] = "#"
						springAdd--
						break
					}
				}
			}

			// Replace remaining ? with "."
			for i, v := range part.testArr {
				if v == "?" {
					part.testArr[i] = "."
				}
			}

			// gather string permutations
			recur(part.testArr, 0, len(part.testArr)-1, &combos, part.lockedIdx, part.springSplit, part.validCombinations)

			// Get the results
			if i == 0 {
				// Add to the final result
				finalValue += len(*part.validCombinations)
			}
			if i == 1 {
				// More math we need to do
				validLen := len(*part.validCombinations)
				multi := float64(validLen / finalValue)
				right := int(math.Pow(multi, 4))
				finalValues = append(finalValues, (finalValue * right))
				finalValue = 0
			}
		}
	}

	// Return the sum of valid combinations
	return funk.SumInt(finalValues)
}

func main() {
	// fmt.Println("Valid combinations p1: ", Part01())
	fmt.Println("Valid combinations p2: ", Part02())
}
