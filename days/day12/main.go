package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
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
	validCombinations []string
}

func Part02() int64 {
	dataInput, err := util.GetInput("12")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	finalValues := int64(0)

	var doit func(input string, start, end int, lockedIdx map[int]bool, springs []int, validCombos int64, memo *map[string]int64, hashLen int) int64
	doit = func(input string, start, end int, lockedIdx map[int]bool, springs []int, validCombos int64, memo *map[string]int64, hashLen int) int64 {
		check := input[start:]
		if start > 0 {
			check = input[start-1:]
		}
		memoKey := check + strconv.Itoa(start) + fmt.Sprint(springs) + strconv.Itoa(hashLen)
		if val, ok := (*memo)[memoKey]; ok {
			return val
		}

		if start > 0 && string(check[0]) == "#" {
			hashLen++
		}

		// at the end
		if start == len(input) {
			// Extra #, not valid
			if len(springs) == 0 && hashLen > 0 {
				return 0
			}

			// Check the last hash len
			if hashLen > 0 && hashLen == springs[0] {
				// Pop the first
				springs = springs[1:]
			}

			// Success!
			if len(springs) == 0 {
				return 1
			}
			return 0
		}

		if string(check[0]) == "." && hashLen > 0 && len(springs) > 0 {
			// If we had a valid amount, continue else return
			if hashLen != springs[0] {
				return 0
			}
			// Pop the first and continue
			springs = springs[1:]
			hashLen = 0
		}

		// If we're at the index, skip
		if lockedIdx[start] {
			return doit(input, start+1, end, lockedIdx, springs, validCombos, memo, hashLen)
		}

		// if we're at a question mark, replace with . & # and move on
		var dotStrRes int64 = 0
		var hashStrRes int64 = 0
		if string(input[start]) == "?" {
			dotString := input[:start] + string(".") + input[start+1:]
			dotStrRes = doit(dotString, start+1, end, lockedIdx, springs, validCombos, memo, hashLen)

			hashString := input[:start] + string("#") + input[start+1:]
			hashStrRes = doit(hashString, start+1, end, lockedIdx, springs, validCombos, memo, hashLen)
		}

		(*memo)[memoKey] = dotStrRes + hashStrRes
		return dotStrRes + hashStrRes
	}

	for _, v := range inputArr {
		inputSplit := strings.Split(v, " ")
		testString := inputSplit[0]
		springs := inputSplit[1]

		newTest := []string{testString, testString, testString, testString, testString}
		// newTest := []string{testString}
		newTestStr := strings.Join(newTest, "?")
		testStringSplit := strings.Split(newTestStr, "")

		newSpring := []string{springs, springs, springs, springs, springs}
		// newSpring := []string{springs}
		newSpringStr := strings.Join(newSpring, ",")

		// Set "locked" indexes (those that were already set)
		var lockedIdx = map[int]bool{}
		for i, v := range testStringSplit {
			if v != "?" {
				lockedIdx[i] = true
			}
		}

		springSplit := funk.Map(strings.Split(newSpringStr, ","), func(s string) int {
			return util.ToInt(s)
		}).([]int)

		var memo = make(map[string]int64, 0)
		validPoint := doit(newTestStr, 0, len(newTestStr), lockedIdx, springSplit, 0, &memo, 0)

		fmt.Println(validPoint)
		finalValues += validPoint
	}
	// Return the sum of valid combinations

	return finalValues
}

func main() {
	// fmt.Println("Valid combinations p1: ", Part01())
	fmt.Println("Valid combinations p2: ", Part02())
}
