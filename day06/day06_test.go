package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		datastreamBuffer: "bvwbjplbgvbhsrlpgdmjqwftvncz",
	}
	input := `bvwbjplbgvbhsrlpgdmjqwftvncz`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		datastreamBuffer: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "7", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		datastreamBuffer: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "", answer)
}

func TestGetFirstStartOfPacketMarkerShould(t *testing.T) {
	tests := map[string]struct {
		datastreamBuffer datastreamBuffer
		expected         int
	}{
		"work with example 1": {
			datastreamBuffer: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected:         5,
		},
		"work with example 2": {
			datastreamBuffer: "nppdvjthqldpwncqszvftbrmjlhg",
			expected:         6,
		},
		"work with example 3": {
			datastreamBuffer: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected:         10,
		},
		"work with example 4": {
			datastreamBuffer: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected:         11,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.datastreamBuffer.getFirstStartOfPacketMarker())
		})
	}
}

func TestGetFirstStartOfMessageMarkerShould(t *testing.T) {
	tests := map[string]struct {
		datastreamBuffer datastreamBuffer
		expected         int
	}{
		"work with example 1": {
			datastreamBuffer: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected:         23,
		},
		"work with example 2": {
			datastreamBuffer: "nppdvjthqldpwncqszvftbrmjlhg",
			expected:         23,
		},
		"work with example 3": {
			datastreamBuffer: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected:         29,
		},
		"work with example 4": {
			datastreamBuffer: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected:         26,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.datastreamBuffer.getFirstStartOfMessageMarker())
		})
	}
}
