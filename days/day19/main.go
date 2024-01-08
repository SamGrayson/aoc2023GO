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

type rule struct {
	start       string
	comparitive string
	value       int
	route       string
}

type ruleSet struct {
	rules []rule
}

func Part01() {
	dataInput, err := util.GetInput("19")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	directions := map[string]ruleSet{}
	testSets := []map[string]int{}

	var solve func(vals map[string]int, curr string) string
	solve = func(vals map[string]int, curr string) string {
		dir := directions[curr]
		var nextRoute string
		for _, rule := range dir.rules {
			if rule.comparitive != "" && rule.comparitive == ">" {
				if vals[rule.start] > rule.value {
					nextRoute = rule.route
					break
				}
			}

			if rule.comparitive != "" && rule.comparitive == "<" {
				if vals[rule.start] < rule.value {
					nextRoute = rule.route
					break
				}
			}

			nextRoute = rule.route
		}
		if nextRoute == "A" || nextRoute == "R" {
			return nextRoute
		} else {
			return solve(vals, nextRoute)
		}
	}

	spaceFound := false
	for _, input := range inputArr {
		if input == "" {
			spaceFound = true
			continue
		}

		// if the space is found, do the input parse instead
		if spaceFound {
			line := input[1 : len(input)-1]
			lineVals := strings.Split(line, ",")
			sets := map[string]int{}
			for _, v := range lineVals {
				split := strings.Split(v, "=")
				sets[split[0]] = util.ToInt(split[1])
			}
			testSets = append(testSets, sets)
			continue
		}

		nameSplut := strings.Split(input, "{")
		ruleName := nameSplut[0]
		// Take off the }
		objSplit := nameSplut[1][:len(nameSplut[1])-1]
		rules := strings.Split(objSplit, ",")
		ruleSetBld := ruleSet{rules: []rule{}}
		// Create the rules
		for _, r := range rules {
			ruleSplit := strings.Split(r, "")
			var start string
			var comparitive string = ""
			var value string
			var route string
			var colonFound bool = false
			for _, c := range ruleSplit {
				if c == ">" || c == "<" {
					comparitive = c
					continue
				}
				if colonFound {
					route += c
					continue
				}
				if c == ":" {
					colonFound = true
					continue
				}
				if comparitive != "" {
					value += c
				} else {
					start += c
				}
			}
			// If no comparitive, its actually a route.
			if comparitive == "" {
				ruleSetBld.rules = append(ruleSetBld.rules, rule{
					route: start,
				})
				continue
			}
			ruleSetBld.rules = append(ruleSetBld.rules, rule{
				start:       start,
				comparitive: comparitive,
				value:       util.ToInt(value),
				route:       route,
			})
		}
		directions[ruleName] = ruleSetBld
	}

	acceptedPatterns := []map[string]int{}
	for _, test := range testSets {
		if solve(test, "in") == "A" {
			acceptedPatterns = append(acceptedPatterns, test)
		}

	}

	total := 0
	for _, pat := range acceptedPatterns {
		for _, v := range pat {
			total += v
		}
	}

	fmt.Println("Part 1 total: ", total)
}

func Part02() {
	dataInput, err := util.GetInput("02")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Fields(dataInput)

	fmt.Println(inputArr)
}
