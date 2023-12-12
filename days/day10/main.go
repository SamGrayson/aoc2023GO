package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"util"
)

func main() {
	// Part01()
	Part02()
}

type pipe struct {
	shape      string
	up         []string
	down       []string
	left       []string
	right      []string
	directions [][]int
}

var pipes = make(map[string]pipe)

// Helper to add pipes
func addPipe(shape string, up, down, left, right []string, directions [][]int) {
	pipes[shape] = pipe{
		shape:      shape,
		up:         up,
		down:       down,
		left:       left,
		right:      right,
		directions: directions,
	}
}

// Possible Pipe Directions
var possibleDirections = [4][]string{
	{"|", "7", "F"},
	{"|", "L", "J"},
	{"-", "F", "L"},
	{"-", "J", "7"}}

// Create Pipe Dictionary
func createPipes() {
	/*
		-- | is a vertical pipe connecting north and south.
		-- - is a horizontal pipe connecting east and west.
		-- L is a 90-degree bend connecting north and east.
		-- J is a 90-degree bend connecting north and west.
		-- 7 is a 90-degree bend connecting south and west.
		-- F is a 90-degree bend connecting south and east.
		-- . is ground; there is no pipe in this tile.
		-- S is the starting position of the animal.
	*/
	// shape up down left right {-1, 0}, {1, 0} /*{0,0}, */, {0, -1}, {0, 1},
	addPipe("|", possibleDirections[0], possibleDirections[1], nil, nil, [][]int{{-1, 0}, {1, 0}})
	addPipe("-", nil, nil, possibleDirections[2], possibleDirections[3], [][]int{{0, -1}, {0, 1}})
	addPipe("L", possibleDirections[0], nil, nil, possibleDirections[3], [][]int{{-1, 0}, {0, 1}})
	addPipe("J", possibleDirections[0], nil, possibleDirections[2], nil, [][]int{{-1, 0}, {0, -1}})
	addPipe("7", nil, possibleDirections[1], possibleDirections[2], nil, [][]int{{1, 0}, {0, -1}})
	addPipe("F", nil, possibleDirections[1], nil, possibleDirections[3], [][]int{{1, 0}, {0, 1}})
	addPipe(".", nil, nil, nil, nil, [][]int{})
}

type startingPipe struct {
	coord [2]int
	pipe  pipe
}

// Get the first two neighbors..
func getStaringNeighborPipes(startingPoint [2]int, matrix [][]string, pipes map[string]pipe) []startingPipe {
	var neighbors = []startingPipe{}
	// Util returns in up down left right
	for i, n := range util.GetNeighborsPlus() {
		row := n[0] + startingPoint[0]
		col := n[1] + startingPoint[1]
		// Are we in bounds?
		if row >= 0 && row <= len(matrix) && col <= len(matrix[0]) && col >= 0 {
			// If the direction & the pipe makes sense..
			if slices.Contains(possibleDirections[i], matrix[row][col]) {
				neighbors = append(neighbors, startingPipe{
					pipe:  pipes[matrix[row][col]],
					coord: [2]int{row, col},
				})
			}
		}
	}
	return neighbors
}

// Get the next pipe (can only be uni-directional)
func getNextPipe(startingPoint, prev [2]int, matrix [][]string, pipes map[string]pipe) (startingPipe, bool) {
	// Util returns in up down left right
	for _, n := range pipes[matrix[startingPoint[0]][startingPoint[1]]].directions {
		row := n[0] + startingPoint[0]
		col := n[1] + startingPoint[1]
		// Skip the previous "valid" pipe
		if [2]int{row, col} == prev {
			continue
		}
		// Are we in bounds?
		if row >= 0 && row < len(matrix) && col < len(matrix[0]) && col >= 0 {
			// If the direction & the pipe makes sense..
			return startingPipe{
				pipe:  pipes[matrix[row][col]],
				coord: [2]int{row, col},
			}, true
		}
	}
	return startingPipe{}, false
}

