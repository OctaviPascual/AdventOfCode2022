package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		cave: map[position]material{
			// 498,4 -> 498,6
			position{x: 498, y: 4}: rock, position{x: 498, y: 5}: rock, position{x: 498, y: 6}: rock,
			// 498,6 -> 496,6
			position{x: 496, y: 6}: rock, position{x: 497, y: 6}: rock, position{x: 498, y: 6}: rock,
			// 503,4 -> 502,4
			position{x: 502, y: 4}: rock, position{x: 503, y: 4}: rock,
			// 502,4 -> 502,9
			position{x: 502, y: 4}: rock, position{x: 502, y: 5}: rock, position{x: 502, y: 6}: rock,
			position{x: 502, y: 7}: rock, position{x: 502, y: 8}: rock, position{x: 502, y: 9}: rock,
			// 502,9 -> 494,9
			position{x: 494, y: 9}: rock, position{x: 495, y: 9}: rock, position{x: 496, y: 9}: rock,
			position{x: 497, y: 9}: rock, position{x: 498, y: 9}: rock, position{x: 499, y: 9}: rock,
			position{x: 500, y: 9}: rock, position{x: 501, y: 9}: rock, position{x: 502, y: 9}: rock,
		},
	}
	input := `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		cave: map[position]material{
			// 498,4 -> 498,6
			position{x: 498, y: 4}: rock, position{x: 498, y: 5}: rock, position{x: 498, y: 6}: rock,
			// 498,6 -> 496,6
			position{x: 496, y: 6}: rock, position{x: 497, y: 6}: rock, position{x: 498, y: 6}: rock,
			// 503,4 -> 502,4
			position{x: 502, y: 4}: rock, position{x: 503, y: 4}: rock,
			// 502,4 -> 502,9
			position{x: 502, y: 4}: rock, position{x: 502, y: 5}: rock, position{x: 502, y: 6}: rock,
			position{x: 502, y: 7}: rock, position{x: 502, y: 8}: rock, position{x: 502, y: 9}: rock,
			// 502,9 -> 494,9
			position{x: 494, y: 9}: rock, position{x: 495, y: 9}: rock, position{x: 496, y: 9}: rock,
			position{x: 497, y: 9}: rock, position{x: 498, y: 9}: rock, position{x: 499, y: 9}: rock,
			position{x: 500, y: 9}: rock, position{x: 501, y: 9}: rock, position{x: 502, y: 9}: rock,
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "24", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		cave: map[position]material{
			// 498,4 -> 498,6
			position{x: 498, y: 4}: rock, position{x: 498, y: 5}: rock, position{x: 498, y: 6}: rock,
			// 498,6 -> 496,6
			position{x: 496, y: 6}: rock, position{x: 497, y: 6}: rock, position{x: 498, y: 6}: rock,
			// 503,4 -> 502,4
			position{x: 502, y: 4}: rock, position{x: 503, y: 4}: rock,
			// 502,4 -> 502,9
			position{x: 502, y: 4}: rock, position{x: 502, y: 5}: rock, position{x: 502, y: 6}: rock,
			position{x: 502, y: 7}: rock, position{x: 502, y: 8}: rock, position{x: 502, y: 9}: rock,
			// 502,9 -> 494,9
			position{x: 494, y: 9}: rock, position{x: 495, y: 9}: rock, position{x: 496, y: 9}: rock,
			position{x: 497, y: 9}: rock, position{x: 498, y: 9}: rock, position{x: 499, y: 9}: rock,
			position{x: 500, y: 9}: rock, position{x: 501, y: 9}: rock, position{x: 502, y: 9}: rock,
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "93", answer)
}
