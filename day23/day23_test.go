package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		elves: []*elf{
			{position: position{i: 1, j: 2}},
			{position: position{i: 1, j: 3}},
			{position: position{i: 2, j: 2}},
			{position: position{i: 4, j: 2}},
			{position: position{i: 4, j: 3}},
		},
	}
	input := `.....
..##.
..#..
.....
..##.
.....`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	input := `..............
..............
.......#......
.....###.#....
...#...#.#....
....#...##....
...#.###......
...##.#.##....
....#..#......
..............
..............
..............`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "110", answer)
}

func TestSolvePartTwo(t *testing.T) {
	input := `..............
..............
.......#......
.....###.#....
...#...#.#....
....#...##....
...#.###......
...##.#.##....
....#..#......
..............
..............
..............`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "20", answer)
}

func TestAdjacentPositions(t *testing.T) {
	e := elf{position: position{i: 10, j: 5}}

	expectedAdjacentPositions := []position{
		{i: 9, j: 4},
		{i: 9, j: 5},
		{i: 9, j: 6},
		{i: 10, j: 4},
		{i: 10, j: 6},
		{i: 11, j: 4},
		{i: 11, j: 5},
		{i: 11, j: 6},
	}
	assert.ElementsMatch(t, expectedAdjacentPositions, e.adjacentPositions())

	var d direction
	d = north

	expectedAdjacentPositions = []position{
		{i: 9, j: 4},
		{i: 9, j: 5},
		{i: 9, j: 6},
	}
	assert.ElementsMatch(t, expectedAdjacentPositions, d.adjacentPositions(position{i: 10, j: 5}))
}
