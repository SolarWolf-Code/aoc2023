package day3

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

type idxLocation struct {
	x int
	y int
}

type partRange struct {
	number   string
	idxRange []idxLocation
}

func part1(text string) {
	lines := strings.Split(text, "\n")
	// get len of line for width to use for adjacent
	// height can be indx of lines
	lineWidth := len(lines[0])
	sum := 0
	for lineIdx, line := range lines {
		nums := make(map[string][]idxLocation)
		part := partRange{}
		for charIdx, char := range line {
			if isDigit(char) {
				part.number += string(char)
				part.idxRange = append(part.idxRange, idxLocation{x: charIdx, y: lineIdx})
				if charIdx == len(line)-1 {
					nums[fmt.Sprintf("%d-%d|%s", lineIdx, charIdx, part.number)] = part.idxRange
					part = partRange{}
				}

			} else {
				if part.number != "" {
					nums[fmt.Sprintf("%d-%d|%s", lineIdx, charIdx, part.number)] = part.idxRange
					part = partRange{}
				}
			}

		}

		// loop each part number and check for adjacent symbols
		for partNumStr, idxRanges := range nums {
			// fmt.Println(part, idxRanges)
			adjacentSymbol := false
			// loop through each number in part number and check left, right, up, down, and diagonals
			for _, coords := range idxRanges {
				// check left
				if coords.x != 0 && !adjacentSymbol {
					if string(lines[coords.y][coords.x-1]) != "." && !isDigit(rune(lines[coords.y][coords.x-1])) {
						adjacentSymbol = true
						break
					}
				}
				// check right
				if coords.x != lineWidth-1 && !adjacentSymbol {
					if string(lines[coords.y][coords.x+1]) != "." && !isDigit(rune(lines[coords.y][coords.x+1])) {
						adjacentSymbol = true
						break
					}
				}
				// check up
				if coords.y != 0 && !adjacentSymbol {
					if string(lines[coords.y-1][coords.x]) != "." && !isDigit(rune(lines[coords.y-1][coords.x])) {
						adjacentSymbol = true
						break
					}
				}
				// check down
				if coords.y != len(lines)-1 && !adjacentSymbol {
					if string(lines[coords.y+1][coords.x]) != "." && !isDigit(rune(lines[coords.y+1][coords.x])) {
						adjacentSymbol = true
						break
					}
				}

				// check left up
				if coords.y != 0 && coords.x != 0 && !adjacentSymbol {
					if string(lines[coords.y-1][coords.x-1]) != "." && !isDigit(rune(lines[coords.y-1][coords.x-1])) {
						adjacentSymbol = true
						break
					}
				}
				// check right up
				if coords.y != 0 && coords.x != lineWidth-1 && !adjacentSymbol {
					if string(lines[coords.y-1][coords.x+1]) != "." && !isDigit(rune(lines[coords.y-1][coords.x+1])) {
						adjacentSymbol = true
						break
					}
				}
				// check left down
				if coords.y != len(lines)-1 && coords.x != 0 && !adjacentSymbol {
					if string(lines[coords.y+1][coords.x-1]) != "." && !isDigit(rune(lines[coords.y+1][coords.x-1])) {
						adjacentSymbol = true
						break
					}
				}

				// check right down
				if coords.y != len(lines)-1 && coords.x != lineWidth-1 && !adjacentSymbol {
					if string(lines[coords.y+1][coords.x+1]) != "." && !isDigit(rune(lines[coords.y+1][coords.x+1])) {
						adjacentSymbol = true
						break
					}
				}

			}
			if adjacentSymbol {
				partNum, err := strconv.Atoi(strings.Split(partNumStr, "|")[1])
				if err != nil {
					log.Fatal("Failed to convert partNumStr to int")
				}
				sum += partNum
			}
		}
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func part2(text string) {
	lines := strings.Split(text, "\n")
	sum := 0
	adjacentGears := make(map[string][]string) // adjacentGears["0,1"] = [123, 456]
	for lineIdx, line := range lines {
		nums := make(map[string][]idxLocation)
		part := partRange{}
		lineWidth := len(lines[0])
		for charIdx, char := range line {
			if isDigit(char) {
				part.number += string(char)
				part.idxRange = append(part.idxRange, idxLocation{x: charIdx, y: lineIdx})
				if charIdx == len(line)-1 {
					nums[fmt.Sprintf("%d-%d|%s", lineIdx, charIdx, part.number)] = part.idxRange
					part = partRange{}
				}

			} else {
				if part.number != "" {
					nums[fmt.Sprintf("%d-%d|%s", lineIdx, charIdx, part.number)] = part.idxRange
					part = partRange{}
				}
			}

		}

		for partNumStr, idxRanges := range nums {
			// fmt.Println(part, idxRanges)
			adjacentSymbol := false
			// loop through each number in part number and check left, right, up, down, and diagonals
			for _, coords := range idxRanges {
				// check left
				if coords.x != 0 && !adjacentSymbol {
					if string(lines[coords.y][coords.x-1]) == "*" {
						// TODO: replace this logic with adding to map with * index
						gear_loc := fmt.Sprintf("%d,%d", coords.x-1, coords.y)
						adjacentGears[gear_loc] = append(adjacentGears[gear_loc], partNumStr)
						adjacentSymbol = true
						break
					}
				}
				// check right
				if coords.x != lineWidth-1 && !adjacentSymbol {
					if string(lines[coords.y][coords.x+1]) == "*" {
						// TODO: replace this logic with adding to map with * index
						gear_loc := fmt.Sprintf("%d,%d", coords.x+1, coords.y)
						adjacentGears[gear_loc] = append(adjacentGears[gear_loc], partNumStr)
						adjacentSymbol = true
						break
					}
				}
				// check up
				if coords.y != 0 && !adjacentSymbol {
					if string(lines[coords.y-1][coords.x]) == "*" {
						// TODO: replace this logic with adding to map with * index
						gear_loc := fmt.Sprintf("%d,%d", coords.x, coords.y-1)
						adjacentGears[gear_loc] = append(adjacentGears[gear_loc], partNumStr)
						adjacentSymbol = true
						break
					}
				}
				// check down
				if coords.y != len(lines)-1 && !adjacentSymbol {
					if string(lines[coords.y+1][coords.x]) == "*" {
						// TODO: replace this logic with adding to map with * index
						gear_loc := fmt.Sprintf("%d,%d", coords.x, coords.y+1)
						adjacentGears[gear_loc] = append(adjacentGears[gear_loc], partNumStr)
						adjacentSymbol = true
						break
					}
				}

				// check left up
				if coords.y != 0 && coords.x != 0 && !adjacentSymbol {
					if string(lines[coords.y-1][coords.x-1]) == "*" {
						// TODO: replace this logic with adding to map with * index
						gear_loc := fmt.Sprintf("%d,%d", coords.x-1, coords.y-1)
						adjacentGears[gear_loc] = append(adjacentGears[gear_loc], partNumStr)
						adjacentSymbol = true
						break
					}
				}
				// check right up
				if coords.y != 0 && coords.x != lineWidth-1 && !adjacentSymbol {
					if string(lines[coords.y-1][coords.x+1]) == "*" {
						// TODO: replace this logic with adding to map with * index
						gear_loc := fmt.Sprintf("%d,%d", coords.x+1, coords.y-1)
						adjacentGears[gear_loc] = append(adjacentGears[gear_loc], partNumStr)
						adjacentSymbol = true
						break
					}
				}
				// check left down
				if coords.y != len(lines)-1 && coords.x != 0 && !adjacentSymbol {
					if string(lines[coords.y+1][coords.x-1]) == "*" {
						// TODO: replace this logic with adding to map with * index
						gear_loc := fmt.Sprintf("%d,%d", coords.x-1, coords.y+1)
						adjacentGears[gear_loc] = append(adjacentGears[gear_loc], partNumStr)
						adjacentSymbol = true
						break
					}
				}

				// check right down
				if coords.y != len(lines)-1 && coords.x != lineWidth-1 && !adjacentSymbol {
					if string(lines[coords.y+1][coords.x+1]) == "*" {
						// TODO: replace this logic with adding to map with * index
						gear_loc := fmt.Sprintf("%d,%d", coords.x+1, coords.y+1)
						adjacentGears[gear_loc] = append(adjacentGears[gear_loc], partNumStr)
						adjacentSymbol = true
						break
					}
				}

			}
		}

	}
	for _, partsNums := range adjacentGears {
		if len(partsNums) == 2 {
			partNum1, err := strconv.Atoi(strings.Split(partsNums[0], "|")[1])
			if err != nil {
				log.Fatal("Could not convert first partNumStr to int")
			}
			partNum2, err := strconv.Atoi(strings.Split(partsNums[1], "|")[1])
			if err != nil {
				log.Fatal("Could not convert second partNumStr to int")
			}
			sum += partNum1 * partNum2
		}
	}
	fmt.Printf("Part 2: %d\n", sum)

}

func Day3(text string, part int) {
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
