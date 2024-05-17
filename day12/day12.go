package day12

import (
	"fmt"
	"math"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2022/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	heightmap [][]elevation
}

type elevation rune

type positionState struct {
	position position
	steps    int
}

type position struct {
	i, j int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	heightmapString := strings.Split(input, "\n")

	heightmap, err := parseHeightmap(heightmapString)
	if err != nil {
		return nil, fmt.Errorf("could not parse heightmap: %w", err)
	}

	return &Day{
		heightmap: heightmap,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	startingPosition, err := d.getStartingPosition()
	if err != nil {
		return "", fmt.Errorf("could not get starting position: %w", err)
	}

	steps, err := d.stepsToFinalPosition(startingPosition)
	if err != nil {
		return "", fmt.Errorf("could not get number of steps: %w", err)
	}

	return fmt.Sprintf("%d", steps), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	startingPositions := d.getStartingPositions()

	minSteps := math.MaxInt
	for _, startingPosition := range startingPositions {
		steps, err := d.stepsToFinalPosition(startingPosition)
		if err != nil {
			continue
		}

		minSteps = min(minSteps, steps)
	}

	if minSteps == math.MaxInt {
		return "", fmt.Errorf("final position is unreachable")
	}

	return fmt.Sprintf("%d", minSteps), nil
}

func parseHeightmap(heightmapString []string) ([][]elevation, error) {
	heightmap := make([][]elevation, 0, len(heightmapString))
	for _, elevations := range heightmapString {
		row := make([]elevation, 0, len(elevations))
		for _, elevationRune := range elevations {
			row = append(row, elevation(elevationRune))
		}
		heightmap = append(heightmap, row)
	}

	return heightmap, nil
}

func (d Day) stepsToFinalPosition(startingPosition position) (int, error) {
	visited := util.NewSet[position]()
	queue := []positionState{{position: startingPosition, steps: 0}}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if visited.Contains(state.position) {
			continue
		}
		visited.Add(state.position)

		if d.isFinalPosition(state.position) {
			return state.steps, nil
		}

		nextPositions := d.getNextPositions(state)
		queue = append(queue, nextPositions...)
	}
	return 0, fmt.Errorf("could not reach final position")
}

func (d Day) getNextPositions(state positionState) []positionState {
	var nextPositions []positionState

	steps := state.steps + 1

	up := position{i: state.position.i - 1, j: state.position.j}
	if d.isAccessible(state.position, up) {
		nextPositions = append(nextPositions, positionState{position: up, steps: steps})
	}

	down := position{i: state.position.i + 1, j: state.position.j}
	if d.isAccessible(state.position, down) {
		nextPositions = append(nextPositions, positionState{position: down, steps: steps})
	}

	right := position{i: state.position.i, j: state.position.j + 1}
	if d.isAccessible(state.position, right) {
		nextPositions = append(nextPositions, positionState{position: right, steps: steps})
	}

	left := position{i: state.position.i, j: state.position.j - 1}
	if d.isAccessible(state.position, left) {
		nextPositions = append(nextPositions, positionState{position: left, steps: steps})
	}

	return nextPositions
}

func (d Day) isAccessible(current position, next position) bool {
	n := len(d.heightmap)
	m := len(d.heightmap[0])
	if next.i < 0 || next.i >= n || next.j < 0 || next.j >= m {
		return false
	}

	if d.isStartingPosition(current) {
		return elevation('a')+1 >= d.heightmap[next.i][next.j]
	}

	if d.isFinalPosition(next) {
		return d.heightmap[current.i][current.j]+1 >= elevation('z')
	}

	return d.heightmap[current.i][current.j]+1 >= d.heightmap[next.i][next.j]
}

func (d Day) isStartingPosition(position position) bool {
	return d.isPosition(position, elevation('S'))
}

func (d Day) isFinalPosition(position position) bool {
	return d.isPosition(position, elevation('E'))
}

func (d Day) isPosition(position position, elevation elevation) bool {
	n := len(d.heightmap)
	m := len(d.heightmap[0])

	if position.i < 0 || position.i >= n || position.j < 0 || position.j >= m {
		return false
	}

	return d.heightmap[position.i][position.j] == elevation
}

func (d Day) getStartingPosition() (position, error) {
	n := len(d.heightmap)
	m := len(d.heightmap[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			position := position{i: i, j: j}
			if d.isStartingPosition(position) {
				return position, nil
			}
		}
	}
	return position{}, fmt.Errorf("starting position S not found")
}

func (d Day) getStartingPositions() []position {
	n := len(d.heightmap)
	m := len(d.heightmap[0])

	var startingPositions []position
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			position := position{i: i, j: j}
			if d.isStartingPosition(position) || d.heightmap[i][j] == 'a' {
				startingPositions = append(startingPositions, position)
			}
		}
	}
	return startingPositions
}
