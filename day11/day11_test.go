package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		monkeys: []monkey{
			{[]item{{79}, {98}},
				operation{operator: "*", operand: 19},
				test{divisibleBy: 23, monkeyIfTrue: 2, monkeyIfFalse: 3},
			},
			{[]item{{54}, {65}, {75}, {74}},
				operation{operator: "+", operand: 6},
				test{divisibleBy: 19, monkeyIfTrue: 2, monkeyIfFalse: 0},
			},
			{[]item{{79}, {60}, {97}},
				operation{operator: "* old"},
				test{divisibleBy: 13, monkeyIfTrue: 1, monkeyIfFalse: 3},
			},
			{[]item{{74}},
				operation{operator: "+", operand: 3},
				test{divisibleBy: 17, monkeyIfTrue: 0, monkeyIfFalse: 1},
			},
		},
	}
	input := `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		monkeys: []monkey{
			{[]item{{79}, {98}},
				operation{operator: "*", operand: 19},
				test{divisibleBy: 23, monkeyIfTrue: 2, monkeyIfFalse: 3},
			},
			{[]item{{54}, {65}, {75}, {74}},
				operation{operator: "+", operand: 6},
				test{divisibleBy: 19, monkeyIfTrue: 2, monkeyIfFalse: 0},
			},
			{[]item{{79}, {60}, {97}},
				operation{operator: "* old"},
				test{divisibleBy: 13, monkeyIfTrue: 1, monkeyIfFalse: 3},
			},
			{[]item{{74}},
				operation{operator: "+", operand: 3},
				test{divisibleBy: 17, monkeyIfTrue: 0, monkeyIfFalse: 1},
			},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "10605", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		monkeys: []monkey{
			{[]item{{79}, {98}},
				operation{operator: "*", operand: 19},
				test{divisibleBy: 23, monkeyIfTrue: 2, monkeyIfFalse: 3},
			},
			{[]item{{54}, {65}, {75}, {74}},
				operation{operator: "+", operand: 6},
				test{divisibleBy: 19, monkeyIfTrue: 2, monkeyIfFalse: 0},
			},
			{[]item{{79}, {60}, {97}},
				operation{operator: "* old"},
				test{divisibleBy: 13, monkeyIfTrue: 1, monkeyIfFalse: 3},
			},
			{[]item{{74}},
				operation{operator: "+", operand: 3},
				test{divisibleBy: 17, monkeyIfTrue: 0, monkeyIfFalse: 1},
			},
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "2713310158", answer)
}
