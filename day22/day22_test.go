package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		board: [][]square{
			{empty, empty, empty, empty, empty, empty, empty, empty, tile, tile, tile, wall, empty, empty, empty, empty},
			{empty, empty, empty, empty, empty, empty, empty, empty, tile, wall, tile, tile, empty, empty, empty, empty},
			{empty, empty, empty, empty, empty, empty, empty, empty, wall, tile, tile, tile, empty, empty, empty, empty},
			{empty, empty, empty, empty, empty, empty, empty, empty, tile, tile, tile, tile, empty, empty, empty, empty},

			{tile, tile, tile, wall, tile, tile, tile, tile, tile, tile, tile, wall, empty, empty, empty, empty},
			{tile, tile, tile, tile, tile, tile, tile, tile, wall, tile, tile, tile, empty, empty, empty, empty},
			{tile, tile, wall, tile, tile, tile, tile, wall, tile, tile, tile, tile, empty, empty, empty, empty},
			{tile, tile, tile, tile, tile, tile, tile, tile, tile, tile, wall, tile, empty, empty, empty, empty},

			{empty, empty, empty, empty, empty, empty, empty, empty, tile, tile, tile, wall, tile, tile, tile, tile},
			{empty, empty, empty, empty, empty, empty, empty, empty, tile, tile, tile, tile, tile, wall, tile, tile},
			{empty, empty, empty, empty, empty, empty, empty, empty, tile, wall, tile, tile, tile, tile, tile, tile},
			{empty, empty, empty, empty, empty, empty, empty, empty, tile, tile, tile, tile, tile, tile, wall, tile},
		},
		path: []step{
			{tilesToMove: 10},
			{turn: clockwise},
			{tilesToMove: 5},
			{turn: counterClockwise},
			{tilesToMove: 5},
			{turn: clockwise},
			{tilesToMove: 10},
			{turn: counterClockwise},
			{tilesToMove: 4},
			{turn: clockwise},
			{tilesToMove: 5},
			{turn: counterClockwise},
			{tilesToMove: 5},
		},
	}
	input := `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	input := `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "6032", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "", answer)
}
