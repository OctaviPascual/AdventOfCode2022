package day17

import (
	"errors"
	"fmt"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	jetPattern string
}

const (
	chamberWidth  = 7
	chamberHeight = 100_000

	air  rune = '.'
	rock rune = '#'
	wall rune = '|'
)

var (
	errRockLanded = errors.New("rock landed")
)

type shape interface {
	appear(int, map[position]rune) position
	fallDown(position, map[position]rune) (position, error)
	pushRight(position, map[position]rune) position
	pushLeft(position, map[position]rune) position
}

type position struct {
	x, y int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	return &Day{
		jetPattern: input,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	const totalRocks = 2022

	chamber := createChamber()
	shapes := []shape{shape1{}, shape2{}, shape3{}, shape4{}, shape5{}}
	yMax := -1
	i := 0
	for rock := 0; rock < totalRocks; rock++ {
		s := shapes[rock%5]
		rockPosition := s.appear(yMax, chamber)
		for {
			if d.jetPattern[i%len(d.jetPattern)] == '>' {
				rockPosition = s.pushRight(rockPosition, chamber)
			} else {
				rockPosition = s.pushLeft(rockPosition, chamber)
			}
			i++
			landing, err := s.fallDown(rockPosition, chamber)
			if errors.Is(err, errRockLanded) {
				yMax = max(yMax, landing.y)
				break
			}
			rockPosition = landing
		}
	}
	// we must add 1 to yMax because we use 0-based coordinates
	return fmt.Sprintf("%d", yMax+1), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	// TODO: Figure out how to find a recurring pattern in the layout to compute arbitrarily big total rocks.
	return "", nil
}

func createChamber() map[position]rune {
	chamber := make(map[position]rune)

	for x := -1; x < chamberWidth; x++ {
		chamber[position{x: x, y: -1}] = wall
	}
	for y := -1; y <= chamberHeight; y++ {
		chamber[position{x: -1, y: y}] = wall
		chamber[position{x: chamberWidth, y: y}] = wall
	}
	for x := 0; x < chamberWidth; x++ {
		for y := 0; y < chamberHeight; y++ {
			chamber[position{x: x, y: y}] = air
		}
	}
	return chamber
}

// Represents shape with form:
// ####
type shape1 struct{}

func (s shape1) appear(yMax int, chamber map[position]rune) position {
	p := position{x: 2, y: yMax + 4}
	chamber[position{x: p.x, y: p.y}] = rock
	chamber[position{x: p.x + 1, y: p.y}] = rock
	chamber[position{x: p.x + 2, y: p.y}] = rock
	chamber[position{x: p.x + 3, y: p.y}] = rock
	return p
}

func (s shape1) fallDown(p position, chamber map[position]rune) (position, error) {
	down1 := position{x: p.x, y: p.y - 1}
	down2 := position{x: p.x + 1, y: p.y - 1}
	down3 := position{x: p.x + 2, y: p.y - 1}
	down4 := position{x: p.x + 3, y: p.y - 1}
	if chamber[down1] == air && chamber[down2] == air && chamber[down3] == air && chamber[down4] == air {
		chamber[position{x: p.x, y: p.y}] = air
		chamber[position{x: p.x + 1, y: p.y}] = air
		chamber[position{x: p.x + 2, y: p.y}] = air
		chamber[position{x: p.x + 3, y: p.y}] = air

		chamber[down1] = rock
		chamber[down2] = rock
		chamber[down3] = rock
		chamber[down4] = rock

		return position{x: p.x, y: p.y - 1}, nil
	}
	return p, errRockLanded
}

func (s shape1) pushRight(p position, chamber map[position]rune) position {
	right1 := position{x: p.x + 4, y: p.y}
	if chamber[right1] == air {
		chamber[position{x: p.x, y: p.y}] = air
		chamber[right1] = rock
		return position{x: p.x + 1, y: p.y}
	}
	return p
}

func (s shape1) pushLeft(p position, chamber map[position]rune) position {
	left1 := position{x: p.x - 1, y: p.y}
	if chamber[left1] == air {
		chamber[position{x: p.x + 3, y: p.y}] = air
		chamber[left1] = rock
		return position{x: p.x - 1, y: p.y}
	}
	return p
}

// Represents shape with form:
// .#.
// ###
// .#.
type shape2 struct{}

func (s shape2) appear(yMax int, chamber map[position]rune) position {
	p := position{x: 2, y: yMax + 6}
	chamber[position{x: p.x + 1, y: p.y}] = rock
	chamber[position{x: p.x, y: p.y - 1}] = rock
	chamber[position{x: p.x + 1, y: p.y - 1}] = rock
	chamber[position{x: p.x + 2, y: p.y - 1}] = rock
	chamber[position{x: p.x + 1, y: p.y - 2}] = rock
	return p
}

func (s shape2) fallDown(p position, chamber map[position]rune) (position, error) {
	down1 := position{x: p.x, y: p.y - 2}
	down2 := position{x: p.x + 1, y: p.y - 3}
	down3 := position{x: p.x + 2, y: p.y - 2}

	if chamber[down1] == air && chamber[down2] == air && chamber[down3] == air {
		chamber[position{x: p.x, y: p.y - 1}] = air
		chamber[position{x: p.x + 1, y: p.y}] = air
		chamber[position{x: p.x + 2, y: p.y - 1}] = air

		chamber[down1] = rock
		chamber[down2] = rock
		chamber[down3] = rock

		return position{x: p.x, y: p.y - 1}, nil
	}
	return p, errRockLanded
}

func (s shape2) pushRight(p position, chamber map[position]rune) position {
	right1 := position{x: p.x + 2, y: p.y}
	right2 := position{x: p.x + 3, y: p.y - 1}
	right3 := position{x: p.x + 2, y: p.y - 2}

	if chamber[right1] == air && chamber[right2] == air && chamber[right3] == air {
		chamber[position{x: p.x + 1, y: p.y}] = air
		chamber[position{x: p.x, y: p.y - 1}] = air
		chamber[position{x: p.x + 1, y: p.y - 2}] = air

		chamber[right1] = rock
		chamber[right2] = rock
		chamber[right3] = rock

		return position{x: p.x + 1, y: p.y}
	}
	return p
}

func (s shape2) pushLeft(p position, chamber map[position]rune) position {
	left1 := position{x: p.x, y: p.y}
	left2 := position{x: p.x - 1, y: p.y - 1}
	left3 := position{x: p.x, y: p.y - 2}

	if chamber[left1] == air && chamber[left2] == air && chamber[left3] == air {
		chamber[position{x: p.x + 1, y: p.y}] = air
		chamber[position{x: p.x + 2, y: p.y - 1}] = air
		chamber[position{x: p.x + 1, y: p.y - 2}] = air

		chamber[left1] = rock
		chamber[left2] = rock
		chamber[left3] = rock

		return position{x: p.x - 1, y: p.y}
	}
	return p
}

// Represents shape of with form:
// ..#
// ..#
// ###
type shape3 struct{}

func (s shape3) appear(yMax int, chamber map[position]rune) position {
	p := position{x: 2, y: yMax + 6}
	chamber[position{x: p.x + 2, y: p.y}] = rock
	chamber[position{x: p.x + 2, y: p.y - 1}] = rock
	chamber[position{x: p.x, y: p.y - 2}] = rock
	chamber[position{x: p.x + 1, y: p.y - 2}] = rock
	chamber[position{x: p.x + 2, y: p.y - 2}] = rock
	return p
}

func (s shape3) fallDown(p position, chamber map[position]rune) (position, error) {
	down1 := position{x: p.x, y: p.y - 3}
	down2 := position{x: p.x + 1, y: p.y - 3}
	down3 := position{x: p.x + 2, y: p.y - 3}

	if chamber[down1] == air && chamber[down2] == air && chamber[down3] == air {
		chamber[position{x: p.x + 2, y: p.y}] = air
		chamber[position{x: p.x, y: p.y - 2}] = air
		chamber[position{x: p.x + 1, y: p.y - 2}] = air

		chamber[down1] = rock
		chamber[down2] = rock
		chamber[down3] = rock

		return position{x: p.x, y: p.y - 1}, nil
	}
	return p, errRockLanded
}

func (s shape3) pushRight(p position, chamber map[position]rune) position {
	right1 := position{x: p.x + 3, y: p.y}
	right2 := position{x: p.x + 3, y: p.y - 1}
	right3 := position{x: p.x + 3, y: p.y - 2}

	if chamber[right1] == air && chamber[right2] == air && chamber[right3] == air {
		chamber[position{x: p.x + 2, y: p.y}] = air
		chamber[position{x: p.x + 2, y: p.y - 1}] = air
		chamber[position{x: p.x, y: p.y - 2}] = air

		chamber[right1] = rock
		chamber[right2] = rock
		chamber[right3] = rock

		return position{x: p.x + 1, y: p.y}
	}
	return p
}

func (s shape3) pushLeft(p position, chamber map[position]rune) position {
	left1 := position{x: p.x + 1, y: p.y}
	left2 := position{x: p.x + 1, y: p.y - 1}
	left3 := position{x: p.x - 1, y: p.y - 2}

	if chamber[left1] == air && chamber[left2] == air && chamber[left3] == air {
		chamber[position{x: p.x + 2, y: p.y}] = air
		chamber[position{x: p.x + 2, y: p.y - 1}] = air
		chamber[position{x: p.x + 2, y: p.y - 2}] = air

		chamber[left1] = rock
		chamber[left2] = rock
		chamber[left3] = rock

		return position{x: p.x - 1, y: p.y}
	}
	return p
}

// Represents shape with form:
// #
// #
// #
// #
type shape4 struct{}

func (s shape4) appear(yMax int, chamber map[position]rune) position {
	p := position{x: 2, y: yMax + 7}
	chamber[position{x: p.x, y: p.y}] = rock
	chamber[position{x: p.x, y: p.y - 1}] = rock
	chamber[position{x: p.x, y: p.y - 2}] = rock
	chamber[position{x: p.x, y: p.y - 3}] = rock
	return p
}

func (s shape4) fallDown(p position, chamber map[position]rune) (position, error) {
	down1 := position{x: p.x, y: p.y - 4}
	if chamber[down1] == air {
		chamber[position{x: p.x, y: p.y}] = air

		chamber[down1] = rock

		return position{x: p.x, y: p.y - 1}, nil
	}
	return p, errRockLanded
}

func (s shape4) pushRight(p position, chamber map[position]rune) position {
	right1 := position{x: p.x + 1, y: p.y}
	right2 := position{x: p.x + 1, y: p.y - 1}
	right3 := position{x: p.x + 1, y: p.y - 2}
	right4 := position{x: p.x + 1, y: p.y - 3}
	if chamber[right1] == air && chamber[right2] == air && chamber[right3] == air && chamber[right4] == air {
		chamber[position{x: p.x, y: p.y}] = air
		chamber[position{x: p.x, y: p.y - 1}] = air
		chamber[position{x: p.x, y: p.y - 2}] = air
		chamber[position{x: p.x, y: p.y - 3}] = air

		chamber[right1] = rock
		chamber[right2] = rock
		chamber[right3] = rock
		chamber[right4] = rock
		return position{x: p.x + 1, y: p.y}
	}
	return p
}

func (s shape4) pushLeft(p position, chamber map[position]rune) position {
	left1 := position{x: p.x - 1, y: p.y}
	left2 := position{x: p.x - 1, y: p.y - 1}
	left3 := position{x: p.x - 1, y: p.y - 2}
	left4 := position{x: p.x - 1, y: p.y - 3}
	if chamber[left1] == air && chamber[left2] == air && chamber[left3] == air && chamber[left4] == air {
		chamber[position{x: p.x, y: p.y}] = air
		chamber[position{x: p.x, y: p.y - 1}] = air
		chamber[position{x: p.x, y: p.y - 2}] = air
		chamber[position{x: p.x, y: p.y - 3}] = air

		chamber[left1] = rock
		chamber[left2] = rock
		chamber[left3] = rock
		chamber[left4] = rock
		return position{x: p.x - 1, y: p.y}
	}
	return p
}

// Represents shape with form:
// ##
// ##
type shape5 struct{}

func (s shape5) appear(yMax int, chamber map[position]rune) position {
	p := position{x: 2, y: yMax + 5}
	chamber[position{x: p.x, y: p.y}] = rock
	chamber[position{x: p.x + 1, y: p.y}] = rock
	chamber[position{x: p.x, y: p.y - 1}] = rock
	chamber[position{x: p.x + 1, y: p.y - 1}] = rock
	return p
}

func (s shape5) fallDown(p position, chamber map[position]rune) (position, error) {
	down1 := position{x: p.x, y: p.y - 2}
	down2 := position{x: p.x + 1, y: p.y - 2}
	if chamber[down1] == air && chamber[down2] == air {
		chamber[position{x: p.x, y: p.y}] = air
		chamber[position{x: p.x + 1, y: p.y}] = air

		chamber[down1] = rock
		chamber[down2] = rock

		return position{x: p.x, y: p.y - 1}, nil
	}
	return p, errRockLanded
}

func (s shape5) pushRight(p position, chamber map[position]rune) position {
	right1 := position{x: p.x + 2, y: p.y}
	right2 := position{x: p.x + 2, y: p.y - 1}
	if chamber[right1] == air && chamber[right2] == air {
		chamber[position{x: p.x, y: p.y}] = air
		chamber[position{x: p.x, y: p.y - 1}] = air
		chamber[right1] = rock
		chamber[right2] = rock
		return position{x: p.x + 1, y: p.y}
	}
	return p
}

func (s shape5) pushLeft(p position, chamber map[position]rune) position {
	left1 := position{x: p.x - 1, y: p.y}
	left2 := position{x: p.x - 1, y: p.y - 1}
	if chamber[left1] == air && chamber[left2] == air {
		chamber[position{x: p.x + 1, y: p.y}] = air
		chamber[position{x: p.x + 1, y: p.y - 1}] = air
		chamber[left1] = rock
		chamber[left2] = rock
		return position{x: p.x - 1, y: p.y}
	}
	return p
}