func Part01() int {
	dataInput, err := util.GetInput("10")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	// Create Pipes Dictionary
	createPipes()

	matrix := make([][]string, len(inputArr))
	// row, col
	var startingPoint [2]int

	for i := 0; i < len(inputArr); i++ {
		matrix[i] = []string{}
		split := strings.Split(inputArr[i], "")
		matrix[i] = make([]string, len(split))
		for j := 0; j < len(split); j++ {
			if split[j] == "S" && startingPoint == [2]int{0, 0} {
				startingPoint = [2]int{i, j}
			}
			matrix[i][j] = split[j]
		}
	}

	// determine starting neighbors
	startingPipes := getStaringNeighborPipes(startingPoint, matrix, pipes)

	maxFound := false
	currentPipeOne := startingPipes[0]
	onePrev := startingPoint
	currentPipeTwo := startingPipes[1]
	twoPrev := startingPoint
	distance := 1
	for !maxFound {
		distance++
		// Start one
		resultOne, _ := getNextPipe(currentPipeOne.coord, onePrev, matrix, pipes)
		onePrev = currentPipeOne.coord
		currentPipeOne = resultOne

		// Start two
		resultTwo, _ := getNextPipe(currentPipeTwo.coord, twoPrev, matrix, pipes)
		twoPrev = currentPipeTwo.coord
		currentPipeTwo = resultTwo

		if currentPipeTwo.coord == currentPipeOne.coord {
			maxFound = true
		}
	}

	fmt.Println(distance)
	return distance
}

func expandMatrix(matrix [][]string, pipes map[string]pipe) [][]string {
	expanded := make([][]string, len(matrix)*2)
	for i := 0; i < len(expanded); i++ {
		expanded[i] = make([]string, len(matrix[0])*2)
	}

	// Space out.
	expandedI := 0
	for _, row := range matrix {
		expandedJ := 0
		for j, v := range row {
			expanded[expandedI][j+expandedJ] = v
			for _, dir := range util.GetNeighborsSquare() {
				newX, newY := dir[0], dir[1]
				if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) {
					expanded[expandedI+dir[0]][expandedJ+dir[1]+expandedJ] = " "
				}
			}
			expandedJ = 1 + j
		}
		expandedI += 2
	}

	// make another copy - gross
	var expandedPipes = [][]string{}
	for i := 0; i < len(expanded); i++ {
		exRow := append([]string(nil), expanded[i]...)
		expandedPipes = append(expandedPipes, exRow)
	}

	// Add pipes back in
	for i, row := range expanded {
		for j, v := range row {
			if v == "-" {
				for _, dir := range pipes["-"].directions {
					newX, newY := i+dir[0], j+dir[1]
					if newX >= 0 && newX < len(expandedPipes) && newY >= 0 && newY < len(expandedPipes[0]) {
						expandedPipes[newX][newY] = "-"
					}
				}
			}
			if v == "|" {
				for _, dir := range pipes["|"].directions {
					newX, newY := i+dir[0], j+dir[1]
					if newX >= 0 && newX < len(expandedPipes) && newY >= 0 && newY < len(expandedPipes[0]) {
						expandedPipes[newX][newY] = "|"
					}
				}
			}
			if v == "J" {
				for d, dir := range pipes["J"].directions {
					newX, newY := i+dir[0], j+dir[1]
					if newX >= 0 && newX < len(expandedPipes) && newY >= 0 && newY < len(expandedPipes[0]) {
						if d == 0 {
							expandedPipes[newX][newY] = "|"
						}
						if d == 1 {
							expandedPipes[newX][newY] = "-"
						}
					}
				}
			}
			if v == "F" {
				for d, dir := range pipes["F"].directions {
					newX, newY := i+dir[0], j+dir[1]
					if newX >= 0 && newX < len(expandedPipes) && newY >= 0 && newY < len(expandedPipes[0]) {
						if d == 0 {
							expandedPipes[newX][newY] = "|"
						}
						if d == 1 {
							expandedPipes[newX][newY] = "-"
						}
					}
				}
			}
			if v == "7" {
				for d, dir := range pipes["7"].directions {
					newX, newY := i+dir[0], j+dir[1]
					if newX >= 0 && newX < len(expandedPipes) && newY >= 0 && newY < len(expandedPipes[0]) {
						if d == 0 {
							expandedPipes[newX][newY] = "|"
						}
						if d == 1 {
							expandedPipes[newX][newY] = "-"
						}
					}
				}
			}
			if v == "L" {
				for d, dir := range pipes["L"].directions {
					newX, newY := i+dir[0], j+dir[1]
					if newX >= 0 && newX < len(expandedPipes) && newY >= 0 && newY < len(expandedPipes[0]) {
						if d == 0 {
							expandedPipes[newX][newY] = "|"
						}
						if d == 1 {
							expandedPipes[newX][newY] = "-"
						}
					}
				}
			}
		}
	}
	return expandedPipes
}

