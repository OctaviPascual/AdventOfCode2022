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
	input := `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "18", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "", answer)
}

func TestBlizzardMoveShould(t *testing.T) {
	valley := [][]rune{
		{wall, ground, wall, wall, wall, wall, wall},
		{wall, ground, ground, ground, ground, ground, wall},
		{wall, ground, ground, ground, ground, ground, wall},
		{wall, ground, ground, ground, ground, ground, wall},
		{wall, ground, ground, ground, ground, ground, wall},
		{wall, ground, ground, ground, ground, ground, wall},
		{wall, wall, wall, wall, wall, ground, wall},
	}

	t.Run("move right blizzard one step to the right", func(t *testing.T) {
		b := blizzard{
			direction: rightBlizzard,
			position:  position{i: 2, j: 1},
		}

		assert.Equal(t, position{i: 2, j: 2}, b.move(valley))
	})

	t.Run("move down blizzard to the opposite side", func(t *testing.T) {
		b := blizzard{
			direction: downBlizzard,
			position:  position{i: 5, j: 2},
		}

		assert.Equal(t, position{i: 1, j: 2}, b.move(valley))
	})

	t.Run("move down blizzard to the opposite side (beginning current)", func(t *testing.T) {
		b := blizzard{
			direction: downBlizzard,
			position:  position{i: 5, j: 1},
		}

		assert.Equal(t, position{i: 0, j: 1}, b.move(valley))
	})

	t.Run("move right blizzard to the opposite side", func(t *testing.T) {
		b := blizzard{
			direction: rightBlizzard,
			position:  position{i: 1, j: 5},
		}

		assert.Equal(t, position{i: 1, j: 1}, b.move(valley))
	})
}
