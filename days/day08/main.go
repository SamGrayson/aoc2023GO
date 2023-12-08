package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"util"
)

func main() {
	// Part01()
	Part02()
}

type Node struct {
	data  string
	left  *Node
	right *Node
}

func newNode(value string) *Node {
	return &Node{data: value, left: nil, right: nil}
}

type Direction struct {
	left  string
	right string
}

// var STEPS = "LR"

var STEPS = "LLRLRRRLLRRRLRRLRRLRLRRRLRRRLRLLRLRRLRRLRLLRRLRRRLRRLRLRLRLRRRLRRLRLLLRRLRRRLLLRLRRRLRRRLLRRLRRRLRLRRRLLLRRLLRRLRRLLLRRRLRRRLRRRLRRLLRLRLRLRRRLRLRLRRLRRLRLRRRLRRLRRRLRRRLLLRLRRLRRLRLLRRLLRRLRRLLRLRRLRRLRLRLLLRLLRRLRRLRRRLLRRLLRRRLRRLRRRLRRRLLRRRLRRRLLRRRLRLRLLRRLRLRLRRRR"

func Part01() int {
	dataInput, err := util.GetInput("08")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	var directionMap = make(map[string]Direction)

	lettersRe := regexp.MustCompile(`\w+`)
	for _, line := range inputArr {
		letters := lettersRe.FindAllString(line, -1)
		directionMap[letters[0]] = Direction{
			left:  letters[1],
			right: letters[2],
		}
	}

	var treeMap = make(map[string]*Node)
	var handleNodeInsert = func(value string, left string, right string) {
		if _, ok := treeMap[value]; !ok {
			treeMap[value] = newNode(value)
		}
		if _, ok := treeMap[left]; !ok {
			treeMap[left] = newNode(left)
		}
		if _, ok := treeMap[right]; !ok {
			treeMap[right] = newNode(right)
		}
		// Skip circular references
		if value == left && value == right {
			return
		}
		treeMap[value].left = treeMap[left]
		treeMap[value].right = treeMap[right]
	}

	// Create Tree Map
	for k, v := range directionMap {
		handleNodeInsert(k, v.left, v.right)
	}

	// Track how many steps it took
	var steps = 0
	// Storage for when ZZZ found
	var totalSteps = 0
	// Start at AAA
	stepTrack := treeMap["AAA"]
	for totalSteps == 0 {
		for _, v := range STEPS {
			steps++
			dir := string(v)
			if dir == "L" {
				stepTrack = stepTrack.left
			}
			if dir == "R" {
				stepTrack = stepTrack.right
			}
			if stepTrack.data == "ZZZ" {
				totalSteps = steps
			}
		}
	}

	fmt.Println(totalSteps)
	return totalSteps
}

func Part02() int {
	dataInput, err := util.GetInput("08")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	var directionMap = make(map[string]Direction)

	lettersRe := regexp.MustCompile(`\w+`)
	for _, line := range inputArr {
		letters := lettersRe.FindAllString(line, -1)
		directionMap[letters[0]] = Direction{
			left:  letters[1],
			right: letters[2],
		}
	}

	var treeMap = make(map[string]*Node)
	var handleNodeInsert = func(value string, left string, right string) {
		if _, ok := treeMap[value]; !ok {
			treeMap[value] = newNode(value)
		}
		if _, ok := treeMap[left]; !ok {
			treeMap[left] = newNode(left)
		}
		if _, ok := treeMap[right]; !ok {
			treeMap[right] = newNode(right)
		}
		// Skip circular references
		if value == left && value == right {
			return
		}
		treeMap[value].left = treeMap[left]
		treeMap[value].right = treeMap[right]
	}

	// Create Tree Map
	for k, v := range directionMap {
		handleNodeInsert(k, v.left, v.right)
	}

	type pointTrack struct {
		currentStart string
		currentStep  int
		steps        []int
	}

	// Setup
	startingPoints := []*pointTrack{}
	for k := range treeMap {
		if k[len(k)-1:] == "A" {
			startingPoints = append(startingPoints, &pointTrack{
				currentStart: k,
				currentStep:  0,
			})
		}
	}

	var allZExist = func(step int) bool {
		for _, v := range startingPoints {
			if len(v.steps) == 0 {
				return false
			}
		}
		return true
	}

	var allZs = false
	for !allZs {
		for _, point := range startingPoints {
			stepTrack := treeMap[point.currentStart]
			for _, v := range STEPS {
				point.currentStep++
				dir := string(v)
				if dir == "L" {
					stepTrack = stepTrack.left
				}
				if dir == "R" {
					stepTrack = stepTrack.right
				}
				point.currentStart = stepTrack.data
				if stepTrack.data[len(stepTrack.data)-1:] == "Z" {
					point.steps = append(point.steps, point.currentStep)
					if allZExist(point.currentStep) {
						allZs = true
						break
					}
				}
			}
			if allZs {
				break
			}
		}
	}

	lcmInput := []int{}
	for _, v := range startingPoints {
		lcmInput = append(lcmInput, v.steps[0])
	}
	final := util.LCMFromSlice(lcmInput)
	fmt.Println(final)
	return final
}
