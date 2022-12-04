package day03

import (
	"fmt"
	"strings"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	rucksacks []rucksack
}

type rucksack struct {
	firstCompartmentItems  []item
	secondCompartmentItems []item
}

type item struct {
	id rune
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	rucksacksString := strings.Split(input, "\n")

	rucksacks, err := parseRucksacks(rucksacksString)
	if err != nil {
		return nil, fmt.Errorf("could not parse rucksacks: %w", err)
	}

	return &Day{
		rucksacks: rucksacks,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	commonItemSum, err := getCommonItemSum(d.rucksacks)
	if err != nil {
		return "", fmt.Errorf("could not get common item sum: %w", err)
	}

	return fmt.Sprintf("%d", commonItemSum), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	badgeSum, err := getBadgeSum(d.rucksacks)
	if err != nil {
		return "", fmt.Errorf("could not get badge sum: %w", err)
	}

	return fmt.Sprintf("%d", badgeSum), nil
}

func parseRucksacks(rucksacksString []string) ([]rucksack, error) {
	if len(rucksacksString)%3 != 0 {
		return nil, fmt.Errorf("invalid number of items: %d", len(rucksacksString))
	}

	rucksacks := make([]rucksack, 0, len(rucksacksString))

	for _, rucksackString := range rucksacksString {
		rucksack, err := parseRucksack(rucksackString)
		if err != nil {
			return nil, fmt.Errorf("could not parse rucksack: %w", err)
		}

		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks, nil
}

func parseRucksack(rucksackString string) (rucksack, error) {
	if len(rucksackString)%2 != 0 {
		return rucksack{}, fmt.Errorf("invalid number of items: %d", len(rucksackString))
	}

	i := len(rucksackString) / 2

	firstCompartmentItems := make([]item, 0, i)
	for _, itemRune := range rucksackString[:i] {
		firstCompartmentItems = append(firstCompartmentItems, item{id: itemRune})
	}

	secondCompartmentItems := make([]item, 0, i)
	for _, itemRune := range rucksackString[i:] {
		secondCompartmentItems = append(secondCompartmentItems, item{id: itemRune})
	}

	return rucksack{
		firstCompartmentItems:  firstCompartmentItems,
		secondCompartmentItems: secondCompartmentItems,
	}, nil
}

func getCommonItemSum(rucksacks []rucksack) (int, error) {
	total := 0
	for _, rucksack := range rucksacks {
		commonItem, err := rucksack.findCommonItem()
		if err != nil {
			return 0, fmt.Errorf("could not find ruckstack common item: %w", err)
		}
		total += commonItem.priority()
	}
	return total, nil
}

func (r rucksack) findCommonItem() (item, error) {
	firstCompartmentItems := make(map[item]struct{})

	for _, item := range r.firstCompartmentItems {
		firstCompartmentItems[item] = struct{}{}
	}

	for _, item := range r.secondCompartmentItems {
		if _, ok := firstCompartmentItems[item]; ok {
			return item, nil
		}
	}

	return item{}, fmt.Errorf("no common item in rucksack compartments")
}

func (i item) priority() int {
	if i.id >= 'a' && i.id <= 'z' {
		return int(i.id) - int('a') + 1
	}

	maxLowerCasePriority := item{id: 'z'}.priority()
	return int(i.id) - int('A') + 1 + maxLowerCasePriority
}

func getBadgeSum(rucksacks []rucksack) (int, error) {
	total := 0
	i := 0
	for i+2 < len(rucksacks) {
		badge, err := findBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		if err != nil {
			return 0, fmt.Errorf("could not find badge: %w", err)
		}

		total += badge.priority()
		i += 3
	}
	return total, nil
}

func findBadge(rucksack1, rucksack2, rucksack3 rucksack) (item, error) {
	commonItems := make(map[item]struct{})

	firstRucksackItems := make(map[item]struct{})
	for _, item := range rucksack1.firstCompartmentItems {
		firstRucksackItems[item] = struct{}{}
	}
	for _, item := range rucksack1.secondCompartmentItems {
		firstRucksackItems[item] = struct{}{}
	}

	for _, item := range rucksack2.firstCompartmentItems {
		if _, ok := firstRucksackItems[item]; ok {
			commonItems[item] = struct{}{}
		}
	}
	for _, item := range rucksack2.secondCompartmentItems {
		if _, ok := firstRucksackItems[item]; ok {
			commonItems[item] = struct{}{}
		}
	}

	for _, item := range rucksack3.firstCompartmentItems {
		if _, ok := commonItems[item]; ok {
			return item, nil
		}
	}
	for _, item := range rucksack3.secondCompartmentItems {
		if _, ok := commonItems[item]; ok {
			return item, nil
		}
	}

	return item{}, fmt.Errorf("no common item in the three rucksacks")
}
