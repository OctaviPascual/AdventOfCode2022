package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		valley: [][]rune{
			{wall, ground, wall, wall, wall, wall, wall},
			{wall, ground, ground, ground, ground, ground, wall},
			{wall, rightBlizzard, ground, ground, ground, ground, wall},
			{wall, ground, ground, ground, ground, ground, wall},
			{wall, ground, ground, ground, downBlizzard, ground, wall},
			{wall, ground, ground, ground, ground, ground, wall},
			{wall, wall, wall, wall, wall, ground, wall},
		},
	}
	input := `#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "", answer)
}
