package day09

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2022/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	motions []motion
}

type motion struct {
	direction direction
	steps     int
}

type direction string

const (
	right direction = "R"
	left  direction = "L"
	up    direction = "U"
	down  direction = "D"
)

type ropeBridge struct {
	headPosition         position
	tailPosition         position
	tailVisitedPositions util.Set[position]
}

type position struct {
	x, y int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	motionsString := strings.Split(input, "\n")

	motions, err := parseMotions(motionsString)
	if err != nil {
		return nil, fmt.Errorf("could not parse motions: %w", err)
	}

	return &Day{
		motions: motions,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	ropeBridge := newRopeBridge()
	ropeBridge.runSimulation(d.motions)
	return fmt.Sprintf("%d", ropeBridge.totalTailVisitedPositions()), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	return "", nil
}

func parseMotions(motionsString []string) ([]motion, error) {
	motions := make([]motion, 0, len(motionsString))
	for _, motionString := range motionsString {
		motion, err := parseMotion(motionString)
		if err != nil {
			return nil, fmt.Errorf("could not parse motion: %w", err)
		}

		motions = append(motions, motion)
	}
	return motions, nil
}

func parseMotion(motionString string) (motion, error) {
	splitMotion := strings.Split(motionString, " ")
	if len(splitMotion) != 2 {
		return motion{}, fmt.Errorf("invalid format of motion: %s", motionString)
	}

	direction, err := parseDirection(splitMotion[0])
	if err != nil {
		return motion{}, fmt.Errorf("could not parse direction: %w", err)
	}

	steps, err := strconv.Atoi(splitMotion[1])
	if err != nil {
		return motion{}, fmt.Errorf("could not parse steps: %w", err)
	}

	return motion{
		direction: direction,
		steps:     steps,
	}, nil
}

func parseDirection(directionString string) (direction, error) {
	switch directionString {
	case "R":
		return right, nil
	case "L":
		return left, nil
	case "U":
		return up, nil
	case "D":
		return down, nil
	}
	return "", fmt.Errorf("invalid direction: %s", directionString)
}

func newRopeBridge() *ropeBridge {
	initialPosition := position{0, 0}

	tailVisitedPositions := util.Set[position]{}
	tailVisitedPositions.Add(initialPosition)

	return &ropeBridge{
		headPosition:         initialPosition,
		tailPosition:         initialPosition,
		tailVisitedPositions: util.Set[position]{},
	}
}

func (r *ropeBridge) runSimulation(motions []motion) {
	for _, motion := range motions {
		for i := 0; i < motion.steps; i++ {
			r.headPosition.move(motion.direction)
			r.updateTailPosition()
			r.tailVisitedPositions.Add(r.tailPosition)
		}
	}
}

func (r *ropeBridge) totalTailVisitedPositions() int {
	return len(r.tailVisitedPositions)
}

func (r *ropeBridge) updateTailPosition() {
	r.tailPosition.follow(r.headPosition)
}

func (p *position) move(direction direction) {
	switch direction {
	case right:
		p.x++
	case left:
		p.x--
	case up:
		p.y++
	case down:
		p.y--
	}
}

func (p *position) follow(p2 position) {
	offsetX := p2.x - p.x
	offsetY := p2.y - p.y

	distanceX := int(math.Abs(float64(offsetX)))
	distanceY := int(math.Abs(float64(offsetY)))

	if distanceX > 1 || (distanceX > 0 && distanceY > 1) {
		p.x += 1 * (offsetX / distanceX)
	}

	if distanceY > 1 || (distanceY > 0 && distanceX > 1) {
		p.y += 1 * (offsetY / distanceY)
	}
}
