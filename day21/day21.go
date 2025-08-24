package day21

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	monkeys map[string]monkey
}

const (
	rootMonkeyName = "root"
	humanName      = "humn"
)

var (
	// Regex matching an operation as "cczh + lfqf"
	operationRe = regexp.MustCompile(`(\w+) (.) (\w+)`)
)

type monkey struct {
	name      string
	number    int
	operation *operation
}

type operation struct {
	left     string
	operator string
	right    string
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	monkeysString := strings.Split(input, "\n")

	monkeys, err := parseMonkeys(monkeysString)
	if err != nil {
		return nil, fmt.Errorf("could not parse monkeys: %w", err)
	}
	return &Day{monkeys: monkeys}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {

	number, err := yell(rootMonkeyName, d.monkeys)
	if err != nil {
		return "", fmt.Errorf("failed to find number yelled by root monkey: %w", err)
	}

	return fmt.Sprintf("%d", number), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	delete(d.monkeys, humanName)

	name := d.monkeys[rootMonkeyName].operation.right
	target, err := yell(d.monkeys[rootMonkeyName].operation.left, d.monkeys)
	if err != nil {
		name = d.monkeys[rootMonkeyName].operation.left
		target, _ = yell(d.monkeys[rootMonkeyName].operation.right, d.monkeys)
	}

	numberYelledByHuman, err := yellWithTarget(name, d.monkeys, target)
	if err != nil {
		return "", fmt.Errorf("failed to find number yelled by human: %w", err)
	}

	return fmt.Sprintf("%d", numberYelledByHuman), nil
}

func parseMonkeys(monkeysString []string) (map[string]monkey, error) {
	monkeys := make(map[string]monkey, len(monkeysString))
	for _, monkeyString := range monkeysString {
		m, err := parseMonkey(monkeyString)
		if err != nil {
			return nil, fmt.Errorf("coud not parse monkey: %w", err)
		}
		monkeys[m.name] = m
	}
	return monkeys, nil
}

func parseMonkey(monkeyString string) (monkey, error) {
	parts := strings.Split(monkeyString, ":")
	if len(parts) != 2 {
		return monkey{}, fmt.Errorf("invalid line format: %s", monkeyString)
	}

	name := parts[0]
	job := parts[1][1:]

	if unicode.IsDigit(rune(job[0])) {
		number, err := strconv.Atoi(job)
		if err != nil {
			return monkey{}, fmt.Errorf("invalid number format: %w", err)
		}
		return monkey{name: name, number: number}, nil
	}

	matches := operationRe.FindStringSubmatch(job)
	if len(matches) != 4 {
		return monkey{}, fmt.Errorf("invalid operation format: %s", job)
	}
	operation := operation{
		left:     matches[1],
		operator: matches[2],
		right:    matches[3],
	}
	return monkey{
		name:      name,
		operation: &operation,
	}, nil
}

func yell(name string, monkeys map[string]monkey) (int, error) {
	m, ok := monkeys[name]
	if !ok {
		return 0, fmt.Errorf("could not find monkey with name %s", name)
	}

	if m.operation == nil {
		return m.number, nil
	}

	left, err := yell(m.operation.left, monkeys)
	if err != nil {
		return 0, err
	}

	right, err := yell(m.operation.right, monkeys)
	if err != nil {
		return 0, err
	}

	switch m.operation.operator {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		return left / right, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", m.operation.operator)
	}
}

func yellWithTarget(name string, monkeys map[string]monkey, target int) (int, error) {
	if name == humanName {
		return target, nil
	}

	m, ok := monkeys[name]
	if !ok {
		return 0, fmt.Errorf("could not find monkey with name %s", name)
	}

	number, err := yell(m.operation.left, monkeys)
	if err != nil {
		// We know this won't fail because the human is on the left path.
		number, _ = yell(m.operation.right, monkeys)
		switch m.operation.operator {
		case "+":
			return yellWithTarget(m.operation.left, monkeys, target-number)
		case "-":
			return yellWithTarget(m.operation.left, monkeys, target+number)
		case "*":
			return yellWithTarget(m.operation.left, monkeys, target/number)
		case "/":
			return yellWithTarget(m.operation.left, monkeys, target*number)
		default:
			return 0, fmt.Errorf("invalid operator: %s", m.operation.operator)
		}
	}

	switch m.operation.operator {
	case "+":
		return yellWithTarget(m.operation.right, monkeys, target-number)
	case "-":
		return yellWithTarget(m.operation.right, monkeys, number-target)
	case "*":
		return yellWithTarget(m.operation.right, monkeys, target/number)
	case "/":
		return yellWithTarget(m.operation.right, monkeys, number/target)
	default:
		return 0, fmt.Errorf("invalid operator: %s", m.operation.operator)
	}
}
