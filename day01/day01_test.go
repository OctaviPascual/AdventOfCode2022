package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		elves: []elf{
			{
				items: []item{{calories: 1_000}, {calories: 2_000}, {calories: 3_000}},
			},
			{
				items: []item{{calories: 4_000}},
			},
			{
				items: []item{{calories: 5_000}, {calories: 6_000}},
			},
			{
				items: []item{{calories: 7_000}, {calories: 8_000}, {calories: 9_000}},
			},
			{
				items: []item{{calories: 10_000}},
			},
		},
	}
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		elves: []elf{
			{
				items: []item{{calories: 1_000}, {calories: 2_000}, {calories: 3_000}},
			},
			{
				items: []item{{calories: 4_000}},
			},
			{
				items: []item{{calories: 5_000}, {calories: 6_000}},
			},
			{
				items: []item{{calories: 7_000}, {calories: 8_000}, {calories: 9_000}},
			},
			{
				items: []item{{calories: 10_000}},
			},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "24000", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		elves: []elf{
			{
				items: []item{{calories: 1_000}, {calories: 2_000}, {calories: 3_000}},
			},
			{
				items: []item{{calories: 4_000}},
			},
			{
				items: []item{{calories: 5_000}, {calories: 6_000}},
			},
			{
				items: []item{{calories: 7_000}, {calories: 8_000}, {calories: 9_000}},
			},
			{
				items: []item{{calories: 10_000}},
			},
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "45000", answer)
}
