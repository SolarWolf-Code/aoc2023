package day4

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func isDigit(char rune) bool {
	// fmt.Printf("Symbol %s: ", string(char))
	if _, err := strconv.Atoi(string(char)); err == nil {
		// fmt.Printf("true\n")
		return true
	} else {
		// fmt.Printf("false\n")
		return false
	}
}

func sliceIntersection(slice1, slice2 []int) []int {
	intersection := make([]int, 0)

	// Create a map to store the elements of the first slice
	slice1Map := make(map[int]bool)
	for _, element := range slice1 {
		slice1Map[element] = true
	}

	// Iterate through the second slice and check if each element is in the map
	for _, element := range slice2 {
		if slice1Map[element] {
			intersection = append(intersection, element)
		}
	}

	return intersection
}

func part1(text string) {
	lines := strings.Split(text, "\n")
	sum := 0
	for _, line := range lines {
		cardValues := strings.Split(line, ": ")[1]
		winningNumbers := make([]int, 0)
		haveNumbers := make([]int, 0)
		haveSplit := false
		for _, val := range strings.Split(cardValues, " ") {
			if val == "|" {
				haveSplit = true
				continue
			}
			if val != "" {
				valInt, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal("Failed to convert val to int")
				}

				if !haveSplit {
					// winningNumbers.append(valInt)
					winningNumbers = append(winningNumbers, valInt)
				} else {
					haveNumbers = append(haveNumbers, valInt)

				}
			}
		}
		cardMatches := len(sliceIntersection(winningNumbers, haveNumbers))
		result := 0.5
		for i := 1; i <= cardMatches; i++ {
			result *= 2
		}

		if result != 0.5 {
			sum += int(result)
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func part2(text string) {
	lines := strings.Split(text, "\n")
	sum := 0
	cardCopies := make(map[string]int, 0)
	for cardId, line := range lines {
		repeatCount := 1
		if val, ok := cardCopies[fmt.Sprintf("%d", cardId)]; ok {
			repeatCount += val
		}

		for i := 0; i < repeatCount; i++ {
			cardValues := strings.Split(line, ": ")[1]
			winningNumbers := make([]int, 0)
			haveNumbers := make([]int, 0)
			haveSplit := false
			for _, val := range strings.Split(cardValues, " ") {
				if val == "|" {
					haveSplit = true
					continue
				}
				if val != "" {
					valInt, err := strconv.Atoi(val)
					if err != nil {
						log.Fatal("Failed to convert val to int")
					}

					if !haveSplit {
						winningNumbers = append(winningNumbers, valInt)
					} else {
						haveNumbers = append(haveNumbers, valInt)

					}
				}
			}
			cardMatches := len(sliceIntersection(winningNumbers, haveNumbers))
			for i := 1; i <= cardMatches; i++ {
				// fmt.Printf("Incrementing %d\n", cardId+i)
				cardCopies[fmt.Sprintf("%d", cardId+i)] += 1
			}
		}
		sum += cardCopies[fmt.Sprintf("%d", cardId)] + 1
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func Day4(text string) {
	fmt.Println("Running Day 4 Part 1\n------------")
	part1(text)
	fmt.Println("\nRunning Day 4 Part 2\n------------")
	part2(text)
}
