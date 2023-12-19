package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

func checkVilidityO(test []string, hashList []int) bool {
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

func Part01O() int {
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

type partO struct {
	testArr           []string
	lockedIdx         map[int]bool
	springSplit       []int
	validCombinations []string
}

func shouldContinueO(s []string, start int, check []int) bool {
	test := s[0:int(start)]
	// If the last character in the string is a "." check the first # length.
	if start > 0 && test[len(test)-1] == "." {
		test := strings.Join(test, "")
		re := regexp.MustCompile(`#+`)
		matches := re.FindAllString(test, -1)

		if len(matches) == 0 {
			return true
		}

		groupsInt := funk.Map(matches, func(str string) int {
			return len(str)
		}).([]int)

		// If we found more # then are even in check, go ahead and return false
		if len(groupsInt) > len(check) {
			return false
		}

		return slices.Equal(groupsInt, check[0:len(groupsInt)])
	}
	return true
}

func Part02O(MAX float64) float64 {
	dataInput, err := util.GetInput("12")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	finalValue := 0
	finalValues := []float64{}

	for _, v := range inputArr {

		// go func(i int, v string, finalValue int, finalValues ) {
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
			validCombinations: []string{},
		}

		part2End := append([]string{"?"}, testStringSplit...)
		part2 := part{
			testArr:           append(testStringSplit, part2End...),
			lockedIdx:         map[int]bool{},
			validCombinations: []string{},
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

			// Trying to define inside the loop so we don't have to pass the cache around.
			var recur func(s []string, start, end int, check []int)
			recur = func(s []string, start, end int, check []int) {
				stringJoin := strings.Join(s, "")

				// Look at the # that exist as we go. If it doesn't match the first check - we can skip that thread
				if !shouldContinueO(s, start, check) {
					combos[stringJoin] = true
					return
				}

				// Skip iterations where the start is a locked index
				if part.lockedIdx[start] && start != end {
					recur(s, start+1, end, check)
					return
				}

				if combos[stringJoin] {
					return
				}
				if start == end {
					combos[stringJoin] = true
					if checkVilidity(s, check) {
						part.validCombinations = append(part.validCombinations, stringJoin)
					}
					return
				}
				for i := range s {
					// Don't swap "locked" indexes
					if !part.lockedIdx[i] && !part.lockedIdx[start] {
						s[int(start)], s[i] = s[i], s[int(start)]
						str2 := strings.Join(s, "")
						if combos[str2+fmt.Sprint(start+1)+","+fmt.Sprint(end)] {
							s[int(start)], s[i] = s[i], s[int(start)] // flip it back
							continue
						}
						combos[str2+fmt.Sprint(start+1)+","+fmt.Sprint(end)] = true

						recur(s, start+1, end, check)
						s[int(start)], s[i] = s[i], s[int(start)] // flip it back
					} else {
						// I want to make it all the way to the end
						if combos[stringJoin+fmt.Sprint(start+1)+","+fmt.Sprint(end)] {
							continue
						}
						combos[stringJoin+fmt.Sprint(start+1)+","+fmt.Sprint(end)] = true
						recur(s, start+1, end, check)
					}
				}
			}

			// gather string permutations
			recur(part.testArr, 0, len(part.testArr)-1, part.springSplit)

			// Get the results
			if i == 0 {
				// Add to the final result
				finalValue += len(part.validCombinations)
			}
			if i == 1 {
				// More math we need to do
				validLen := len(part.validCombinations)
				multi := float64(validLen / finalValue)
				right := math.Pow(multi, 4)
				finalValues = append(finalValues, (float64(finalValue) * right))
				fmt.Printf("%s, %f \n", testString, (float64(finalValue) * right))
				finalValue = 0
			}
		}
		// 	<-sem // removes an int from sem, allowing another to proceed
		// }(i, v)
	}
	// Return the sum of valid combinations

	return funk.SumFloat64(finalValues)
}

func mainOLD() {
	// fmt.Println("Valid combinations p1: ", Part01())
	fmt.Println("Valid combinations p2: ", strconv.FormatFloat(Part02O(5), 'f', -1, 64))
}
