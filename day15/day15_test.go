package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{}
	input := ``
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
