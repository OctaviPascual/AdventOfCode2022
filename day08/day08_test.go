package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		grid: [][]tree{
			{{3}, {0}, {3}, {7}, {3}},
			{{2}, {5}, {5}, {1}, {2}},
			{{6}, {5}, {3}, {3}, {2}},
			{{3}, {3}, {5}, {4}, {9}},
			{{3}, {5}, {3}, {9}, {0}},
		},
	}
	input := `30373
25512
65332
33549
35390`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		grid: [][]tree{
			{{3}, {0}, {3}, {7}, {3}},
			{{2}, {5}, {5}, {1}, {2}},
			{{6}, {5}, {3}, {3}, {2}},
			{{3}, {3}, {5}, {4}, {9}},
			{{3}, {5}, {3}, {9}, {0}},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "21", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		grid: [][]tree{
			{{3}, {0}, {3}, {7}, {3}},
			{{2}, {5}, {5}, {1}, {2}},
			{{6}, {5}, {3}, {3}, {2}},
			{{3}, {3}, {5}, {4}, {9}},
			{{3}, {5}, {3}, {9}, {0}},
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "8", answer)
}
