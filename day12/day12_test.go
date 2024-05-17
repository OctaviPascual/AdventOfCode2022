package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		[][]elevation{
			{elevation('S'), elevation('a'), elevation('b'), elevation('q'), elevation('p'), elevation('o'), elevation('n'), elevation('m')},
			{elevation('a'), elevation('b'), elevation('c'), elevation('r'), elevation('y'), elevation('x'), elevation('x'), elevation('l')},
			{elevation('a'), elevation('c'), elevation('c'), elevation('s'), elevation('z'), elevation('E'), elevation('x'), elevation('k')},
			{elevation('a'), elevation('c'), elevation('c'), elevation('t'), elevation('u'), elevation('v'), elevation('w'), elevation('j')},
			{elevation('a'), elevation('b'), elevation('d'), elevation('e'), elevation('f'), elevation('g'), elevation('h'), elevation('i')},
		},
	}
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		[][]elevation{
			{elevation('S'), elevation('a'), elevation('b'), elevation('q'), elevation('p'), elevation('o'), elevation('n'), elevation('m')},
			{elevation('a'), elevation('b'), elevation('c'), elevation('r'), elevation('y'), elevation('x'), elevation('x'), elevation('l')},
			{elevation('a'), elevation('c'), elevation('c'), elevation('s'), elevation('z'), elevation('E'), elevation('x'), elevation('k')},
			{elevation('a'), elevation('c'), elevation('c'), elevation('t'), elevation('u'), elevation('v'), elevation('w'), elevation('j')},
			{elevation('a'), elevation('b'), elevation('d'), elevation('e'), elevation('f'), elevation('g'), elevation('h'), elevation('i')},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "31", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		[][]elevation{
			{elevation('S'), elevation('a'), elevation('b'), elevation('q'), elevation('p'), elevation('o'), elevation('n'), elevation('m')},
			{elevation('a'), elevation('b'), elevation('c'), elevation('r'), elevation('y'), elevation('x'), elevation('x'), elevation('l')},
			{elevation('a'), elevation('c'), elevation('c'), elevation('s'), elevation('z'), elevation('E'), elevation('x'), elevation('k')},
			{elevation('a'), elevation('c'), elevation('c'), elevation('t'), elevation('u'), elevation('v'), elevation('w'), elevation('j')},
			{elevation('a'), elevation('b'), elevation('d'), elevation('e'), elevation('f'), elevation('g'), elevation('h'), elevation('i')},
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "29", answer)
}
