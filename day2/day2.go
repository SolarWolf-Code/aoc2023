package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func part1(text string) {
	lines := strings.Split(text, "\n")
	colorRestrictions := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	for _, line := range lines {
		sep := strings.Split(line, ": ")
		gameIdStr := strings.Replace(sep[0], "Game ", "", 1)
		gameId, err := strconv.Atoi(gameIdStr)
		if err != nil {
			log.Fatal("Could not convert gameId to int")
		}
		gameSets := strings.Split(sep[1], "; ")
		possible := true
		for _, gameSet := range gameSets {
			setCubes := strings.Split(gameSet, ", ")
			for _, setCube := range setCubes {
				// get color and amount
				sep := strings.Split(setCube, " ")
				amtStr := sep[0]
				amt, err := strconv.Atoi(amtStr)
				if err != nil {
					log.Fatal("Could not convert amtStr to int")
				}
				color := sep[1]
				if colorRestrictions[color] < amt {
					possible = false
				}
			}
		}
		if possible {
			sum += gameId
		}

	}
	fmt.Printf("Part 1: %d\n", sum)

}

func part2(text string) {
	lines := strings.Split(text, "\n")

	sum := 0
	for _, line := range lines {
		sep := strings.Split(line, ": ")
		gameSets := strings.Split(sep[1], "; ")
		redMin := 0
		greenMin := 0
		blueMin := 0
		for _, gameSet := range gameSets {
			setCubes := strings.Split(gameSet, ", ")
			for _, setCube := range setCubes {
				// get color and amount
				sep := strings.Split(setCube, " ")
				amtStr := sep[0]
				amt, err := strconv.Atoi(amtStr)
				if err != nil {
					log.Fatal("Could not convert amtStr to int")
				}
				color := sep[1]
				switch color {
				case "red":
					if amt > redMin || redMin == 0 {
						redMin = amt
					}
				case "green":
					if amt > greenMin || greenMin == 0 {
						greenMin = amt
					}
				case "blue":
					if amt > blueMin || blueMin == 0 {
						blueMin = amt
					}
				}
			}
		}
		sum += redMin * greenMin * blueMin
	}
	fmt.Printf("Part 2: %d\n", sum)
}

func Day2(text string) {
	fmt.Println("Running Day 2 Part 1\n------------")
	part1(text)
	fmt.Println("\nRunning Day 2 Part 2\n------------")
	part2(text)
}
