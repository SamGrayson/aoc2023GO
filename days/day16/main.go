package main

import (
	"fmt"
	"os"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

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

	var recurLazer func(node *node, direction string)
	recurLazer = func(node *node, direction string) {
		// we went 1 too far, subtract 1
		if visited[fmt.Sprint(node.row)+","+fmt.Sprint(node.col)+direction] {
			return
		}
		visited[fmt.Sprint(node.row)+","+fmt.Sprint(node.col)+direction] = true
		node.visited = true

		row := node.row + directions[direction][0]
		col := node.col + directions[direction][1]

		// Are we in bounds?
		if inBounds(row, col, len(matrix)-1, len(matrix[0])-1) {
			next := matrix[row][col]

			if next.str == "." {
				recurLazer(next, direction)
			}

			if next.str == "\\" {
				if direction == ">" {
					recurLazer(next, "v")
				}
				if direction == "<" {
					recurLazer(next, "^")
				}
				if direction == "^" {
					recurLazer(next, "<")
				}
				if direction == "v" {
					recurLazer(next, ">")
				}
			}
			if next.str == "/" {
				if direction == ">" {
					recurLazer(next, "^")
				}
				if direction == "<" {
					recurLazer(next, "v")
				}
				if direction == "^" {
					recurLazer(next, ">")
				}
				if direction == "v" {
					recurLazer(next, "<")
				}
			}
			if next.str == "|" {
				if direction == ">" || direction == "<" {
					recurLazer(next, "^")
					recurLazer(next, "v")
				} else {
					recurLazer(next, direction)
				}
			}

			// If blank, just go there
			if next.str == "-" {
				if direction == "^" || direction == "v" {
					recurLazer(next, "<")
					recurLazer(next, ">")
				} else {
					recurLazer(next, direction)
				}
			}
		}
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

	recurLazer(startNode, startDirection)

	// For testing
	finalMatrix := [][]string{}
	totalVisited := 0
	for i, row := range matrix {
		finalMatrix = append(finalMatrix, make([]string, len(row)))
		for j, col := range row {
			if col.visited {
				finalMatrix[i][j] = "#"
				totalVisited++
			} else {
				finalMatrix[i][j] = col.str
			}
		}
	}

	// For testing
	// util.PrintMatrix(finalMatrix)

	fmt.Println("Part 1: ", totalVisited)
}

func resetVisitMatrix(matrix [][]*node) {
	for _, row := range matrix {
		for _, col := range row {
			col.visited = false
		}
	}
}

func getVisitedTotal(matrix [][]*node) int {
	totalVisited := 0
	for _, row := range matrix {
		for _, col := range row {
			if col.visited {
				totalVisited++
			}
		}
	}
	return totalVisited
}

func Part02() {
	dataInput, err := util.GetInput("16")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	matrix := make([][]*node, len(inputArr))

	visited := map[string]bool{}

	var recurLazer func(node *node, direction string)
	recurLazer = func(node *node, direction string) {
		// we went 1 too far, subtract 1
		if visited[fmt.Sprint(node.row)+","+fmt.Sprint(node.col)+direction] {
			return
		}
		visited[fmt.Sprint(node.row)+","+fmt.Sprint(node.col)+direction] = true
		node.visited = true

		row := node.row + directions[direction][0]
		col := node.col + directions[direction][1]

		// Are we in bounds?
		if inBounds(row, col, len(matrix)-1, len(matrix[0])-1) {
			next := matrix[row][col]

			if next.str == "." {
				recurLazer(next, direction)
			}

			if next.str == "\\" {
				if direction == ">" {
					recurLazer(next, "v")
				}
				if direction == "<" {
					recurLazer(next, "^")
				}
				if direction == "^" {
					recurLazer(next, "<")
				}
				if direction == "v" {
					recurLazer(next, ">")
				}
			}
			if next.str == "/" {
				if direction == ">" {
					recurLazer(next, "^")
				}
				if direction == "<" {
					recurLazer(next, "v")
				}
				if direction == "^" {
					recurLazer(next, ">")
				}
				if direction == "v" {
					recurLazer(next, "<")
				}
			}
			if next.str == "|" {
				if direction == ">" || direction == "<" {
					recurLazer(next, "^")
					recurLazer(next, "v")
				} else {
					recurLazer(next, direction)
				}
			}

			// If blank, just go there
			if next.str == "-" {
				if direction == "^" || direction == "v" {
					recurLazer(next, "<")
					recurLazer(next, ">")
				} else {
					recurLazer(next, direction)
				}
			}
		}
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

	rows := len(matrix)
	cols := len(matrix[0])

	totals := []int{}

	// Loop through the top edge
	for i := 0; i < cols; i++ {
		currNode := matrix[0][i]
		if currNode.str == "\\" {
			recurLazer(currNode, ">")
		}
		if currNode.str == "-" {
			recurLazer(currNode, ">")
			visited = map[string]bool{}
			totals = append(totals, getVisitedTotal(matrix))
			resetVisitMatrix(matrix)
			recurLazer(currNode, "<")
		}
		if currNode.str == "/" {
			recurLazer(currNode, "<")
		}
		if currNode.str == "." || currNode.str == "|" {
			recurLazer(currNode, "v")
		}
		totals = append(totals, getVisitedTotal(matrix))
		visited = map[string]bool{}
		resetVisitMatrix(matrix)
	}

	// Loop through the right edge
	for i := 1; i < rows; i++ {
		currNode := matrix[i][cols-1]
		if currNode.str == "\\" {
			recurLazer(currNode, "^")
		}
		if currNode.str == "|" {
			recurLazer(currNode, "^")
			visited = map[string]bool{}
			totals = append(totals, getVisitedTotal(matrix))
			resetVisitMatrix(matrix)
			recurLazer(currNode, "v")
		}
		if currNode.str == "/" {
			recurLazer(currNode, "v")
		}
		if currNode.str == "." || currNode.str == "-" {
			recurLazer(currNode, "<")
		}
		totals = append(totals, getVisitedTotal(matrix))
		visited = map[string]bool{}
		resetVisitMatrix(matrix)
	}

	// Loop through the bottom edge
	for i := cols - 1; i >= 0; i-- {
		currNode := matrix[rows-1][i]
		if currNode.str == "\\" {
			recurLazer(currNode, "<")
		}
		if currNode.str == "-" {
			recurLazer(currNode, ">")
			visited = map[string]bool{}
			totals = append(totals, getVisitedTotal(matrix))
			resetVisitMatrix(matrix)
			recurLazer(currNode, "<")
		}
		if currNode.str == "/" {
			recurLazer(currNode, ">")
		}
		if currNode.str == "." || currNode.str == "|" {
			recurLazer(currNode, "^")
		}
		totals = append(totals, getVisitedTotal(matrix))
		visited = map[string]bool{}
		resetVisitMatrix(matrix)
	}

	// Loop through the left edge
	for i := rows - 1; i >= 0; i-- {
		currNode := matrix[i][0]
		if currNode.str == "\\" {
			recurLazer(currNode, "v")
		}
		if currNode.str == "|" {
			recurLazer(currNode, "v")
			visited = map[string]bool{}
			totals = append(totals, getVisitedTotal(matrix))
			resetVisitMatrix(matrix)
			recurLazer(currNode, "^")
		}
		if currNode.str == "/" {
			recurLazer(currNode, "^")
		}
		if currNode.str == "." || currNode.str == "-" {
			recurLazer(currNode, ">")
		}
		totals = append(totals, getVisitedTotal(matrix))
		visited = map[string]bool{}
		resetVisitMatrix(matrix)
	}

	fmt.Println("Part 2: ", funk.MaxInt(totals))
}

func main() {
	// Part01()
	Part02()
}
