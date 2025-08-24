package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		encryptedFile: []int{1, 2, -3, 3, -2, 0, 4},
	}
	input := `1
2
-3
3
-2
0
4`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		encryptedFile: []int{1, 2, -3, 3, -2, 0, 4},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "3", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		encryptedFile: []int{1, 2, -3, 3, -2, 0, 4},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "1623178306", answer)
}
