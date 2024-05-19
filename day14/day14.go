package day14

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	cave map[position]material
}

type position struct {
	x, y int
}

type material rune

const (
	rock material = '#'
	sand material = 'o'
)

var (
	sandSource = position{x: 500, y: 0}
)

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	pathsString := strings.Split(input, "\n")

	cave, err := parsePaths(pathsString)
	if err != nil {
		return nil, fmt.Errorf("could not parse paths: %w", err)
	}

	return &Day{cave: cave}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	d.pourSand()
	unitsOfRestingSand := getUnitsOfRestingSand(d.cave)
	return fmt.Sprintf("%d", unitsOfRestingSand), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	d.pourSandWithFloor()
	unitsOfRestingSand := getUnitsOfRestingSand(d.cave)
	return fmt.Sprintf("%d", unitsOfRestingSand), nil
}

func parsePaths(pathsString []string) (map[position]material, error) {
	cave := make(map[position]material)
	for _, pathString := range pathsString {
		err := parsePath(cave, pathString)
		if err != nil {
			return nil, fmt.Errorf("coud not parse path: %w", err)
		}
	}
	return cave, nil
}

func parsePath(cave map[position]material, pathString string) error {
	positionString := strings.Split(pathString, " -> ")
	for i := 0; i < len(positionString)-1; i++ {
		err := parseLineOfRock(cave, positionString[i], positionString[i+1])
		if err != nil {
			return fmt.Errorf("coud not parse line of rock: %w", err)
		}
	}
	return nil
}

func parseLineOfRock(cave map[position]material, startString, endString string) error {
	start, err := parsePosition(startString)
	if err != nil {
		return fmt.Errorf("coud not parse start position: %w", err)
	}

	end, err := parsePosition(endString)
	if err != nil {
		return fmt.Errorf("coud not parse end position: %w", err)
	}

	if start.x == end.x {
		minY := min(start.y, end.y)
		maxY := max(start.y, end.y)
		for y := minY; y <= maxY; y++ {
			cave[position{x: start.x, y: y}] = rock
		}
		return nil
	}

	if start.y == end.y {
		minX := min(start.x, end.x)
		maxX := max(start.x, end.x)
		for x := minX; x <= maxX; x++ {
			cave[position{x: x, y: start.y}] = rock
		}
		return nil
	}

	return fmt.Errorf("start (%s) and end (%s) don't form a straight line", startString, endString)
}

func parsePosition(positionString string) (position, error) {
	xy := strings.Split(positionString, ",")

	if len(xy) != 2 {
		return position{}, fmt.Errorf("invalid position format: %s", positionString)
	}

	x, err := strconv.Atoi(xy[0])
	if err != nil {
		return position{}, fmt.Errorf("invalid x value: %w", err)
	}

	y, err := strconv.Atoi(xy[1])
	if err != nil {
		return position{}, fmt.Errorf("invalid y value: %w", err)
	}

	return position{x: x, y: y}, nil
}

func (d Day) pourSand() {
	maxY := getMaxY(d.cave)
	sandPosition := sandSource
	for !isFlowingIntoAbyss(sandPosition, maxY) {
		nextPosition := fall(sandPosition, d.cave)
		if nextPosition == sandPosition {
			d.cave[sandPosition] = sand
			sandPosition = sandSource
			continue
		}
		sandPosition = nextPosition
	}
}

func (d Day) pourSandWithFloor() {
	maxY := getMaxY(d.cave)
	sandPosition := sandSource
	for d.cave[sandSource] != sand {
		nextPosition := fall(sandPosition, d.cave)
		if hasReachedFloor(nextPosition, maxY) || nextPosition == sandPosition {
			d.cave[sandPosition] = sand
			sandPosition = sandSource
			continue
		}
		sandPosition = nextPosition
	}
}

func fall(sandPosition position, cave map[position]material) position {
	down := position{x: sandPosition.x, y: sandPosition.y + 1}
	if _, ok := cave[down]; !ok {
		return down
	}

	downLeft := position{x: sandPosition.x - 1, y: sandPosition.y + 1}
	if _, ok := cave[downLeft]; !ok {
		return downLeft
	}

	downRight := position{x: sandPosition.x + 1, y: sandPosition.y + 1}
	if _, ok := cave[downRight]; !ok {
		return downRight
	}

	return sandPosition
}

func getMaxY(cave map[position]material) int {
	maxY := math.MinInt
	for position := range cave {
		maxY = max(maxY, position.y)
	}
	return maxY
}

func getUnitsOfRestingSand(cave map[position]material) int {
	unitsOfRestingSand := 0
	for _, m := range cave {
		if m == sand {
			unitsOfRestingSand++
		}
	}
	return unitsOfRestingSand
}

func isFlowingIntoAbyss(sandPosition position, maxY int) bool {
	return sandPosition.y >= maxY
}

func hasReachedFloor(sandPosition position, maxY int) bool {
	return sandPosition.y == maxY+2
}
