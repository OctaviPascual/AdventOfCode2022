package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	stack1 := stack{
		id:     1,
		crates: []crate{'Z', 'N'},
	}
	stack2 := stack{
		id:     2,
		crates: []crate{'M', 'C', 'D'},
	}
	stack3 := stack{
		id:     3,
		crates: []crate{'P'},
	}
	expected := &Day{
		stacks: map[stackID]*stack{
			1: &stack1,
			2: &stack2,
			3: &stack3,
		},
		rearrangementProcedure: []step{
			{
				crates: 1,
				from:   2,
				to:     1,
			},
			{
				crates: 3,
				from:   1,
				to:     3,
			},
			{
				crates: 2,
				from:   2,
				to:     1,
			},
			{
				crates: 1,
				from:   1,
				to:     2,
			},
		},
	}
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "CMZ", answer)
}

func TestSolvePartTwo(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "MCD", answer)
}
