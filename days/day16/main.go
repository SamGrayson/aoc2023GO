package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"util"

	"github.com/google/uuid"
)

func main() {
	Part01()
	// Part02()
}

type node struct {
	row       int
	col       int
	str       string
	neighbors *[]node
	visited   bool
}

func generateNodeNeighbors(n *node, row, col int, matrix [][]*node) {
	neighborIdx := util.GetNeighborsPlus()
	for _, idx := range neighborIdx {
		nRow := idx[0] + n.row
		nCol := idx[1] + n.col
		if nRow >= 0 && nRow <= len(matrix)-1 && nCol <= len(matrix[0])-1 && nCol >= 0 {
			*n.neighbors = append(*n.neighbors, *matrix[nRow][nCol])
		}
	}
}

var directions = map[string][2]int{
	">": {0, 1},
	"<": {0, -1},
	"v": {1, 0},
	"^": {-1, 0},
}

type lazer struct {
	coord     [2]int
	direction string
	visited   *[]string
}

func inBounds(row, col, rowLen, colLen int) bool {
	return row >= 0 && row <= rowLen && col <= colLen && col >= 0
}

func motion(laz *lazer, direction string, matrix [][]*node, lIM map[string]*lazer) bool {
	currentCoord := laz.coord
	currNode := matrix[currentCoord[0]][currentCoord[1]]
	row := currNode.row + directions[direction][0]
	col := currNode.col + directions[direction][1]
	currNode.visited = true

	// if we've already been there with this lazer, we found a loop
	if slices.Contains(*laz.visited, fmt.Sprint(row)+","+fmt.Sprint(col)) {
		return true
	}

	// Are we in bounds?
	if inBounds(row, col, len(matrix)-1, len(matrix[0])-1) {
		(*laz.visited) = append(*laz.visited, fmt.Sprint(row)+","+fmt.Sprint(col))
		next := matrix[row][col]
		laz.coord = [2]int{next.row, next.col}
		next.visited = true

		if next.str == "." {
			return false
		}

		if next.str == "\\" {
			if direction == ">" {
				laz.direction = "v"
			}
			if direction == "<" {
				laz.direction = "^"
			}
			if direction == "^" {
				laz.direction = "<"
			}
			if direction == "v" {
				laz.direction = ">"
			}
			return false
		}
		if next.str == "/" {
			if direction == ">" {
				laz.direction = "^"
			}
			if direction == "<" {
				laz.direction = "v"
			}
			if direction == "^" {
				laz.direction = ">"
			}
			if direction == "v" {
				laz.direction = "<"
			}
			return false
		}
		if next.str == "|" {
			if direction == ">" || direction == "<" {
				id := uuid.New()
				lIM[id.String()] = &lazer{
					direction: "^",
					coord:     laz.coord,
					visited:   &[]string{},
				}
				id = uuid.New()
				lIM[id.String()] = &lazer{
					direction: "v",
					coord:     laz.coord,
					visited:   &[]string{},
				}
				return true
			}
			return false
		}

		// If blank, just go there
		if next.str == "-" {
			if direction == "^" || direction == "v" {
				id := uuid.New()
				lIM[id.String()] = &lazer{
					direction: "<",
					coord:     laz.coord,
					visited:   &[]string{},
				}
				id = uuid.New()
				lIM[id.String()] = &lazer{
					direction: ">",
					coord:     laz.coord,
					visited:   &[]string{},
				}
				return true
			}
			return false
		}
	}
	// If we're out of bounds, this route is done.
	return true
}

func Part01() {
	dataInput, err := util.GetInput("16")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	matrix := make([][]*node, len(inputArr))

	// Create node matrix
	for i := 0; i < len(inputArr); i++ {
		matrix[i] = make([]*node, len(inputArr[i]))
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = &node{
				row:       i,
				col:       j,
				str:       string(inputArr[i][j]),
				neighbors: &[]node{},
				visited:   false,
			}
		}
	}

	// Create neighbors
	for i := 0; i < len(inputArr); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			generateNodeNeighbors(matrix[i][j], i, j, matrix)
		}
	}

	lazerStart := [2]int{0, 0}
	start_id := uuid.New()
	lazersInMotion := map[string]*lazer{
		start_id.String(): {
			direction: "v",
			coord:     lazerStart,
			visited:   &[]string{},
		},
	}

	// For debugging
	debugVisitedNodes := []*node{}
	resultMatrix := [][]string{}
	//

	prevValueStop := 0

	for len(lazersInMotion) > 0 {
		// Track the lazer through all the neighbors
		for lazerKey, currLazer := range lazersInMotion {
			complete := motion(currLazer, currLazer.direction, matrix, lazersInMotion)
			if complete {
				delete(lazersInMotion, lazerKey)
			}
		}
		//
		visitedSpots := 0
		for i, row := range matrix {
			resultMatrix = append(resultMatrix, make([]string, len(row)))
			for j, col := range row {
				if col.visited {
					visitedSpots += 1
					debugVisitedNodes = append(debugVisitedNodes, col)
					resultMatrix[i][j] = "#"
				} else {
					resultMatrix[i][j] = col.str
				}
			}
		}
		if prevValueStop == visitedSpots {
			fmt.Println("Part 1: ", visitedSpots)
			util.PrintMatrix(resultMatrix)
		}
		prevValueStop = visitedSpots
	}
}

func Part02() {
	dataInput, err := util.GetInput("16")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)

	fmt.Println(inputArr)
}
