package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/solarwolf-code/aoc2023/day1"
	"github.com/solarwolf-code/aoc2023/day2"
	"github.com/solarwolf-code/aoc2023/day3"
	"github.com/solarwolf-code/aoc2023/day4"
	"github.com/solarwolf-code/aoc2023/day5"
)

func readInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Could not read file: %s", filename)
	}

	return string(content)
}

func main() {
	fmt.Println("Welcome to AOC 2023!\n--------------------")
	daySelected := flag.Int("day", 0, "day of aoc")
	partSelected := flag.Int("part", 0, "part 1 or 2 of day")
	flag.Parse()

	content := readInput(fmt.Sprintf("day%d/day%d.txt", *daySelected, *daySelected))
	switch *daySelected {
	case 1:
		day1.Day1(content, *partSelected)
	case 2:
		day2.Day2(content, *partSelected)
	case 3:
		day3.Day3(content, *partSelected)
	case 4:
		day4.Day4(content, *partSelected)
	case 5:
		day5.Day5(content, *partSelected)
	default:
		log.Fatalf("Day %d not found!", *daySelected)
	}
}
