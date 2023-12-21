package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

type matrix struct {
	rows       []string
	cols       []string
	colMatches []int
	rowMatches []int
	colVal     int
	rowVal     int
}

func validMirror(input []string, i, j int) ([]int, bool) {
	matches := []int{}
	for j > i {
		if input[i] == input[j] {
			matches = append(matches, []int{i, j}...)
			j--
			i++
			continue
		} else {
			return matches, false
		}
	}
	return matches, true
}

// Modified to make part 2 easier
func Part01() []*matrix {
	// Part 1 match function
	match := func(input []string) ([]int, bool) {
		// detect col
		i := 0
		j := len(input) - 1

		matches := []int{}
		valid := false
		matchFound := false
		if !matchFound {
			// Start from left, check till equal
			for s := i; s < j; s++ {
				if input[s] == input[j] {
					// Not even, not yet a mirror - continue
					if (j-s)%2 == 0 {
						continue
					}
					matches, valid = validMirror(input, s, j)
					if valid {
						matchFound = true
						break
					}
				}
			}
		}

		if !matchFound {
			// Start from right, check till equal
			for s := j; s > i; s-- {
				if input[s] == input[i] {
					// Not even, not yet a mirror - continue
					if (j-s)%2 == 0 {
						continue
					}
					matches, valid = validMirror(input, i, s)
					if valid {
						matchFound = true
						break
					}
				}
			}
		}

		return matches, matchFound
	}

	dataInput, err := util.GetInput("13")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	matrixList := []*matrix{}

	colLen := len(inputArr[0])
	newMatrix := &matrix{
		rows: []string{},
		cols: make([]string, colLen),
	}
	// Loop through each matrix in the list
	for i, row := range inputArr {
		// We're at the end of the matrix, go ahead and add the prev matrix
		if row == "" || i == len(inputArr)-1 {
			// At the end just do the adding matrix rows / cols again
			if i == len(inputArr)-1 {
				// Add cols & rows
				for j, col := range row {
					newMatrix.cols[j] = newMatrix.cols[j] + string(col)
				}
				newMatrix.rows = append(newMatrix.rows, row)
				matrixList = append(matrixList, newMatrix)
				continue
			}

			matrixList = append(matrixList, newMatrix)
			newMatrix = &matrix{
				rows: []string{},
				cols: make([]string, len(inputArr[i+1])),
			}
			continue
		}
		// Add cols & rows
		for j, col := range row {
			newMatrix.cols[j] = newMatrix.cols[j] + string(col)
		}
		newMatrix.rows = append(newMatrix.rows, row)
	}

	// Track highest value
	finalColVal := 0
	finalRowVal := 0
	for _, matrix := range matrixList {

		matchedCols, colSuccess := match(matrix.cols)
		matchedRows, rowSuccess := match(matrix.rows)

		slices.Sort(matchedCols)
		slices.Sort(matchedRows)

		colVal := 0
		rowVal := 0
		// Matched cols started on the edge, its a mirror

		if colSuccess && len(matchedCols) > 0 && (matchedCols[len(matchedCols)-1] == len(matrix.cols)-1 || matchedCols[0] == 0) {
			mid := (len(matchedCols) - 1) / 2
			// Left of the right side of the mirror
			colVal = matchedCols[mid] + 1
		} else if rowSuccess && len(matchedRows) > 0 && (matchedRows[len(matchedRows)-1] == len(matrix.rows)-1 || matchedRows[0] == 0) {
			mid := (len(matchedRows) - 1) / 2
			// Top of the bottom side of the mirror
			rowVal = matchedRows[mid] + 1
		}

		matrix.colMatches = matchedCols
		matrix.rowMatches = matchedRows
		matrix.colVal = colVal
		matrix.rowVal = rowVal

		finalColVal += colVal
		finalRowVal += (100 * rowVal)
	}
	return matrixList
}

