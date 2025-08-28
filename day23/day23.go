package day23

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2022/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	elves []*elf
}

type grove struct {
	elves      []*elf
	directions []direction
}

type elf struct {
	position position
}

type position struct {
	i, j int
}

type direction string

const (
	north direction = "N"
	south direction = "S"
	west  direction = "W"
	east  direction = "E"
)

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	lines := strings.Split(input, "\n")

	elves, err := parseElves(lines)
	if err != nil {
		return nil, fmt.Errorf("could not parse elves: %w", err)
	}

	return &Day{
		elves: elves,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	grove := newGrove(d.elves)
	for i := 0; i < 10; i++ {
		grove.executeRound()
	}
	return fmt.Sprintf("%d", grove.emptyGroundTiles()), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	grove := newGrove(d.elves)
	rounds := 1
	for grove.executeRound() > 0 {
		rounds++
	}
	return fmt.Sprintf("%d", rounds), nil
}

func parseElves(lines []string) ([]*elf, error) {
	var elves []*elf
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				elves = append(elves, &elf{position{i: i, j: j}})
			}
		}
	}
	return elves, nil
}

func newGrove(elves []*elf) *grove {
	return &grove{
		elves:      elves,
		directions: []direction{north, south, west, east},
	}
}

func (g *grove) executeRound() int {
	proposedPositions := g.executeFirstHalf()
	moves := g.executeSecondHalf(proposedPositions)
	g.rotateListOfDirections()
	return moves
}

func (g *grove) executeFirstHalf() map[position][]*elf {
	proposedPositions := make(map[position][]*elf)
	for _, elf := range g.elves {
		proposedPosition := elf.proposePosition(g.elves, g.directions)
		if proposedPosition != nil {
			proposedPositions[*proposedPosition] = append(proposedPositions[*proposedPosition], elf)
		}
	}
	return proposedPositions
}

func (g *grove) executeSecondHalf(proposedPositions map[position][]*elf) int {
	moves := 0
	for p, elves := range proposedPositions {
		if len(elves) == 1 {
			elves[0].position = p
			moves++
		}
	}
	return moves
}

func (g *grove) rotateListOfDirections() {
	g.directions = append(g.directions[1:4], g.directions[0])
}

func (g *grove) emptyGroundTiles() int {
	minI := slices.MinFunc(g.elves, func(a, b *elf) int { return cmp.Compare(a.position.i, b.position.i) }).position.i
	maxI := slices.MaxFunc(g.elves, func(a, b *elf) int { return cmp.Compare(a.position.i, b.position.i) }).position.i
	minJ := slices.MinFunc(g.elves, func(a, b *elf) int { return cmp.Compare(a.position.j, b.position.j) }).position.j
	maxJ := slices.MaxFunc(g.elves, func(a, b *elf) int { return cmp.Compare(a.position.j, b.position.j) }).position.j
	return (maxI-minI+1)*(maxJ-minJ+1) - len(g.elves)
}

func (e *elf) proposePosition(elves []*elf, directions []direction) *position {
	adjacentPositions := e.getAdjacentPositions()
	adjacentElvesPositions := e.getAdjacentElvesPositions(adjacentPositions, elves)

	if len(adjacentElvesPositions) == 0 {
		return nil
	}

	for _, d := range directions {
		noElves := true
		for _, p := range d.adjacentPositions(e.position) {
			if adjacentElvesPositions.Contains(p) {
				noElves = false
				break
			}
		}
		if noElves {
			proposedPosition := d.move(e.position)
			return &proposedPosition
		}
	}
	return nil
}

func (d direction) move(p position) position {
	switch d {
	case north:
		return position{p.i - 1, p.j}
	case south:
		return position{p.i + 1, p.j}
	case west:
		return position{p.i, p.j - 1}
	case east:
		return position{p.i, p.j + 1}
	}
	panic("BUG! invalid position")
}

func (d direction) adjacentPositions(p position) []position {
	switch d {
	case north:
		return []position{{i: p.i - 1, j: p.j - 1}, {i: p.i - 1, j: p.j}, {i: p.i - 1, j: p.j + 1}}
	case south:
		return []position{{i: p.i + 1, j: p.j - 1}, {i: p.i + 1, j: p.j}, {i: p.i + 1, j: p.j + 1}}
	case west:
		return []position{{i: p.i - 1, j: p.j - 1}, {i: p.i, j: p.j - 1}, {i: p.i + 1, j: p.j - 1}}
	case east:
		return []position{{i: p.i - 1, j: p.j + 1}, {i: p.i, j: p.j + 1}, {i: p.i + 1, j: p.j + 1}}
	}
	return nil
}

func (e *elf) getAdjacentPositions() util.Set[position] {
	adjacentPositions := util.NewSet[position]()
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			p := position{i: e.position.i + i, j: e.position.j + j}
			if p != e.position {
				adjacentPositions.Add(p)
			}
		}
	}
	return adjacentPositions
}

func (e *elf) getAdjacentElvesPositions(positions util.Set[position], elves []*elf) util.Set[position] {
	adjacentElvesPositions := util.NewSet[position]()
	for _, e := range elves {
		if positions.Contains(e.position) {
			adjacentElvesPositions.Add(e.position)
		}
	}
	return adjacentElvesPositions
}
