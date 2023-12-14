package day5

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"math"
)

func convertStringSliceToInt(stringSlice []string) []int {
	intSlice := []int{}
	for _, val := range stringSlice {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("Failed to convert val to int")
		}
		intSlice = append(intSlice, valInt)
	}
	return intSlice
}

func part1(text string) {
	lines := strings.Split(text, "\n\n")
	locations := []int{}
	seeds := convertStringSliceToInt(strings.Split(strings.Split(lines[0], "seeds: ")[1], " "))
	categories := []string{}
	for _, line := range lines {
		if strings.Contains(line, "map") {
			categories = append(categories, line)
		}
	}
	for _, categoryReturn := range seeds {
		for _, category := range categories {
			categoryMapping := -1
			for _, mapping := range strings.Split(category, "\n")[1:] {
				mappingVals := convertStringSliceToInt(strings.Split(mapping, " "))
				if categoryReturn >= mappingVals[1] && categoryReturn < mappingVals[1]+mappingVals[2] {
					categoryMapping = mappingVals[0] + (categoryReturn - mappingVals[1])
				}
			}
			if categoryMapping == -1 {
				categoryMapping = categoryReturn
			}
			categoryReturn = categoryMapping
		}
		locations = append(locations, categoryReturn)
	}
	lowestLocation := slices.Min(locations)
	fmt.Printf("Part 1: %d\n", lowestLocation)
}

func part2(text string) {
	lines := strings.Split(text, "\n\n")
	locations := []int{}
	seeds := convertStringSliceToInt(strings.Split(strings.Split(lines[0], "seeds: ")[1], " "))
	categories := []string{}
	for _, line := range lines {
		if strings.Contains(line, "map") {
			categories = append(categories, line)
		}
	}


	rangeSeeds := []int{}
	start := 0
	end := 0
	for idx, seed := range seeds {
		if math.Mod(float64(idx+1), 2) == 0{
			end = seed
			for i := 0; i < end; i++ {
				rangeSeeds = append(rangeSeeds, start+i)
			}
		} else {
			start = seed
		}
	}

	for _, categoryReturn := range rangeSeeds {
		for _, category := range categories {
			categoryMapping := -1
			for _, mapping := range strings.Split(category, "\n")[1:] {
				mappingVals := convertStringSliceToInt(strings.Split(mapping, " "))
				if categoryReturn >= mappingVals[1] && categoryReturn < mappingVals[1]+mappingVals[2] {
					categoryMapping = mappingVals[0] + (categoryReturn - mappingVals[1])
					break
				}
			}
			if categoryMapping == -1 {
				categoryMapping = categoryReturn
			}
			categoryReturn = categoryMapping
		}
		locations = append(locations, categoryReturn)
	}
	lowestLocation := slices.Min(locations)
	fmt.Printf("Part 2: %d\n", lowestLocation)
}

func Day5(text string, part int) {
	switch part {
	case 0:
		part1(text)
		part2(text)
	case 1:
		part1(text)
	case 2:
		part2(text)
	default:
		log.Fatalf("Part %d not found", part)
	}
}
