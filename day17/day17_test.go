package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{jetPattern: ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"}
	input := `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{jetPattern: ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "3068", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{jetPattern: ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "", answer)
}
