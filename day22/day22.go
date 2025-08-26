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

type facing rune

const (
	up    facing = '^'
	down  facing = 'v'
	right facing = '>'
	left  facing = '<'
)

type position struct {
	i, j   int
	facing facing
}

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
	p := startingPosition(d.board)
	for _, s := range d.path {
		p = s.execute(p, d.board)
	}
	return fmt.Sprintf("%d", p.finalPassword()), nil
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

	path := make([]step, 0, len(tilesToMoveMatches)+len(turnsMatches))

	for i := 0; i < len(tilesToMoveMatches); i++ {
		tilesToMove, err := strconv.Atoi(tilesToMoveMatches[i][1])
		if err != nil {
			return nil, fmt.Errorf("could not parse tiles to move: %w", err)
		}
		path = append(path, step{tilesToMove: tilesToMove})

		// There is one more tile to move than turn in the path.
		if i < len(tilesToMoveMatches)-1 {
			path = append(path, step{turn: turn(turnsMatches[i][1][0])})
		}

	}

	return path, nil
}

func startingPosition(board [][]square) position {
	j := 0
	for board[0][j] != tile {
		j++
	}

	return position{
		i:      0,
		j:      j,
		facing: right,
	}
}

func (s step) execute(p position, board [][]square) position {
	if s.turn == noTurn {
		return executeMove(p, board, s.tilesToMove)
	}
	return executeTurn(p, s.turn)
}

func executeMove(p position, board [][]square, tilesToMove int) position {
	if tilesToMove <= 0 {
		return p
	}
	nextP := nextPosition(p, board)
	switch board[nextP.i][nextP.j] {
	case tile:
		return executeMove(nextP, board, tilesToMove-1)
	case wall:
		return p
	}
	panic("BUG! next position can't be empty")
}

func executeTurn(p position, t turn) position {
	switch t {
	case clockwise:
		return position{
			i:      p.i,
			j:      p.j,
			facing: p.facing.turnClockwise(),
		}
	case counterClockwise:
		return position{
			i:      p.i,
			j:      p.j,
			facing: p.facing.turnCounterClockwise(),
		}
	}
	panic("BUG! invalid turn")
}

func (f facing) turnClockwise() facing {
	switch f {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	panic("BUG! invalid facing")
}

func (f facing) turnCounterClockwise() facing {
	return f.turnClockwise().turnClockwise().turnClockwise()
}

func (f facing) code() int {
	switch f {
	case right:
		return 0
	case down:
		return 1
	case left:
		return 2
	case up:
		return 3
	}
	panic("BUG! invalid facing")
}

func nextPosition(p position, board [][]square) position {
	rows := len(board)
	columns := len(board[0])

	switch p.facing {
	case up:
		p.i = (p.i - 1 + rows) % rows
	case down:
		p.i = (p.i + 1) % rows
	case left:
		p.j = (p.j - 1 + columns) % columns
	case right:
		p.j = (p.j + 1) % columns
	}
	// We can't move to an empty position, so we keep moving in the same direction.
	if board[p.i][p.j] == empty {
		return nextPosition(p, board)
	}
	return p
}

func (p position) finalPassword() int {
	return 1000*(p.i+1) + 4*(p.j+1) + p.facing.code()
}
