package day01

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct {
	elves []elf
}

type elf struct {
	items []item
}

type item struct {
	calories int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	elvesString := strings.Split(input, "\n")

	elves, err := parseElves(elvesString)
	if err != nil {
		return nil, fmt.Errorf("could not parse elves: %w", err)
	}

	return &Day{elves: elves}, nil
}

// SolvePartOne solves part one
func (d *Day) SolvePartOne() (string, error) {
	maxCalories, _ := maxCalories(d.elves)
	return fmt.Sprintf("%d", maxCalories), nil
}

// SolvePartTwo solves part two
func (d *Day) SolvePartTwo() (string, error) {
	top1Calories, position := maxCalories(d.elves)
	top2Calories, position := maxCalories(removeElf(d.elves, position))
	top3Calories, _ := maxCalories(removeElf(d.elves, position))

	return fmt.Sprintf("%d", top1Calories+top2Calories+top3Calories), nil
}

func parseElves(elvesString []string) ([]elf, error) {
	var elves []elf
	i := 0
	for i < len(elvesString) {
		var items []item
		for i < len(elvesString) && elvesString[i] != "" {
			item, err := parseItem(elvesString[i])
			if err != nil {
				return nil, fmt.Errorf("could not parse item: %w", err)
			}

			items = append(items, item)
			i++
		}
		i++
		elves = append(elves, elf{items: items})
	}
	return elves, nil
}

func parseItem(itemString string) (item, error) {
	calories, err := strconv.Atoi(itemString)
	if err != nil {
		return item{}, fmt.Errorf("could not parse calories: %w", err)
	}

	return item{calories: calories}, nil
}

func (e elf) calories() int {
	totalCalories := 0
	for _, item := range e.items {
		totalCalories += item.calories
	}
	return totalCalories
}

func maxCalories(elves []elf) (int, int) {
	maxCalories := 0
	maxCaloriesPosition := 0
	for i, elf := range elves {
		calories := elf.calories()
		if calories > maxCalories {
			maxCalories = calories
			maxCaloriesPosition = i
		}
	}
	return maxCalories, maxCaloriesPosition
}

func removeElf(elves []elf, i int) []elf {
	elves[i] = elves[len(elves)-1]
	return elves[:len(elves)-1]
}
