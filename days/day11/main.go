package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"util"

	"github.com/thoas/go-funk"
	"gonum.org/v1/gonum/mat"
)

func addEmptyColumn(matrix [][]float64, index int) [][]float64 {
	result := make([][]float64, len(matrix))

	for i, row := range matrix {
		// Create a new row with an additional empty element at the specified index
		newRow := make([]float64, len(row)+1)
		copy(newRow[:index], row[:index])
		copy(newRow[index+1:], row[index:])
		result[i] = newRow
	}

	return result
}

func addEmptyRow(matrix [][]float64, index int) [][]float64 {
	numRows := len(matrix)
	numCols := len(matrix[0]) // Assuming all rows have the same length

	// Create a new row with the same number of columns filled with zeros
	emptyRow := make([]float64, numCols)
	result := make([][]float64, numRows+1)

	for i := 0; i <= numRows; i++ {
		if i < index {
			result[i] = matrix[i]
		} else if i == index {
			result[i] = emptyRow
		} else {
			result[i] = matrix[i-1]
		}
	}

	return result
}

func Part01() {
	dataInput, err := util.GetInput("11")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	inputMatrix := make([][]float64, 0)
	galaxyI := 1
	for i := 0; i < len(inputArr); i++ {
		inputMatrix = append(inputMatrix, make([]float64, 0))
		split := strings.Split(inputArr[i], "")
		for _, v := range split {
			if v == "." {
				inputMatrix[i] = append(inputMatrix[i], float64(0))
			} else {
				inputMatrix[i] = append(inputMatrix[i], float64(galaxyI))
				galaxyI++
			}
		}
	}

	flatMatrix := funk.Flatten(inputMatrix).([]float64)
	inputMatrixMat := mat.NewDense(len(inputMatrix), len(inputMatrix[0]), flatMatrix)
	// Look for empty columns & append extra
	addedColumns := 0
	for i := range inputMatrix[0] {
		col := mat.Col(nil, i, inputMatrixMat)
		if funk.Sum(col) == 0 {
			inputMatrix = addEmptyColumn(inputMatrix, i+addedColumns)
			addedColumns++
		}
	}

	// Look for empty rows & append extra
	addedRows := 0
	for i := range inputMatrix {
		row := mat.Row(nil, i, inputMatrixMat)
		if funk.Sum(row) == 0 {
			inputMatrix = addEmptyRow(inputMatrix, i+addedRows)
			addedRows++
		}
	}

	// Loop through expanded matrix and get the mappings for each galaxy & their location.
	var galaxyMappings = make(map[int][2]float64, 0)
	for i, row := range inputMatrix {
		for j, col := range row {
			if col != 0 {
				galaxyMappings[int(col)] = [2]float64{float64(i), float64(j)}
			}
		}
	}

	// util.PrintMatrixFloat(inputMatrix)

	// Find all the shortest distances
	var distances = []float64{}
	for i := 1; i <= len(galaxyMappings); i++ {
		start := galaxyMappings[i]
		for j := len(galaxyMappings); j > i; j-- {
			distance := util.GetManhattanDistance(start, galaxyMappings[j])
			distances = append(distances, distance)
		}
	}

	distanceStr := strconv.FormatFloat(funk.Sum(distances), 'f', -1, 64)
	fmt.Println(distanceStr)
}

func Part02(multiplier float64) {
	dataInput, err := util.GetInput("11")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	inputMatrix := make([][]float64, 0)
	galaxyI := 1
	for i := 0; i < len(inputArr); i++ {
		inputMatrix = append(inputMatrix, make([]float64, 0))
		split := strings.Split(inputArr[i], "")
		for _, v := range split {
			if v == "." {
				inputMatrix[i] = append(inputMatrix[i], float64(0))
			} else {
				inputMatrix[i] = append(inputMatrix[i], float64(galaxyI))
				galaxyI++
			}
		}
	}

	// Loop through expanded matrix and get the mappings for each galaxy & their location.
	var galaxyMappings = make(map[int][2]float64, 0)
	for i, row := range inputMatrix {
		for j, col := range row {
			if col != 0 {
				I := float64(i)
				J := float64(j)
				galaxyMappings[int(col)] = [2]float64{I, J}
			}
		}
	}

	flatMatrix := funk.Flatten(inputMatrix).([]float64)
	inputMatrixMat := mat.NewDense(len(inputMatrix), len(inputMatrix[0]), flatMatrix)

	// Detect empty cols, add the multiplier to every [_, y] after the point.
	rowMultiplier := 0
	for i := range inputMatrix {
		row := mat.Row(nil, i, inputMatrixMat)
		if funk.Sum(row) == 0 {
			rowMultiplier++
		}
		columnMultiplier := 0
		for j, v := range row {
			col := mat.Col(nil, j, inputMatrixMat)
			if funk.Sum(col) == 0 {
				columnMultiplier++
			}
			if v != 0 {
				mapVal := galaxyMappings[int(v)]
				galaxyMappings[int(v)] = [2]float64{
					mapVal[0] + (float64(rowMultiplier) * multiplier),
					mapVal[1] + (float64(columnMultiplier) * multiplier),
				}
			}
		}
	}

	// Find all the shortest distances
	var distances = []float64{}
	for i := 1; i <= len(galaxyMappings); i++ {
		start := galaxyMappings[i]
		for j := len(galaxyMappings); j > i; j-- {
			distance := util.GetManhattanDistance(start, galaxyMappings[j])
			distances = append(distances, distance)
		}
	}

	distanceStr := strconv.FormatFloat(funk.Sum(distances), 'f', -1, 64)
	fmt.Println(distanceStr)
}

func main() {
	// Part01()
	Part02(999999)
}
