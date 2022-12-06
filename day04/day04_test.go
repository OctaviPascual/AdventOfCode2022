package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		assignments: []assignment{
			{
				firstElf:  sectionIDRange{2, 4},
				secondElf: sectionIDRange{6, 8},
			},
			{
				firstElf:  sectionIDRange{2, 3},
				secondElf: sectionIDRange{4, 5},
			},
			{
				firstElf:  sectionIDRange{5, 7},
				secondElf: sectionIDRange{7, 9},
			},
			{
				firstElf:  sectionIDRange{2, 8},
				secondElf: sectionIDRange{3, 7},
			},
			{
				firstElf:  sectionIDRange{6, 6},
				secondElf: sectionIDRange{4, 6},
			},
			{
				firstElf:  sectionIDRange{2, 6},
				secondElf: sectionIDRange{4, 8},
			},
		},
	}
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "2", answer)
}

func TestSolvePartTwo(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "4", answer)
}
