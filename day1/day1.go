package day1

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func isDigit(char rune) bool {
	if _, err := strconv.Atoi(string(char)); err == nil {
		return true
	} else {
		return false
	}
}

func part1(text string) {
	sum := 0
	// split text on each line
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		first := ""
		last := ""
		for _, char := range line {
			if isDigit(char) && first == "" {
				first = string(char)
			}
			if isDigit(char) {
				last = string(char)
			}
		}
		num, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatalf("Failed to convert num to int")
		}
		sum += num
	}
	fmt.Printf("Part 1: %d\n", sum)
}

type PossibleMatch struct {
	startIndex int
	num        string
}

func findAllSubstrings(str, substr string) []int {
	indices := []int{}
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr {
			indices = append(indices, i)
		}
	}
	return indices
}

func part2(text string) {
	numMappings := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	lines := strings.Split(text, "\n")
	sum := 0
	for _, line := range lines {
		// what if we check where the substring index starts for all keys in the line and we pick the first one
		possibleMatches := []PossibleMatch{}
		for k := range numMappings {
			subStrings := findAllSubstrings(line, k)
			for _, index := range subStrings {
				// fmt.Println(index, k)
				possibleMatches = append(possibleMatches, PossibleMatch{index, k})
			}
		}
		// order by lowest to highest index
		sort.Slice(possibleMatches, func(i, j int) bool {
			return possibleMatches[i].startIndex < possibleMatches[j].startIndex
		})

		// eight:  len of 5

		ogLen := len(line)
		// fix line with corrections
		for _, match := range possibleMatches {
			lastChar := string(match.num[len(match.num)-1])
			line = strings.Replace(line, match.num, numMappings[match.num]+lastChar, 1)
		}
		first := ""
		last := ""

		lenOffset := ogLen - len(line)

		for charIdx, char := range line {
			if isDigit(char) && first == "" {
				first = string(char)
			}
			if isDigit(char) {
				last = string(char)
				if len(possibleMatches) > 0 && possibleMatches[len(possibleMatches)-1].startIndex > charIdx+lenOffset {
					last = numMappings[possibleMatches[len(possibleMatches)-1].num]
				} else {
					last = string(char)
				}
			}
		}
		num, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatalf("Failed to convert num to int")
		}
		sum += num
	}
	fmt.Printf("Part 2: %d\n", sum)
} // 54194

func Day1(text string) {
	fmt.Println("Running Day 1 Part 1\n------------")
	part1(text)
	fmt.Println("\nRunning Day 1 Part 2\n------------")
	part2(text)
}
