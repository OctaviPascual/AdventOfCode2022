package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func buildItems(itemsString string) []item {
	items := make([]item, 0, len(itemsString))
	for _, itemString := range itemsString {
		items = append(items, item{id: itemString})
	}
	return items
}

func TestNewDay(t *testing.T) {
	expected := &Day{
		rucksacks: []rucksack{
			{
				firstCompartmentItems:  buildItems("vJrwpWtwJgWr"),
				secondCompartmentItems: buildItems("hcsFMMfFFhFp"),
			},
			{
				firstCompartmentItems:  buildItems("jqHRNqRjqzjGDLGL"),
				secondCompartmentItems: buildItems("rsFMfFZSrLrFZsSL"),
			},
			{
				firstCompartmentItems:  buildItems("PmmdzqPrV"),
				secondCompartmentItems: buildItems("vPwwTWBwg"),
			},
			{
				firstCompartmentItems:  buildItems("wMqvLMZHhHMvwLH"),
				secondCompartmentItems: buildItems("jbvcjnnSBnvTQFn"),
			},
			{
				firstCompartmentItems:  buildItems("ttgJtRGJ"),
				secondCompartmentItems: buildItems("QctTZtZT"),
			},
			{
				firstCompartmentItems:  buildItems("CrZsJsPPZsGz"),
				secondCompartmentItems: buildItems("wwsLwLmpwMDw"),
			},
		},
	}
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "157", answer)
}

func TestSolvePartTwo(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "70", answer)
}

func TestItemPriority(t *testing.T) {
	assert.Equal(t, 1, item{id: 'a'}.priority())
	assert.Equal(t, 26, item{id: 'z'}.priority())
	assert.Equal(t, 27, item{id: 'A'}.priority())
	assert.Equal(t, 52, item{id: 'Z'}.priority())
}
