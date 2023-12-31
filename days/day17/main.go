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
	shortest     int
	shortestPrev *node
	path         []string
}

var dirMap map[string]string = map[string]string{
	"-1,0": "^",
	"1,0":  "v",
	"0,1":  ">",
	"0,-1": "<",
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
			unvisitedNodes[genKey(i, j)] = matrix[i][j]
		}
	}

	currNode := matrix[0][0]
	currNode.loss = 0
	currNode.shortest = 0
	currNode.shortestPrev = nil
	delete(unvisitedNodes, genKey(currNode.row, currNode.col))

	for len(unvisitedNodes) > 0 {
		neighborIdx := util.GetNeighborsPlus()
		for _, idx := range neighborIdx {
			nRow := idx[0] + currNode.row
			nCol := idx[1] + currNode.col
			direction := dirMap[genKey(idx[0], idx[1])]
			if nRow >= 0 && nRow <= len(matrix)-1 && nCol <= len(matrix[0])-1 && nCol >= 0 {
				next := matrix[nRow][nCol]
				if _, ok := visited[genKey(next.row, next.col)]; !ok {
					if currNode.shortest+next.loss < next.shortest {
						(*unvisitedNodes[genKey(next.row, next.col)]).shortest = currNode.shortest + next.loss
						(*unvisitedNodes[genKey(next.row, next.col)]).shortestPrev = currNode
						(*unvisitedNodes[genKey(next.row, next.col)]).path = append(currNode.path, direction)
					}
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
		fmt.Println(currNode.row, ", ", currNode.col)
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
