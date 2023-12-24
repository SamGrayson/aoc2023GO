package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"util"

	"github.com/thoas/go-funk"
)

func hashCalculation(start int, currentValue int) int {
	return ((currentValue + start) * 17) % 256
}

func Part01() {
	dataInput, err := util.GetInput("15")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, ",")

	values := []int{}
	currentValue := 0
	for _, v := range inputArr {
		for _, s := range v {
			currentValue = hashCalculation(int(s), currentValue)
		}
		values = append(values, currentValue)
		currentValue = 0
	}

	fmt.Println("part 1: ", funk.Sum(values))
}

func getBoxInput(str string) int {
	currentValue := 0
	for _, s := range str {
		currentValue = hashCalculation(int(s), currentValue)
	}
	return currentValue
}

type box struct {
	order  *[]string
	lenses map[string]bool
}

func removeLabel(labels []string, labelToRemove string) []string {
	return funk.Filter(labels, func(l string) bool {
		return !strings.Contains(l, labelToRemove)
	}).([]string)
}

func replaceLabel(labels []string, labelToReplace, number string) []string {
	return funk.Map(labels, func(l string) string {
		if strings.Contains(l, labelToReplace) {
			return labelToReplace + " " + number
		}
		return l
	}).([]string)
}

func Part02() {
	dataInput, err := util.GetInput("15")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, ",")

	boxes := map[int]box{}
	for _, v := range inputArr {
		// Split string into directions
		if strings.Contains(v, "-") {
			split := strings.Split(v, "-")
			label := split[0]
			currBox := getBoxInput(label)
			// If the box already exists, check the lense next
			if bx, ok := boxes[currBox]; ok {
				bx.lenses[label] = false
				*bx.order = removeLabel(*bx.order, label)
			}
		}
		if strings.Contains(v, "=") {
			split := strings.Split(v, "=")
			label := split[0]
			number := split[1]
			currBox := getBoxInput(label)
			// If the box already exists, check the lense next
			if bx, ok := boxes[currBox]; ok {
				bx.lenses[label] = true
				*bx.order = replaceLabel(*bx.order, label, number)
				idx := slices.Index(*bx.order, label+" "+number)
				if idx == -1 {
					*bx.order = append(*bx.order, label+" "+number)
				}
			} else {
				boxes[currBox] = box{
					order:  &[]string{label + " " + number},
					lenses: map[string]bool{label: true},
				}
			}
		}
	}

	boxValues := 0

	for k, bx := range boxes {
		bValue := k + 1
		for i, lense := range *bx.order {
			split := strings.Split(lense, " ")
			if len(split) == 2 {
				focusValue := util.ToInt(split[1])
				boxValues += (bValue * (i + 1) * focusValue)
			}
		}
	}

	fmt.Println("part 2: ", boxValues)
}

func main() {
	// Part01()
	Part02()
}
