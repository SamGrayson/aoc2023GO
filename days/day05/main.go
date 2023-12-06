package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"util"

	"github.com/thoas/go-funk"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func main() {
	// fmt.Println("Min location p1: ", Part01())
	fmt.Println("Min location p2: ", Part02())
}

func getRangeValue(conv int, ranges []farmMapping) int {
	for _, r := range ranges {
		if r.srcStart <= conv && conv <= (r.srcStart+r.Range) {
			diff := conv - r.srcStart
			return r.destStart + diff
		}
	}
	return conv
}

// Source/Destination just for debugging
type soilMap struct {
	uniqueRanges    []farmMapping
	undesireableSrc map[int]bool
}

func Part01() int {
	dataInput, err := util.GetInput("05")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	digitRe := regexp.MustCompile(`\d+`)

	// Set seed values
	seeds := util.ArrStringToInt(digitRe.FindAllString(inputArr[0], -1))

	// Generate maps
	mapTracker := orderedmap.New[string, *soilMap]()
	currentMap := ""
	for _, line := range inputArr[2:] {
		if len(line) == 0 {
			continue
		}
		if digitRe.FindStringIndex(line) == nil {
			currentMap = line
			mapTracker.Set(currentMap, &soilMap{
				uniqueRanges: []farmMapping{},
			})
		} else {
			inputLine := util.ArrStringToInt(digitRe.FindAllString(line, -1))
			conversion, _ := mapTracker.Get(currentMap)
			conversion.uniqueRanges = append(conversion.uniqueRanges, farmMapping{
				destStart: inputLine[0],
				srcStart:  inputLine[1],
			})
			mapTracker.Set(currentMap, conversion)
		}
	}

	var locations []int
	// Seed Lookup
	for i := 0; i < len(seeds); i++ {
		start := seeds[i]

		for pair := mapTracker.Oldest(); pair != nil; pair = pair.Next() {
			start = getRangeValue(start, pair.Value.uniqueRanges)
		}
		locations = append(locations, start)
	}

	return funk.MinInt(locations)
}

type farmMapping struct {
	destStart int
	destEnd   int
	srcStart  int
	srcEnd    int
	Range     int
}

func Part02() int {
	dataInput, err := util.GetInput("05")
	if err != nil {
		os.Exit(1)
	}
	inputArr := strings.Split(dataInput, "\n")

	digitRe := regexp.MustCompile(`\d+`)
	seedGroupRe := regexp.MustCompile(`(\d+) (\d+)`)

	// Set seed values
	seedStr := seedGroupRe.FindAllString(inputArr[0], -1)
	var seedRanges [][]int
	for _, s := range seedStr {
		seedStartStr := strings.Split(s, " ")[0]
		seedRangeStr := strings.Split(s, " ")[1]
		sStart, _ := strconv.Atoi(seedStartStr)
		sRange, _ := strconv.Atoi(seedRangeStr)
		seedRange := []int{sStart, sRange}
		seedRanges = append(seedRanges, seedRange)
	}

	// Generate maps
	mapTracker := orderedmap.New[string, soilMap]()
	currentMap := ""
	for _, line := range inputArr[2:] {
		if len(line) == 0 {
			continue
		}
		if digitRe.FindStringIndex(line) == nil {
			currentMap = line
			mapTracker.Set(currentMap, soilMap{
				uniqueRanges:    []farmMapping{},
				undesireableSrc: make(map[int]bool),
			})
		} else {
			inputLine := util.ArrStringToInt(digitRe.FindAllString(line, -1))
			conversion, _ := mapTracker.Get(currentMap)
			conversion.uniqueRanges = append(conversion.uniqueRanges, farmMapping{
				destStart: inputLine[0],
				destEnd:   inputLine[0] + inputLine[2],
				srcStart:  inputLine[1],
				srcEnd:    inputLine[1] + inputLine[2],
				Range:     inputLine[2],
			})
			sort.Slice(conversion.uniqueRanges, func(i int, j int) bool {
				return conversion.uniqueRanges[i].srcStart < conversion.uniqueRanges[j].srcStart
			})
			mapTracker.Set(currentMap, conversion)
		}
	}

	var doSeedLookup = func(results *[]int, incSeed []int) {
		minLocation := math.MaxInt
		for seed := incSeed[0]; seed < incSeed[0]+incSeed[1]; seed += 1 {
			// Look through the pairs and track the lowest next val
			lowestInc := seed
			for pair := mapTracker.Oldest(); pair != nil; pair = pair.Next() {
				for _, mRange := range pair.Value.uniqueRanges {
					if mRange.srcStart <= lowestInc && lowestInc < mRange.srcEnd {
						distance := lowestInc - mRange.srcStart
						lowestInc = mRange.destStart + distance
					}
					if pair.Key == "humidity-to-location" {
						if lowestInc < minLocation {
							minLocation = lowestInc
						}
					}
				}
			}
		}
		*results = append(*results, minLocation)
	}

	// Try and use the ranges
	results := []int{}
	for _, v := range seedRanges {
		// Do all the seeds - figure out how to do parallelize
		doSeedLookup(&results, v)
	}

	return funk.MinInt(results)
}
