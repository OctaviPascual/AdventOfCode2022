package day11

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	monkeys []monkey
}

var (
	// Regex matching a line of the form "  Starting items: 79, 98"
	startingItemsRe = regexp.MustCompile(`^  Starting items: (.*)$`)
	// Regex matching a line of the form "  Operation: new = old * 19"
	operationRe = regexp.MustCompile(`^  Operation: new = old (.) (.*)$`)
	// Regex matching a line of the form "  Test: divisible by 23"
	testRe = regexp.MustCompile(`^  Test: divisible by (\d+)$`)
	// Regex matching a line of the form "    If true: throw to monkey 2"
	ifTrueRe = regexp.MustCompile(`^    If true: throw to monkey (\d+)$`)
	// Regex matching a line of the form "    If false: throw to monkey 3"
	ifFalseRe = regexp.MustCompile(`^    If false: throw to monkey (\d+)$`)
)

type monkey struct {
	startingItems []item
	operation     operation
	test          test
}

type item struct {
	worryLevel int
}

type operation struct {
	operator string
	operand  int
}

type test struct {
	divisibleBy   int
	monkeyIfTrue  int
	monkeyIfFalse int
}

type monkeyState struct {
	itemsHolding        [][]int
	totalItemsInspected []int
}

type reliefWorryLevelFn func(int) int

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	lines := strings.Split(input, "\n")

	monkeys, err := parseMonkeys(lines)
	if err != nil {
		return nil, fmt.Errorf("could not parse monkeys: %w", err)
	}

	return &Day{monkeys: monkeys}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	monkeyState := newMonkeyState(d.monkeys)
	reliefWorryLevelFn := func(worryLevel int) int { return worryLevel / 3 }

	for i := 0; i < 20; i++ {
		monkeyState = d.playRound(monkeyState, reliefWorryLevelFn)
	}
	return fmt.Sprintf("%d", monkeyBusiness(monkeyState)), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	monkeyState := newMonkeyState(d.monkeys)
	totalModulo := d.computeTotalModulo()
	reliefWorryLevelFn := func(worryLevel int) int { return worryLevel % totalModulo }

	for i := 0; i < 10_000; i++ {
		monkeyState = d.playRound(monkeyState, reliefWorryLevelFn)
	}
	return fmt.Sprintf("%d", monkeyBusiness(monkeyState)), nil
}

func parseMonkeys(lines []string) ([]monkey, error) {
	var monkeys []monkey
	for i := 0; i < len(lines); i += 7 {
		monkey, err := parseMonkey(lines[i : i+6])
		if err != nil {
			return nil, fmt.Errorf("could not parse monkey: %w", err)
		}

		monkeys = append(monkeys, monkey)
	}

	return monkeys, nil
}

func parseMonkey(lines []string) (monkey, error) {
	matches := startingItemsRe.FindStringSubmatch(lines[1])
	if len(matches) != 2 {
		return monkey{}, fmt.Errorf("invalid starting items list: %s", lines[1])
	}

	itemsString := strings.Split(matches[1], ",")
	startingItems, err := parseStartingItems(itemsString)
	if err != nil {
		return monkey{}, fmt.Errorf("could not parse items: %w", err)
	}

	operation, err := parseOperation(lines[2])
	if err != nil {
		return monkey{}, fmt.Errorf("could not parse operation: %w", err)
	}

	test, err := parseTest(lines[3], lines[4], lines[5])
	if err != nil {
		return monkey{}, fmt.Errorf("could not parse test: %w", err)
	}

	return monkey{startingItems: startingItems, operation: operation, test: test}, nil
}

func parseStartingItems(itemsString []string) ([]item, error) {
	var items []item
	for _, itemString := range itemsString {
		item, err := parseItem(strings.TrimSpace(itemString))
		if err != nil {
			return nil, fmt.Errorf("could not parse item: %w", err)
		}

		items = append(items, item)
	}
	return items, nil
}

func parseItem(itemString string) (item, error) {
	worryLevel, err := strconv.Atoi(itemString)
	if err != nil {
		return item{}, fmt.Errorf("invalid worry level value: %w", err)
	}

	return item{worryLevel: worryLevel}, nil
}

