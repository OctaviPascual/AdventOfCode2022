package day24

import (
	"fmt"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2022/util"
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

type expedition struct {
	valley            [][]rune
	blizzards         []*blizzard
	current           position
	goal              position
	minutes           int
	minDistanceToGoal int
}

type position struct {
	i, j int
}

type blizzard struct {
	direction rune
	position  position
}

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
	e := newExpedition(d.valley)
	bfs(e)

	return fmt.Sprintf("%d", bfs(e)), nil
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

func newExpedition(valley [][]rune) expedition {
	start := 0
	for valley[0][start] == wall {
		start++
	}

	goal := 0
	for valley[len(valley)-1][goal] == wall {
		goal++
	}

	var blizzards []*blizzard
	for i := 0; i < len(valley); i++ {
		for j := 0; j < len(valley[0]); j++ {
			if isBlizzard(valley[i][j]) {
				b := blizzard{
					direction: valley[i][j],
					position:  position{i: i, j: j},
				}
				blizzards = append(blizzards, &b)
			}
		}
	}

	return expedition{
		valley:            valley,
		blizzards:         blizzards,
		current:           position{i: 0, j: start},
		goal:              position{i: len(valley) - 1, j: goal},
		minDistanceToGoal: position{i: 0, j: start}.manhattanDistance(position{i: len(valley) - 1, j: goal}),
	}
}

func bfs(e expedition) int {
	expeditions := []expedition{e}

	for len(expeditions) > 0 {
		e, expeditions = expeditions[0], expeditions[1:]

		blizzards := e.moveBlizzards()
		blockedPositions := e.getBlockedPositions()
		for _, nextPos := range e.availablePositions(blockedPositions) {
			if e.goal == nextPos {
				return e.minutes
			}

			distanceToGoal := e.current.manhattanDistance(e.goal)
			if distanceToGoal > e.minDistanceToGoal+1 {
				continue
			}

			e := expedition{
				valley:            e.valley,
				blizzards:         blizzards,
				current:           nextPos,
				goal:              e.goal,
				minutes:           e.minutes + 1,
				minDistanceToGoal: min(e.minDistanceToGoal, distanceToGoal),
			}
			expeditions = append(expeditions, e)
		}
	}
	panic("BUG! impossible to reach goal")
}

func (e expedition) moveBlizzards() []*blizzard {
	blizzards := make([]*blizzard, 0, len(e.blizzards))
	for _, b := range e.blizzards {
		blizzards = append(blizzards, &blizzard{
			direction: b.direction,
			position:  b.move(e.valley),
		})
	}
	return blizzards
}

func (e expedition) getBlockedPositions() [][]bool {
	var blockedPositions = make([][]bool, 0, len(e.valley))
	for i := range len(e.valley) {
		blockedPositions = append(blockedPositions, make([]bool, len(e.valley[i])))
	}

	for i := 0; i < len(e.valley); i++ {
		for j := 0; j < len(e.valley[0]); j++ {
			if e.valley[i][j] == wall {
				blockedPositions[i][j] = true
			}
		}
	}

	for _, b := range e.blizzards {
		blockedPositions[b.position.i][b.position.j] = true
	}

	return blockedPositions
}

func (e expedition) availablePositions(blockedPositions [][]bool) []position {
	var availablePositions []position

	down := position{i: e.current.i + 1, j: e.current.j}
	if e.isValid(down) && !blockedPositions[down.i][down.j] {
		availablePositions = append(availablePositions, down)
	}

	wait := e.current
	if e.isValid(wait) && !blockedPositions[wait.i][wait.j] {
		availablePositions = append(availablePositions, wait)
	}

	up := position{i: e.current.i - 1, j: e.current.j}
	if e.isValid(up) && !blockedPositions[up.i][up.j] {
		availablePositions = append(availablePositions, up)
	}

	right := position{i: e.current.i, j: e.current.j + 1}
	if e.isValid(right) && !blockedPositions[right.i][right.j] {
		availablePositions = append(availablePositions, right)
	}

	left := position{i: e.current.i, j: e.current.j - 1}
	if e.isValid(left) && !blockedPositions[left.i][left.j] {
		availablePositions = append(availablePositions, left)
	}

	return availablePositions
}

func (e expedition) isValid(p position) bool {
	rows := len(e.valley)
	columns := len(e.valley[0])
	return p.i >= 0 && p.i < rows && p.j >= 0 && p.j < columns
}

func (b blizzard) move(valley [][]rune) position {
	rows := len(valley)
	columns := len(valley[0])

	switch b.direction {
	case upBlizzard:
		nextPos := position{i: (b.position.i - 1 + rows) % rows, j: b.position.j}
		if valley[nextPos.i][nextPos.j] == wall {
			b.position = nextPos
			nextPos = b.move(valley)
		}
		return nextPos
	case downBlizzard:
		nextPos := position{i: (b.position.i + 1) % rows, j: b.position.j}
		if valley[nextPos.i][nextPos.j] == wall {
			b.position = nextPos
			nextPos = b.move(valley)
		}
		return nextPos
	case rightBlizzard:
		nextPos := position{i: b.position.i, j: (b.position.j + 1) % columns}
		if valley[nextPos.i][nextPos.j] == wall {
			b.position = nextPos
			nextPos = b.move(valley)
		}
		return nextPos
	case leftBlizzard:
		nextPos := position{i: b.position.i, j: (b.position.j - 1 + columns) % columns}
		if valley[nextPos.i][nextPos.j] == wall {
			b.position = nextPos
			nextPos = b.move(valley)
		}
		return nextPos
	}
	panic("BUG! invalid blizzard direction")
}

func isBlizzard(r rune) bool {
	return r == upBlizzard || r == downBlizzard || r == rightBlizzard || r == leftBlizzard
}

func (p position) manhattanDistance(other position) int {
	return util.Abs(p.i-other.i) + util.Abs(p.j-other.j)
}
