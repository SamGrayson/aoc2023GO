package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"util"
)

type node struct {
	row          int
	col          int
	loss         int
	neighbors    *[]node
	shortest     int
	shortestPrev *node
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

func genKey(row, col int) string {
	return fmt.Sprint(row) + "," + fmt.Sprint(col)
}

func getLowestNext(unvisted map[string]*node) *node {
	shortest := &node{shortest: math.MaxInt}
	for _, v := range unvisted {
		if v.shortest <= shortest.shortest {
			shortest = v
		}
	}
	return shortest
}

func Part01() {
	dataInput, err := util.GetInput("17")
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
				row:          i,
				col:          j,
				loss:         util.ToInt(string(inputArr[i][j])),
				neighbors:    &[]node{},
				shortest:     math.MaxInt,
				shortestPrev: &node{},
			}
		}
	}

	visited := map[string]*node{}
	unvisitedNodes := map[string]*node{}

	// Create neighbors
	for i := 0; i < len(inputArr); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			generateNodeNeighbors(matrix[i][j], i, j, matrix)
			unvisitedNodes[genKey(i, j)] = matrix[i][j]
		}
	}

	currNode := matrix[0][0]
	currNode.loss = 0
	currNode.shortest = 0
	currNode.shortestPrev = nil
	delete(unvisitedNodes, genKey(currNode.row, currNode.col))

	for len(unvisitedNodes) > 0 {
		for i, neighbor := range *currNode.neighbors {
			if _, ok := visited[genKey(neighbor.row, neighbor.col)]; !ok {
				if currNode.shortest+neighbor.loss < neighbor.shortest {
					(*currNode.neighbors)[i].shortest = currNode.shortest + neighbor.loss
					(*currNode.neighbors)[i].shortestPrev = currNode
					(*unvisitedNodes[genKey(neighbor.row, neighbor.col)]).shortest = currNode.shortest + neighbor.loss
					(*unvisitedNodes[genKey(neighbor.row, neighbor.col)]).shortestPrev = currNode
				}
			}
		}
		visited[genKey(currNode.row, currNode.col)] = currNode
		delete(unvisitedNodes, genKey(currNode.row, currNode.col))

		currNode = getLowestNext(unvisitedNodes)
	}

	var gogo bool = true
	currNode = visited[genKey(len(matrix)-1, len(matrix[0])-1)]
	totalLoss := 0
	for gogo {
		if currNode.row == 0 && currNode.col == 0 {
			break
		}
		totalLoss += currNode.loss
		currNode = currNode.shortestPrev
	}

	fmt.Println("Part 1: ", totalLoss)
}

func Part02() {
	dataInput, err := util.GetInput("17")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)

	fmt.Println(inputArr)
}

func main() {
	Part01()
	// Part02()
}