func parseOperation(operationString string) (operation, error) {
	matches := operationRe.FindStringSubmatch(operationString)
	if len(matches) != 3 {
		return operation{}, fmt.Errorf("invalid operation format: %s", operationString)
	}

	operator := matches[1]

	if matches[2] == "old" {
		return operation{operator: "* old"}, nil
	}

	operand, err := strconv.Atoi(matches[2])
	if err != nil {
		return operation{}, fmt.Errorf("could not parse operand: %w", err)
	}
	return operation{operator: operator, operand: operand}, nil
}

func parseTest(testString, monkeyIfTrueString, monkeyIfFalseString string) (test, error) {
	matches := testRe.FindStringSubmatch(testString)
	if len(matches) != 2 {
		return test{}, fmt.Errorf("invalid test format: %s", testString)
	}

	divisibleBy, err := strconv.Atoi(matches[1])
	if err != nil {
		return test{}, fmt.Errorf("could not parse divisible by: %w", err)
	}

	matches = ifTrueRe.FindStringSubmatch(monkeyIfTrueString)
	if len(matches) != 2 {
		return test{}, fmt.Errorf("invalid if true format: %s", testString)
	}

	monkeyIfTrue, err := strconv.Atoi(matches[1])
	if err != nil {
		return test{}, fmt.Errorf("could not parse if true monkey: %w", err)
	}

	matches = ifFalseRe.FindStringSubmatch(monkeyIfFalseString)
	if len(matches) != 2 {
		return test{}, fmt.Errorf("invalid if false format: %s", testString)
	}

	monkeyIfFalse, err := strconv.Atoi(matches[1])
	if err != nil {
		return test{}, fmt.Errorf("could not parse if false monkey: %w", err)
	}

	return test{
		divisibleBy:   divisibleBy,
		monkeyIfTrue:  monkeyIfTrue,
		monkeyIfFalse: monkeyIfFalse,
	}, nil
}

func newMonkeyState(monkeys []monkey) *monkeyState {
	itemsHolding := make([][]int, len(monkeys))
	for i, monkey := range monkeys {
		itemsHolding[i] = make([]int, len(monkey.startingItems))
		for j, item := range monkey.startingItems {
			itemsHolding[i][j] = item.worryLevel
		}
	}

	totalItemsInspected := make([]int, len(monkeys))
	return &monkeyState{itemsHolding: itemsHolding, totalItemsInspected: totalItemsInspected}
}

func (d Day) playRound(monkeyState *monkeyState, reliefFn reliefWorryLevelFn) *monkeyState {
	for i := 0; i < len(d.monkeys); i++ {
		monkeyState = d.executeTurn(i, monkeyState, reliefFn)
	}
	return monkeyState
}

func (d Day) executeTurn(
	currentMonkey int,
	state *monkeyState,
	reliefFn reliefWorryLevelFn,
) *monkeyState {
	for _, currentItem := range state.itemsHolding[currentMonkey] {
		worryLevel := d.monkeys[currentMonkey].operation.apply(currentItem)
		worryLevel = reliefFn(worryLevel)

		recipient := d.monkeys[currentMonkey].test.throwItemTo(worryLevel)

		state.itemsHolding[recipient] = append(state.itemsHolding[recipient], worryLevel)

		state.totalItemsInspected[currentMonkey]++
	}
	state.itemsHolding[currentMonkey] = nil
	return state
}

func (d Day) computeTotalModulo() int {
	totalModulo := 1
	for i := 0; i < len(d.monkeys); i++ {
		totalModulo *= d.monkeys[i].test.divisibleBy
	}
	return totalModulo
}

func monkeyBusiness(monkeyState *monkeyState) int {
	n := len(monkeyState.totalItemsInspected)
	slices.Sort(monkeyState.totalItemsInspected)
	return monkeyState.totalItemsInspected[n-1] * monkeyState.totalItemsInspected[n-2]
}

func (o operation) apply(worryLevel int) int {
	if o.operator == "* old" {
		return worryLevel * worryLevel
	}

	if o.operator == "*" {
		return worryLevel * o.operand
	}

	return worryLevel + o.operand
}

func (t test) throwItemTo(worryLevel int) int {
	if worryLevel%t.divisibleBy == 0 {
		return t.monkeyIfTrue
	}
	return t.monkeyIfFalse
}
