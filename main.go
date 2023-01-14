package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2022/day01"
	"github.com/OctaviPascual/AdventOfCode2022/day02"
	"github.com/OctaviPascual/AdventOfCode2022/day03"
	"github.com/OctaviPascual/AdventOfCode2022/day04"
	"github.com/OctaviPascual/AdventOfCode2022/day05"
	"github.com/OctaviPascual/AdventOfCode2022/day06"
	"github.com/OctaviPascual/AdventOfCode2022/day07"
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
	{
		filename: "./day02/day02.txt",
		constructor: func(input string) (Day, error) {
			return day02.NewDay(input)
		},
	},
	{
		filename: "./day03/day03.txt",
		constructor: func(input string) (Day, error) {
			return day03.NewDay(input)
		},
	},
	{
		filename: "./day04/day04.txt",
		constructor: func(input string) (Day, error) {
			return day04.NewDay(input)
		},
	},
	{
		filename: "./day05/day05.txt",
		constructor: func(input string) (Day, error) {
			return day05.NewDay(input)
		},
	},
	{
		filename: "./day06/day06.txt",
		constructor: func(input string) (Day, error) {
			return day06.NewDay(input)
		},
	},
	{
		filename: "./day07/day07.txt",
		constructor: func(input string) (Day, error) {
			return day07.NewDay(input)
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
