package main

import (
	"fmt"
	"os"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

func main() {
	Part01()
	// Part02()
}

func Part01() {
	dataInput, err := util.GetInput("16")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)
	funk.All(true)

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
