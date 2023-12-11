package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"util"
)

func main() {
	fmt.Println("Part 1 sum: ", Part01())
	fmt.Println("Part 1 gearTotal: ", Part02())
}

var neighbors = util.GetNeighborsSquare()

type find struct {
	value string
	found bool
	idx   []string
}

func Part01() int {
	dataInput, err := util.GetInput("03")
	if err != nil {
		os.Exit(1)
	}

	// Get a list of all the special characters in the output
	regexPattern := `[^0-9.]`
	// Compile the regular expression pattern
	regexpPattern := regexp.MustCompile(regexPattern)
	// Find all matches in the text based on the pattern
	symbols := regexpPattern.FindAllString(dataInput, -1)
	// ArrForFound
	var foundList []int

	var neighborCrawl = func(dataInput [][]string, rw []int, symbols []string) bool {
		for _, v := range neighbors {
			// OOB Checks
			if rw[0]+v[0] < 0 {
				continue
			}
			if rw[0]+v[0] > len(dataInput)-1 {
				continue
			}
			if rw[1]+v[1] < 0 {
				continue
			}
			if rw[1]+v[1] > len(dataInput[0])-1 {
				continue
			}
			neighbor := dataInput[rw[0]+v[0]][rw[1]+v[1]]
			// If a symbol is detected, return that it was - mark the number index as seen
			if slices.Contains(symbols, neighbor) {
				return true
			}
		}
		return false
	}

	inputArr := strings.Fields(dataInput)

	// turn input to 2d array
	twoDInput := make([][]string, len(inputArr))
	for i, v := range inputArr {
		twoDInput[i] = strings.Split(v, "")
	}

	// Track the current number we're checking
	curNum := find{
		value: "",
		found: false,
	}
	// Do your twoD crawling.
	for r, v := range twoDInput {
		// Get the end of the line..
		if len(curNum.value) > 0 {
			// if curNum had adjacent symbol
			if curNum.found {
				f, _ := strconv.Atoi(curNum.value)
				foundList = append(foundList, f)
			}
			curNum.value = ""
			curNum.found = false
		}
		for w, nv := range v {
			// if we found a "." or a symbol curNum needs reset
			if (nv == "." || slices.Contains(symbols, nv)) && len(curNum.value) > 0 {
				// if curNum had adjacent symbol
				if curNum.found {
					f, _ := strconv.Atoi(curNum.value)
					foundList = append(foundList, f)
				}
				curNum.value = ""
				curNum.found = false
			}
			// if it's a number, do a neighbor check
			if nv != "." && !slices.Contains(symbols, nv) {
				curNum.value = curNum.value + nv
				rw := []int{r, w}
				found := neighborCrawl(twoDInput, rw, symbols)
				if !curNum.found {
					curNum.found = found
				}
			}
		}
	}

	// Just in case there's a number..
	if len(curNum.value) > 0 {
		// if curNum had adjacent symbol
		if curNum.found {
			f, _ := strconv.Atoi(curNum.value)
			foundList = append(foundList, f)
		}
		curNum.value = ""
		curNum.found = false
	}

	// Get sum
	sum := 0
	for _, v := range foundList {
		sum += v
	}
	return sum
}

func Part02() int {
	dataInput, err := util.GetInput("03")
	if err != nil {
		os.Exit(1)
	}

	// Get a list of all the special characters in the output
	regexPattern := `[^0-9.]`
	// Compile the regular expression pattern
	regexpPattern := regexp.MustCompile(regexPattern)
	// Find all matches in the text based on the pattern
	SYMBOLS := regexpPattern.FindAllString(dataInput, -1)
	COG := "*"
	// ArrForFound
	var foundList []int
	// Map for all index & their number
	var numIdx = make(map[string]int)
	// Track the gear total
	gearTotal := 0

	var findValidCog = func(dataInput [][]string, rw []int, numIdx map[string]int) {
		var numbers = make(util.Set)
		for _, v := range neighbors {
			// OOB Checks
			if rw[0]+v[0] < 0 {
				continue
			}
			if rw[0]+v[0] > len(dataInput)-1 {
				continue
			}
			if rw[1]+v[1] < 0 {
				continue
			}
			if rw[1]+v[1] > len(dataInput[0])-1 {
				continue
			}
			nr := rw[0] + v[0]
			nw := rw[1] + v[1]
			neighbor := dataInput[rw[0]+v[0]][rw[1]+v[1]]
			if util.IsNum(neighbor) {
				nFromIdx := numIdx[fmt.Sprintf("%d,%d", nr, nw)]
				numbers.Add(nFromIdx)
			}
		}
		if len(numbers) == 2 {
			value := 1
			// Get just the numbers
			for k := range numbers {
				value *= k.(int)
			}
			gearTotal += value
		}
	}

	inputArr := strings.Fields(dataInput)

	// turn input to 2d array
	twoDInput := make([][]string, len(inputArr))
	for i, v := range inputArr {
		twoDInput[i] = strings.Split(v, "")
	}

	var setIdxMap = func(class *find) {
		for _, v := range class.idx {
			n, _ := strconv.Atoi(class.value)
			numIdx[v] = n
		}
		class.value = ""
		class.idx = []string{}
	}

	// Track the current number we're checking
	curNum := find{
		value: "",
	}
	// Crawl 2d & set number for all index.
	for r, v := range twoDInput {
		// Get end of line nums
		if len(curNum.value) > 0 {
			setIdxMap(&curNum)
		}
		// Reset curNum each line
		for w, nv := range v {
			if nv != "." && !slices.Contains(SYMBOLS, nv) {
				curNum.value = curNum.value + nv
				curNum.idx = append(curNum.idx, fmt.Sprintf("%d,%d", r, w))
			}
			// if we found a "." or a symbol curNum needs reset
			if (nv == "." || slices.Contains(SYMBOLS, nv)) && len(curNum.value) > 0 {
				setIdxMap(&curNum)
			}
		}
	}
	// In case number at very end
	if len(curNum.value) > 0 {
		setIdxMap(&curNum)
	}

	// CRAWL 2D to find cog & numbers that make sense
	for r, v := range twoDInput {
		// Reset curNum each line
		for w, nv := range v {
			if nv == COG {
				rw := []int{r, w}
				findValidCog(twoDInput, rw, numIdx)
			}
		}
	}

	foundList = append(foundList, 1)

	return gearTotal
}
