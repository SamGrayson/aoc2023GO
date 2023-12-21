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

func goUp(above string, curr string) (string, string) {
	// skip first row
	if above == "" {
		return "", curr
	}
	for i, r := range curr {
		if r == '#' || r == '.' || above[i] == '#' || above[i] == 'O' {
			continue
		} else {
			above = util.ReplaceAtIdx(above, r, i)
			curr = util.ReplaceAtIdx(curr, '.', i)
		}
	}
	return above, curr
}

func goDown(below string, curr string) (string, string) {
	// skip first row
	if below == "" {
		return "", curr
	}
	for i, r := range curr {
		i = len(curr) - 1 - i
		if r == '#' || r == '.' || below[i] == '#' || below[i] == 'O' {
			continue
		} else {
			below = util.ReplaceAtIdx(below, r, i)
			curr = util.ReplaceAtIdx(curr, '.', i)
		}
	}
	return below, curr
}

func goLeft(line string) string {
	// compare left to right
	for i := len(line) - 2; i >= 0; i-- {
		r := line[i]
		if r == '#' || r == '.' || line[i+1] == '#' || line[i+1] == 'O' {
			continue
		}
		line = util.ReplaceAtIdx(line, rune(line[i+1]), i)
	}
	return line
}

func goRight(line string) string {
	// compare left to right
	for i := 1; i <= len(line); i++ {
		r := line[i]
		if r == '#' || r == '.' || line[i-1] == '#' || line[i-1] == 'O' {
			continue
		}
		line = util.ReplaceAtIdx(line, rune(line[i-1]), i)
	}
	return line
}

func Part02() {
	dataInput, err := util.GetInput("14")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)

	cycles := 1000000000
	// north, then west, then south, then east
	for cycles > 0 {
		for i := 0; i < len(inputArr); i++ {
			// North
			if i > 0 {
				inputArr[i-1], inputArr[i] = goUp(inputArr[i-1], inputArr[i])
			}
		}
		for i := 0; i < len(inputArr); i++ {
			// West
			inputArr[i] = goLeft(inputArr[i])
		}
		for i := 0; i < len(inputArr); i++ {
			// South
			if i < len(inputArr) {
				inputArr[i-1], inputArr[i] = goDown(inputArr[i+1], inputArr[i])
			}
		}
		for i := 0; i < len(inputArr); i++ {
			// West
			inputArr[i] = goRight(inputArr[i])
		}
	}

	fmt.Println(inputArr)
}

func main() {
	// Part01()
	Part02()
}