// Idea: Add the pipe circle to a new matrix, any fields that are between two pipes in the same row are "in the circle".
func Part02() int {
	dataInput, err := util.GetInput("10")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	// Create Pipes Dictionary
	createPipes()

	matrix := make([][]string, len(inputArr))

	// row, col
	var startingPoint [2]int

	for i := 0; i < len(inputArr); i++ {
		matrix[i] = []string{}
		split := strings.Split(inputArr[i], "")
		matrix[i] = make([]string, len(split))
		for j := 0; j < len(split); j++ {
			if split[j] == "S" && startingPoint == [2]int{0, 0} {
				startingPoint = [2]int{i, j}
			}
			matrix[i][j] = split[j]
		}
	}

	// determine starting neighbors
	startingPipes := getStaringNeighborPipes(startingPoint, matrix, pipes)

	maxFound := false
	currentPipeOne := startingPipes[0]
	onePrev := startingPoint
	currentPipeTwo := startingPipes[1]
	twoPrev := startingPoint

	// boundry indx tracker
	boundry := map[string]string{}

	// Helper to conv idx to string
	var idxStringConv = func(idx [2]int) string {
		idxStr := fmt.Sprint(idx[0]) + "," + fmt.Sprint(idx[1])
		return idxStr
	}

	// Setup starting points for matrix
	boundry[idxStringConv([2]int{startingPoint[0], startingPoint[1]})] = matrix[startingPoint[0]][startingPoint[1]]
	boundry[idxStringConv([2]int{currentPipeOne.coord[0], currentPipeOne.coord[1]})] = matrix[currentPipeOne.coord[0]][currentPipeOne.coord[1]]
	boundry[idxStringConv([2]int{currentPipeTwo.coord[0], currentPipeTwo.coord[1]})] = matrix[currentPipeTwo.coord[0]][currentPipeTwo.coord[1]]

	// Create the new circle & matrix
	for !maxFound {
		// Start one
		resultOne, _ := getNextPipe(currentPipeOne.coord, onePrev, matrix, pipes)
		onePrev = currentPipeOne.coord
		currentPipeOne = resultOne

		// Start two
		resultTwo, _ := getNextPipe(currentPipeTwo.coord, twoPrev, matrix, pipes)
		twoPrev = currentPipeTwo.coord
		currentPipeTwo = resultTwo

		boundry[idxStringConv([2]int{currentPipeOne.coord[0], currentPipeOne.coord[1]})] = matrix[currentPipeOne.coord[0]][currentPipeOne.coord[1]]
		boundry[idxStringConv([2]int{currentPipeTwo.coord[0], currentPipeTwo.coord[1]})] = matrix[currentPipeTwo.coord[0]][currentPipeTwo.coord[1]]

		if currentPipeTwo.coord == currentPipeOne.coord {
			maxFound = true
		}
	}

	// Create empty matrix, set anything that isn't the line to "." for future reference
	for i, row := range matrix {
		for j := range row {
			strIdx := idxStringConv([2]int{i, j})
			_, ok := boundry[strIdx]
			if !ok {
				matrix[i][j] = "."
			}
		}
	}

	// Expand the matrix
	expandedMatrix := expandMatrix(matrix, pipes)
	util.PrintMatrix(expandedMatrix)

	// Manually put first place to start, reset tracked
	stack := [][2]int{{130, 141}}
	tracked := map[string]bool{}
	var pipeTypes = map[string]bool{"J": true, "|": true, "-": true, "S": true, "F": true, "7": true, "L": true}

	// Track "." count
	insideCount := 0
	// Do a BFS
	for len(stack) > 0 {
		// Pop the first value
		curr := stack[0]
		stack = stack[1:]

		if expandedMatrix[curr[0]][curr[1]] == "." {
			insideCount++
		}

		// If we're on an edge, don't worry about it.
		if _, ok := pipeTypes[expandedMatrix[curr[0]][curr[1]]]; ok {
			continue
		}

		for _, n := range util.GetNeighborsPlus() {
			row := n[0] + curr[0]
			col := n[1] + curr[1]
			// Are we out of bounds?
			if row >= 0 && row <= len(expandedMatrix) && col <= len(expandedMatrix[0]) && col >= 0 {
				strIdx := idxStringConv([2]int{row, col})
				if _, ok := tracked[strIdx]; !ok {
					tracked[idxStringConv([2]int{row, col})] = true
					stack = append(stack, [2]int{row, col})
				}
			}
		}
	}

	fmt.Println(insideCount)
	return insideCount
}
