package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"util"
)

func Part01() int {

	var goUp = func(above string, curr string) (string, string, bool) {
		// skip first row
		if above == "" {
			return "", curr, false
		}
		rolled := false
		for i, r := range curr {
			if r == '#' || r == '.' || above[i] == '#' || above[i] == 'O' {
				continue
			} else {
				above = util.ReplaceAtIdx(above, r, i)
				curr = util.ReplaceAtIdx(curr, '.', i)
				if !rolled {
					rolled = true
				}
			}
		}
		return above, curr, rolled
	}

	dataInput, err := util.GetInput("14")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	rolling := true
	for rolling {
		rolled := false
		for i := 0; i < len(inputArr); i++ {
			if i > 0 {
				currRolled := false
				inputArr[i-1], inputArr[i], currRolled = goUp(inputArr[i-1], inputArr[i])
				if !rolled && currRolled {
					rolled = currRolled
				}
			}
		}
		if !rolled {
			break
		}
	}

	// final count
	finalCount := 0
	slices.Reverse(inputArr)
	for i := len(inputArr) - 1; i >= 0; i-- {
		for _, v := range inputArr[i] {
			if v == 'O' {
				finalCount += (i + 1)
			}
		}
	}
	fmt.Println("Part 1: ", finalCount)
	return finalCount

}

func goUp(above string, curr string) (string, string, bool) {
	// skip first row
	if above == "" {
		return "", curr, false
	}
	rolled := false
	for i, r := range curr {
		if r == '#' || r == '.' || above[i] == '#' || above[i] == 'O' {
			continue
		} else {
			above = util.ReplaceAtIdx(above, r, i)
			curr = util.ReplaceAtIdx(curr, '.', i)
			if !rolled {
				rolled = true
			}
		}
	}
	return above, curr, rolled
}

func goDown(below string, curr string) (string, string, bool) {
	// skip first row
	if below == "" {
		return "", curr, false
	}
	rolled := false
	for i := 0; i < len(curr)-1; i++ {
		r := curr[i]
		if r == '#' || r == '.' || below[i] == '#' || below[i] == 'O' {
			continue
		} else {
			below = util.ReplaceAtIdx(below, rune(r), i)
			curr = util.ReplaceAtIdx(curr, '.', i)
			if !rolled {
				rolled = true
			}
		}
	}
	return below, curr, rolled
}

func goLeft(line string) string {
	newLine := []string{}
	rockCount := 0
	gap := 0
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == 'O' {
			rockCount++
		}
		if line[i] != '#' {
			gap++
		}
		if line[i] == '#' || i == 0 {
			// rocks
			rocks := util.ArrWithDefaultStr(rockCount, "O")
			gaps := util.ArrWithDefaultStr(gap-rockCount, ".")
			// Create new line
			if i != 0 {
				newLine = append(newLine, "#"+strings.Join(rocks, "")+strings.Join(gaps, ""))
			} else {
				newLine = append(newLine, strings.Join(rocks, "")+strings.Join(gaps, ""))
			}
			rockCount = 0
			gap = 0
		}
		// If the last char is "#", add it
		if line[i] == '#' && i == 0 {
			newLine = append(newLine, "#")
		}
	}
	slices.Reverse(newLine)
	return strings.Join(newLine, "")
}

func goRight(line string) string {
	newLine := []string{}
	rockCount := 0
	gap := 0
	for i := 0; i <= len(line)-1; i++ {
		if line[i] == 'O' {
			rockCount++
		}
		if line[i] != '#' {
			gap++
		}
		if line[i] == '#' || i == len(line)-1 {
			// rocks
			rocks := util.ArrWithDefaultStr(rockCount, "O")
			gaps := util.ArrWithDefaultStr(gap-rockCount, ".")
			// Create new line
			if i != len(line)-1 {
				newLine = append(newLine, strings.Join(gaps, "")+strings.Join(rocks, "")+"#")
			} else {
				newLine = append(newLine, strings.Join(gaps, "")+strings.Join(rocks, ""))
			}
			rockCount = 0
			gap = 0
		}
		// If the last char is "#", add it
		if line[i] == '#' && i == len(line)-1 {
			newLine = append(newLine, "#")
		}
	}
	return strings.Join(newLine, "")
}

func Part02() int {
	dataInput, err := util.GetInput("14")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)

	// counts := []int{}

	var first string

	cycles := 1000000000
	// north, then west, then south, then east
	for cycles > 0 {
		// Speed Track
		if cycles%1000 == 0 {
			fmt.Println("Processed: ", cycles)
		}

		rolling := true

		for rolling {
			rolled := false
			for i := 0; i < len(inputArr); i++ {
				currRolled := false
				// North
				if i > 0 {
					inputArr[i-1], inputArr[i], currRolled = goUp(inputArr[i-1], inputArr[i])
				}
				if !rolled && currRolled {
					rolled = currRolled
				}
			}
			if !rolled {
				break
			}
		}

		for i := 0; i < len(inputArr); i++ {
			inputArr[i] = goLeft(inputArr[i])
		}

		// South
		for rolling {
			rolled := false
			for i := 0; i < len(inputArr)-1; i++ {
				currRolled := false
				if i < len(inputArr) {
					inputArr[i+1], inputArr[i], currRolled = goDown(inputArr[i+1], inputArr[i])
				}
				if !rolled && currRolled {
					rolled = currRolled
				}
			}
			if !rolled {
				break
			}
		}

		// East
		for i := 0; i < len(inputArr); i++ {
			inputArr[i] = goRight(inputArr[i])
		}

		if cycles == 1000000000 {
			first = strings.Join(inputArr, "")
			fmt.Println(first)
		}

		finalCount := 0
		slices.Reverse(inputArr)
		for i := len(inputArr) - 1; i >= 0; i-- {
			for _, v := range inputArr[i] {
				if v == 'O' {
					finalCount += (i + 1)
				}
			}
		}
		counts = append(counts, finalCount)

		cycles--
	}

	fmt.Println(inputArr)

	// final count
	finalCount := 0
	slices.Reverse(inputArr)
	for i := len(inputArr) - 1; i >= 0; i-- {
		for _, v := range inputArr[i] {
			if v == 'O' {
				finalCount += (i + 1)
			}
		}
	}
	fmt.Println("Part 2: ", finalCount)
	return finalCount
}

func main() {
	// Part01()
	Part02()
}
