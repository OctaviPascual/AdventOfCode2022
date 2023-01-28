package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		motions: []motion{
			{right, 4}, {up, 4}, {left, 3}, {down, 1}, {right, 4}, {down, 1}, {left, 5}, {right, 2},
		},
	}
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		motions: []motion{
			{right, 4}, {up, 4}, {left, 3}, {down, 1}, {right, 4}, {down, 1}, {left, 5}, {right, 2},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "13", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		motions: []motion{
			{right, 5}, {up, 8}, {left, 8}, {down, 3}, {right, 17}, {down, 10}, {left, 25}, {up, 20},
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "36", answer)
}
