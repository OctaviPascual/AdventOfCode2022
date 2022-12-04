package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		strategyGuide: []round{
			{column1: "A", column2: "Y"},
			{column1: "B", column2: "X"},
			{column1: "C", column2: "Z"},
		},
	}
	input := `A Y
B X
C Z`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		strategyGuide: []round{
			{column1: "A", column2: "Y"},
			{column1: "B", column2: "X"},
			{column1: "C", column2: "Z"},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "15", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		strategyGuide: []round{
			{column1: "A", column2: "Y"},
			{column1: "B", column2: "X"},
			{column1: "C", column2: "Z"},
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "12", answer)
}
