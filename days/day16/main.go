package main

import (
	"fmt"
	"os"
	"strings"
	"util"
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
	"^": {-0, 0},
}

type lazer struct {
	coord     [2]int
	direction string
}

func motion(laz *lazer, direction string, matrix [][]*node, lIM map[int]*lazer) bool {
	currentCoord := laz.coord
	currNode := matrix[currentCoord[0]][currentCoord[1]]
	row := currNode.row + directions[direction][0]
	col := currNode.col + directions[direction][1]
	// Are we in bounds?
	if row >= 0 && row <= len(matrix) && col <= len(matrix[0]) && col >= 0 {
		next := matrix[row][col]

		// If blank, just go there
		if next.str == "." {
			laz.coord = [2]int{next.row, next.col}
			return false
		}

		if next.str == "\\" {
			if direction == ">"
			lazer.coord = 2[int]{next.row, next.col}
			return false
		}
		// If blank, just go there
		if next.str == "/" {
			lazer.coord = 2[int]{next.row, next.col}
			return false
		}
		// If blank, just go there
		if next.str == "|" {
			lazer.coord = 2[int]{next.row, next.col}
			return false
		}

		// If blank, just go there
		if next.str == "-" {
			lazer.coord = 2[int]{next.row, next.col}
			return false
		}
	} else {
		return true
	}
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
	lazersInMotion := map[int]*lazer{
		0: {
			direction: ">",
			coord:     lazerStart,
		},
	}

	for len(lazersInMotion) > 0 {
		// Track the lazer through all the neighbors
		for lazerKey, currLazer := range lazersInMotion {
			complete := motion(*currLazer, ">", matrix, lazersInMotion)
			if complete {
				delete(lazersInMotion, lazerKey)
			}
		}

		// If the lazer gets to the end of the matrix, its done. (or maybe makes a loop??)
	}

	fmt.Println(inputArr)
}

func Part02() {
	dataInput, err := util.GetInput("16")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)

	fmt.Println(inputArr)
}
