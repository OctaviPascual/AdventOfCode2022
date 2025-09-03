package day24

import (
	"fmt"
	"strings"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	valley [][]rune
}

const (
	upBlizzard    = '^'
	downBlizzard  = 'v'
	rightBlizzard = '>'
	leftBlizzard  = '<'
	wall          = '#'
	ground        = '.'
)

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	lines := strings.Split(input, "\n")

	valley, err := parseValley(lines)
	if err != nil {
		return nil, fmt.Errorf("could not parse valley: %w", err)
	}

	return &Day{
		valley: valley,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	return "", nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	return "", nil
}

func parseValley(lines []string) ([][]rune, error) {
	var valley = make([][]rune, 0, len(lines))
	for _, line := range lines {
		valley = append(valley, []rune(line))
	}
	return valley, nil
}
