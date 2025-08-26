package day22

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	board [][]square
	path  []step
}

var (
	// Regex capturing all the tiles to move in "10R5L5R10L4R5L5"
	tilesToMoveRe = regexp.MustCompile(`(\d+)[RL]?`)

	// Regex capturing all the turns in "10R5L5R10L4R5L5"
	turnsRe = regexp.MustCompile(`\d+([RL])`)
)

type position struct {
	i, j int
}

type square rune

const (
	empty square = ' '
	tile  square = '.'
	wall  square = '#'
)

type step struct {
	tilesToMove int
	turn        turn
}

type turn rune

const (
	noTurn           turn = 0
	clockwise        turn = 'R'
	counterClockwise turn = 'L'
)

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	lines := strings.Split(input, "\n")

	n := len(lines)

	board, err := parseBoard(lines[:n-2])
	if err != nil {
		return nil, fmt.Errorf("could not parse board: %w", err)
	}

	path, err := parsePath(lines[n-1])
	if err != nil {
		return nil, fmt.Errorf("could not parse path: %w", err)
	}
	return &Day{board: board, path: path}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	return "", nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	return "", nil
}

func parseBoard(lines []string) ([][]square, error) {
	rows := len(lines)
	longestLine := slices.MaxFunc(lines, func(a, b string) int { return cmp.Compare(len(a), len(b)) })
	columns := len(longestLine)

	board := make([][]square, 0, rows)

	for i := 0; i < rows; i++ {
		row := make([]square, 0, columns)
		for j := 0; j < columns; j++ {
			if j >= len(lines[i]) {
				row = append(row, empty)
				continue
			}
			row = append(row, square(lines[i][j]))
		}
		board = append(board, row)
	}
	return board, nil
}

func parsePath(line string) ([]step, error) {
	tilesToMoveMatches := tilesToMoveRe.FindAllStringSubmatch(line, -1)
	turnsMatches := turnsRe.FindAllStringSubmatch(line, -1)

	path := make([]step, 0, len(tilesToMoveMatches)+len(turnsMatches)-2)

	for i := 0; i < len(tilesToMoveMatches); i++ {
		tilesToMove, err := strconv.Atoi(tilesToMoveMatches[i][1])
		if err != nil {
			return nil, fmt.Errorf("could not parse tiles to move: %w", err)
		}
		path = append(path, step{tilesToMove: tilesToMove})

		// There is one more tile to move than turn in the path.
		if i == len(tilesToMoveMatches)-1 {
			break
		}

		path = append(path, step{turn: turn(turnsMatches[i][1][0])})
	}

	return path, nil
}
