package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/solarwolf-code/aoc2023/day1"
	"github.com/solarwolf-code/aoc2023/day2"
	"github.com/solarwolf-code/aoc2023/day3"
	"github.com/solarwolf-code/aoc2023/day4"
)

func readInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Could not read file: %s", filename)
	}

	return string(content)
}

func main() {
	// fmt.Println("Welcome to AOC 2023!\n--------------------")
	if len(os.Args) < 2 {
		log.Fatal("You need specify the day you wish to run. Such as `go run main.go 1` for Day 1")
	}

	daySelected, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to convert %s to int", os.Args[1])
	}
	// TODO maybe have it so if there is another arg, we can parse the parts. i.e: go run main.go 1 1 would be Day1 part 1
	content := readInput(fmt.Sprintf("day%d/day%d.txt", daySelected, daySelected))
	switch daySelected {
	case 1:
		day1.Day1(content)
	case 2:
		day2.Day2(content)
	case 3:
		day3.Day3(content)
	case 4:
		day4.Day4(content)
	default:
		log.Fatalf("Day %d not found!", daySelected)
	}
}
