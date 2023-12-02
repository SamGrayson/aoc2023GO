package day2

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

func Part01() {
	dataInput, err := util.GetInput("02")
	if err != nil { os.Exit(1) }
	inputArr := strings.Fields(dataInput)

	fmt.Println(inputArr)
}

func Part02() {
	dataInput, err := util.GetInput("02")
	if err != nil { os.Exit(1) }
	inputArr := strings.Fields(dataInput)

	fmt.Println(inputArr)
}
