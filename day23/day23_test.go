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
}

func TestDirectionsAdjacentPositionsShould(t *testing.T) {
	tests := map[string]struct {
		input    direction
		expected []position
	}{
		"return N, NE, NW positions when direction is north": {
			input: north,
			expected: []position{
				{i: 9, j: 4},
				{i: 9, j: 5},
				{i: 9, j: 6},
			},
		},
		"return S, SE, SW positions when direction is south": {
			input: south,
			expected: []position{
				{i: 11, j: 4},
				{i: 11, j: 5},
				{i: 11, j: 6},
			},
		},
		"return W, NW, SW positions when direction is west": {
			input: west,
			expected: []position{
				{i: 9, j: 4},
				{i: 10, j: 4},
				{i: 11, j: 4},
			},
		},
		"return E, NE, SE positions when direction is east": {
			input: east,
			expected: []position{
				{i: 9, j: 6},
				{i: 10, j: 6},
				{i: 11, j: 6},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.input.adjacentPositions(position{i: 10, j: 5}))
		})
	}
}

func TestDirectionsMoveShould(t *testing.T) {
	tests := map[string]struct {
		input    direction
		expected position
	}{
		"return N position when direction is north": {
			input:    north,
			expected: position{i: 9, j: 5},
		},
		"return S position when direction is south": {
			input:    south,
			expected: position{i: 11, j: 5},
		},
		"return W position when direction is west": {
			input:    west,
			expected: position{i: 10, j: 4},
		},
		"return E position when direction is east": {
			input:    east,
			expected: position{i: 10, j: 6},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.input.move(position{i: 10, j: 5}))
		})
	}
}
