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

func inBounds(row, col, rowLen, colLen int) bool {
	return row >= 0 && row <= rowLen && col <= colLen && col >= 0
}

func Part01() {
	dataInput, err := util.GetInput("16")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	matrix := make([][]*node, len(inputArr))

	visited := map[string]bool{}

	var recurLazer func(node *node, direction string) int
	recurLazer = func(node *node, direction string) int {
		// we went 1 too far, subtract 1
		if visited[fmt.Sprint(node.row)+","+fmt.Sprint(node.col)+direction] {
			return -1
		}
		visited[fmt.Sprint(node.row)+","+fmt.Sprint(node.col)+direction] = true

		node.visited = true

		row := node.row + directions[direction][0]
		col := node.col + directions[direction][1]

		allVisit := 0
		// Are we in bounds?
		if inBounds(row, col, len(matrix)-1, len(matrix[0])-1) {
			allVisit = 1
			next := matrix[row][col]

			if next.str == "." {
				allVisit += recurLazer(next, direction)
			}

			if next.str == "\\" {
				if direction == ">" {
					allVisit += recurLazer(next, "v")
				}
				if direction == "<" {
					allVisit += recurLazer(next, "^")
				}
				if direction == "^" {
					allVisit += recurLazer(next, "<")
				}
				if direction == "v" {
					allVisit += recurLazer(next, ">")
				}
			}
			if next.str == "/" {
				if direction == ">" {
					allVisit += recurLazer(next, "^")
				}
				if direction == "<" {
					allVisit += recurLazer(next, "v")
				}
				if direction == "^" {
					allVisit += recurLazer(next, ">")
				}
				if direction == "v" {
					allVisit += recurLazer(next, "<")
				}
			}
			if next.str == "|" {
				if direction == ">" || direction == "<" {
					allVisit += recurLazer(next, "^")
					allVisit += recurLazer(next, "v")
				} else {
					allVisit += recurLazer(next, direction)
				}
			}

			// If blank, just go there
			if next.str == "-" {
				if direction == "^" || direction == "v" {
					allVisit += recurLazer(next, "<")
					allVisit += recurLazer(next, ">")
				} else {
					allVisit += recurLazer(next, direction)
				}
			}
		}
		// If we're out of bounds, this route is done.
		return allVisit
	}

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

	startNode := matrix[0][0]
	startDirection := "v"

	totalVisited := recurLazer(startNode, startDirection)
	fmt.Println("Part 1: ", totalVisited)
}

func Part02() {
	dataInput, err := util.GetInput("16")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)

	fmt.Println(inputArr)
}