func Part02() int {
	// Part 2 match function
	// Need to find multiple valid mirrors
	match := func(input []string, originalMatch []int) ([]int, bool) {
		smugesToTry := [][]string{}

		// Find all strings with only ONE difference for every row (may be slow :/)
		// Going to do this for every combo...
		for i, s := range input {
			newInput := make([]string, len(input))
			// why empty
			copy(newInput, input)
			for j := 0; j < len(input); j++ {
				sCopy := strings.Clone(s)
				sArr := strings.Split(s, "")
				iArr := strings.Split(input[j], "")
				diffStr, diffIdx := util.Difference(sArr, iArr)
				if len(diffStr) == 1 {
					if diffStr[0] == "#" {
						sCopy = sCopy[:diffIdx[0]] + string(".") + sCopy[diffIdx[0]+1:]
					}
					if diffStr[0] == "." {
						sCopy = sCopy[:diffIdx[0]] + string("#") + sCopy[diffIdx[0]+1:]
					}
				}
				// If we fixed a smuge, add to list to try
				if sCopy != s {
					newInput[i] = sCopy
					smugesToTry = append(smugesToTry, newInput)
				}
			}
		}

		matches := []int{}

		for _, in := range smugesToTry {
			matchTry := []int{}
			valid := false
			matchFound := false

			// detect col
			i := 0
			j := len(in) - 1

			if !matchFound {
				// Start from left, check till equal
				for s := i; s < j; s++ {
					if in[s] == in[j] {
						// Not even, not yet a mirror - continue
						if (j-s)%2 == 0 {
							continue
						}
						matchTry, valid = validMirror(in, s, j)
						// Sort for comparison
						slices.Sort(matchTry)
						if valid && !funk.Equal(matchTry, originalMatch) {
							matchFound = true
							break
						}
					}
				}
			}

			if !matchFound {
				// Start from right, check till equal
				for s := j; s > i; s-- {
					if in[s] == in[i] {
						// Not even, not yet a mirror - continue
						if (j-s)%2 == 0 {
							continue
						}
						matchTry, valid = validMirror(in, i, s)
						// Sort for comparison
						slices.Sort(matchTry)
						if valid && !funk.Equal(matchTry, originalMatch) {
							matchFound = true
							break
						}
					}
				}
			}

			// If we found a mirror that didn't equal the previous ran time - the smug was removed and found a new pattern
			if valid && matchFound {
				matches = matchTry
				break
			}
		}
		return matches, len(matches) > 0
	}

	matrixList := Part01()

	// Track highest value
	finalColVal := 0
	finalRowVal := 0
	for i, matrix := range matrixList {

		matchedCols, colSuccess := match(matrix.cols, matrix.colMatches)
		matchedRows, rowSuccess := match(matrix.rows, matrix.rowMatches)

		colVal := 0
		rowVal := 0
		// Matched cols started on the edge, its a mirror

		if colSuccess && len(matchedCols) > 0 && (matchedCols[len(matchedCols)-1] == len(matrix.cols)-1 || matchedCols[0] == 0) {
			mid := (len(matchedCols) - 1) / 2
			// Left of the right side of the mirror
			colVal = matchedCols[mid] + 1
		} else if rowSuccess && len(matchedRows) > 0 && (matchedRows[len(matchedRows)-1] == len(matrix.rows)-1 || matchedRows[0] == 0) {
			mid := (len(matchedRows) - 1) / 2
			// Top of the bottom side of the mirror
			rowVal = matchedRows[mid] + 1
		}

		fmt.Printf("Matrix: %d cols: %s, rows: %s, cVal: %d, rVal %d \n", i, fmt.Sprint(matchedCols), fmt.Sprint(matchedRows), colVal, rowVal)

		finalColVal += colVal
		finalRowVal += (100 * rowVal)
	}
	return finalColVal + finalRowVal
}

func main() {
	// fmt.Println("Part1 Result :", Part01())
	fmt.Println("Part2 Result :", Part02())
}
