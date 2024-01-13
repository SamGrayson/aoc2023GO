package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

func floodFill(matrix *[][]int, row, col int) {
	rows := len(*matrix)
	if rows == 0 {
		return
	}
	cols := len((*matrix)[0])

	// Check if the current position is within the bounds
	if row < 0 || row >= rows || col < 0 || col >= cols || (*matrix)[row][col] == 1 {
		return
	}

	(*matrix)[row][col] = 1

	// Recursive calls for its neighbors
	floodFill(matrix, row+1, col)
	floodFill(matrix, row-1, col)
	floodFill(matrix, row, col+1)
	floodFill(matrix, row, col-1)
}

func Part01(start [2]int) {
	dataInput, err := util.GetInput("18")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	edgeCoords := [][2]int{}
	// up down left right
	directions := map[string][2]int{
		"U": {-1, 0},
		"D": {1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}
	edgeCoords = append(edgeCoords, [2]int{0, 0})
	startCoord := [2]int{0, 0}
	for _, v := range inputArr {
		// First value is always going to be 0,0
		split := strings.Split(v, " ")
		dir := split[0]
		dirCoord := directions[string(dir)]
		count := util.ToInt(split[1])
		for count > 0 {
			newCoord := [2]int{startCoord[0] + dirCoord[0], startCoord[1] + dirCoord[1]}
			startCoord = newCoord
			edgeCoords = append(edgeCoords, newCoord)
			count--
		}
	}

	minRow := funk.Reduce(edgeCoords, func(acc int, e [2]int) int {
		if e[0] < acc {
			return e[0]
		}
		return acc
	}, math.MaxInt)
	maxRow := funk.Reduce(edgeCoords, func(acc int, e [2]int) int {
		if e[0] > acc {
			return e[0]
		}
		return acc
	}, 0)

	minCol := funk.Reduce(edgeCoords, func(acc int, e [2]int) int {
		if e[1] < acc {
			return e[1]
		}
		return acc
	}, math.MaxInt)
	maxCol := funk.Reduce(edgeCoords, func(acc int, e [2]int) int {
		if e[1] > acc {
			return e[1]
		}
		return acc
	}, 0)

	matrix := [][]int{}
	for i := 0; i <= maxRow.(int)-minRow.(int); i++ {
		matrix = append(matrix, make([]int, maxCol.(int)+1-minCol.(int)))
		for j := 0; j <= maxCol.(int)-minCol.(int); j++ {
			if slices.Contains(edgeCoords, [2]int{i + minRow.(int), j + minCol.(int)}) {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}
		}
	}

	// util.PrintIntMatrix(matrix) // DEBUG
	floodFill(&matrix, start[0], start[1])

	count := 0
	for _, row := range matrix {
		for _, col := range row {
			if col == 1 {
				count++
			}
		}
	}

	fmt.Println("Part 1 count:", count)
}

func Part02() {
	dataInput, err := util.GetInput("18")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	// up down left right
	directions := map[string][2]int{
		"3": {-1, 0},
		"1": {1, 0},
		"2": {0, -1},
		"0": {0, 1},
	}
	startCoord := [2]int{0, 0}

	// sort through edge coordinates and add them to a row mapping
	rowMap := map[int][][2]int{}

	for _, v := range inputArr {
		split := strings.Split(v, " ")
		hex := split[2]
		hex = hex[2 : len(hex)-1]
		dirCode := string(hex[len(hex)-1])
		dirCoord := directions[string(dirCode)]
		countCode := hex[:len(hex)-1]
		count, _ := strconv.ParseInt(countCode, 16, 64)
		for count > 0 {
			newCoord := [2]int{startCoord[0] + dirCoord[0], startCoord[1] + dirCoord[1]}
			startCoord = newCoord
			if _, ok := rowMap[startCoord[0]+dirCoord[0]]; !ok {
				rowMap[startCoord[0]+dirCoord[0]] = [][2]int{{startCoord[0] + dirCoord[0], startCoord[1] + dirCoord[1]}}
			} else {
				rowMap[startCoord[0]+dirCoord[0]] = append(rowMap[startCoord[0]+dirCoord[0]], [2]int{startCoord[0] + dirCoord[0], startCoord[1] + dirCoord[1]})
			}
			count--
		}
	}

	var totalCount int64 = 0
	for _, row := range rowMap {
		sort.Slice(row, func(i, j int) bool {
			for x := range row[i] {
				if row[i][x] == row[j][x] {
					continue
				}
				return row[i][x] < row[j][x]
			}
			return false
		})

		for i := 1; i <= len(row)-1; i++ {
			if i%2 != 0 {
				distance := row[i][1] - row[i-1][1]
				totalCount += int64(math.Abs(float64(distance)))
			}
		}
	}

	fmt.Println("Part 2 count:", totalCount)
}

func main() {
	// Part01([2]int{1, 1}) // example
	// Part01([2]int{65, 115}) // real
	// Part02([2]int{1, 1}) // example
	Part02()
}
