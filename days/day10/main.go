package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"util"
)

func main() {
	Part01()
	// Part02()
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
	// Util returns in up down left right, so we need to check in the opposite direction..
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
	// Util returns in up down left right, so we need to check in the opposite direction..
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

func Part02() {
	dataInput, err := util.GetInput("02")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)

	fmt.Println(inputArr)
}
