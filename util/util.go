package util

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strconv"

	"github.com/thoas/go-funk"
)

func ToInt(s string) int {
	res, _ := strconv.Atoi(s)

	return res
}

func ToFloat(s string) float64 {
	res, _ := strconv.ParseFloat(s, 64)

	return res
}

func GetRelativePath(day string, fileName string) (string, error) {
	wd, _ := os.Getwd()

	// Are we debuggin?
	debug := os.Getenv("DEBUG")

	// Set debug path
	var path string
	if !(debug == "true") {
		path = fmt.Sprintf("./days/day%s/"+fileName, day)
	} else {
		path = fileName
	}

	// Read the text file
	pathToInput := filepath.Join(wd, path)
	return pathToInput, nil
}

func GetInput(day string) (string, error) {
	// Get current working directory
	pathToInput, _ := GetRelativePath(day, "input.txt")
	fileContent, err := os.ReadFile(pathToInput)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	return string(fileContent), nil
}

func IsNum(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}
	return false
}

func RemoveSpaceChar(arr []string) []string {
	newArr := []string{}
	for _, v := range arr {
		if v != " " {
			newArr = append(newArr, (v))
		}
	}
	return newArr
}

func RemoveEmptyChar(arr []string) []string {
	newArr := []string{}
	for _, v := range arr {
		if v != "" {
			newArr = append(newArr, (v))
		}
	}
	return newArr
}

func SliceToMap(arr []string) map[string]bool {
	ret := make(map[string]bool)
	for i := 0; i < len(arr); i += 1 {
		ret[arr[i]] = true
	}
	return ret
}

func ArrStringToInt(arr []string) []int {
	ret := funk.Map(
		arr,
		func(n string) int {
			new, _ := strconv.Atoi(n)
			return new
		})
	return ret.([]int)
}

func NumMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NumMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BigMul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}

// Function to calculate the greatest common divisor (GCD) of two numbers using Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate the least common multiple (LCM) of two numbers
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// Function to find the least common denominator (LCD) of multiple numbers
func LCMFromSlice(numbers []int) int {
	if len(numbers) < 2 {
		panic("At least two numbers required")
	}

	lcmValue := numbers[0]
	for i := 1; i < len(numbers); i++ {
		lcmValue = LCM(lcmValue, numbers[i])
	}
	return lcmValue
}

func GetNeighborsSquare() [][]int {
	return [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /*{0, 0}, */, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
}

func GetNeighborsPlus() [][]int {
	// up down left right
	return [][]int{
		{-1, 0}, {1, 0} /*{0,0}, */, {0, -1}, {0, 1},
	}
}

func PrintMatrix(matrix [][]string) {
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%s", cell)
		}
		fmt.Println()
	}
}

func PrintIntMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%d", cell)
		}
		fmt.Println()
	}
}

func PrintMatrixFloat(matrix [][]float64) {
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%d", int(cell))
		}
		fmt.Println()
	}
}

func IdxStringConv(idx [2]int) string {
	idxStr := fmt.Sprint(idx[0]) + "," + fmt.Sprint(idx[1])
	return idxStr
}

func GetManhattanDistance(p1 [2]float64, p2 [2]float64) float64 {
	return math.Abs(p2[0]-p1[0]) + math.Abs(p2[1]-p1[1])
}

// Returns the string difference and the index its located at
func Difference(a, b []string) ([]string, []int) {
	diffStr := []string{}
	diffInt := []int{}
	for i, va := range a {
		if va != b[i] {
			diffStr = append(diffStr, va)
			diffInt = append(diffInt, i)
		}
	}
	return diffStr, diffInt
}

func ReplaceAtIdx(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func ArrWithDefaultStr(len int, def string) []string {
	arr := make([]string, len)
	for i := range arr {
		arr[i] = def
	}
	return arr
}
