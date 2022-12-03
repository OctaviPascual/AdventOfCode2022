package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Day is the interface that wraps SolvePartOne and SolvePartTwo methods
type Day interface {
	SolvePartOne() (string, error)
	SolvePartTwo() (string, error)
}

var days = []struct {
	filename    string
	constructor func(input string) (Day, error)
}{
	{
		filename: "./day01/day01.txt",
		constructor: func(input string) (Day, error) {
			return day01.NewDay(input)
		},
	},
}

func main() {
	for i, day := range days {
		fmt.Printf("\nRunning day %d\n", i+1)
		bytes, err := os.ReadFile(day.filename)
		if err != nil {
			log.Fatalf("could not read file %s: %v", day.filename, err)
		}
		input := string(bytes)
		input = strings.TrimSuffix(input, "\n")

		day, err := day.constructor(input)
		if err != nil {
			log.Fatalf("could not create day %d: %v", i+1, err)
		}

		answer, err := day.SolvePartOne()
		if err != nil {
			log.Fatalf("could not solve part one for day %d: %v", i+1, err)
		}
		fmt.Printf("Part One: %s\n", answer)

		answer, err = day.SolvePartTwo()
		if err != nil {
			log.Fatalf("could not solve part two for day %d: %v", i+1, err)
		}
		fmt.Printf("Part Two: %s\n", answer)
	}
}
