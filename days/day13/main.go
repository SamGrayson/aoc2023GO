package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"util"
)

type matrix struct {
	rows []string
	cols []string
}

func match(input []string) ([]int, bool) {
	// detect col
	i := 0
	j := len(input) - 1
	matches := []int{}

	matchFound := false
	success := false
	for j > i {
		if !matchFound {
			// Start from left, check till equal
			for s := i; s < j; s++ {
				if input[s] == input[j] {
					// Not even, not yet a mirror - continue
					if (j-s)%2 == 0 {
						continue
					}
					i = s
					matchFound = true
					break
				}
			}
		}

		if !matchFound {
			// Start from right, check till equal
			for s := j; s > i; s-- {
				if input[s] == input[i] {
					// Not even, not yet a mirror - continue
					if (j-s)%2 == 0 {
						break
					}
					j = s
					matchFound = true
					break
				}
			}
		}

		if matchFound && input[i] == input[j] {
			matches = append(matches, []int{i, j}...)
			j--
			i++
			continue
		} else {
			break
		}
	}

	if j < i {
		success = true
	}
	return matches, success
}

func Part01() int {
	dataInput, err := util.GetInput("13")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	matrixList := []matrix{}

	colLen := len(inputArr[0])
	newMatrix := matrix{
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
			newMatrix = matrix{
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
	for i, matrix := range matrixList {

		matchedCols, colSuccess := match(matrix.cols)
		matchedRows, rowSuccess := match(matrix.rows)

		slices.Sort(matchedCols)
		slices.Sort(matchedRows)

		colVal := 0
		rowVal := 0
		// Matched cols started on the edge, its a mirror

		if colSuccess && len(matchedCols) > 0 && (matchedCols[len(matchedCols)-1] == len(matrix.cols)-1 || matchedCols[0] == 0) {
			mid := len(matchedCols) / 2
			// Left of the right side of the mirror
			colVal = matchedCols[int(math.Round(float64(mid)))]
		} else if rowSuccess && len(matchedRows) > 0 && (matchedRows[len(matchedRows)-1] == len(matrix.rows)-1 || matchedRows[0] == 0) {
			mid := len(matchedRows) / 2
			// Top of the bottom side of the mirror
			rowVal = matchedRows[int(math.Round(float64(mid)))]
		}

		fmt.Printf("Matrix: %d cols: %s, rows: %s, cVal: %d, rVal %d \n", i, fmt.Sprint(matchedCols), fmt.Sprint(matchedRows), colVal, rowVal)

		finalColVal += colVal
		finalRowVal += (100 * rowVal)
	}
	return finalColVal + finalRowVal
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
	fmt.Println("Part1 Result :", Part01())
	// Part02()
}
