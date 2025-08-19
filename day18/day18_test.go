package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		cubes: []cube{
			{1, 1, 1},
			{2, 1, 1},
		},
	}
	input := `1,1,1
2,1,1`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	input := `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "64", answer)
}

func TestSolvePartTwo(t *testing.T) {
	input := `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "58", answer)
}
