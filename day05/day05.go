package day05

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/OctaviPascual/AdventOfCode2022/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	stacks                 map[stackID]*stack
	rearrangementProcedure []step
}

var (
	// Regex matching a step of the form "move 13 from 7 to 2"
	stepRe = regexp.MustCompile(`^move (\d+) from (\d) to (\d)$`)
)

type stackID int
type crate rune

type stack struct {
	id     stackID
	crates []crate
}

type step struct {
	crates int
	from   stackID
	to     stackID
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	splitInput := strings.Split(input, "\n")

	endStackSection := slices.Index(splitInput, "")
	if endStackSection == -1 {
		return nil, fmt.Errorf("invalid input: no separation between stack section and rearrangement procedure")
	}

	stacks, err := parseStacks(splitInput[:endStackSection])
	if err != nil {
		return nil, fmt.Errorf("could not parse stacks: %w", err)
	}

	rearrangementProcedure, err := parseRearrangementProcedure(splitInput[endStackSection+1:])
	if err != nil {
		return nil, fmt.Errorf("could not parse rearrangement procedure: %w", err)
	}

	return &Day{
		stacks:                 stacks,
		rearrangementProcedure: rearrangementProcedure,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	stacks, err := runRearrangementProcedureWithCrateMover9000(d.rearrangementProcedure, d.stacks)
	if err != nil {
		return "", fmt.Errorf("error running rearrangement procedure: %w", err)
	}

	cratesOnTopOfEachStack, err := getCratesOnTopOfEachStack(stacks)
	if err != nil {
		return "", fmt.Errorf("could not get crates on top of each stack: %w", err)
	}

	return cratesOnTopOfEachStack, nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	stacks, err := runRearrangementProcedureWithCrateMover9001(d.rearrangementProcedure, d.stacks)
	if err != nil {
		return "", fmt.Errorf("error running rearrangement procedure: %w", err)
	}

	cratesOnTopOfEachStack, err := getCratesOnTopOfEachStack(stacks)
	if err != nil {
		return "", fmt.Errorf("could not get crates on top of each stack: %w", err)
	}

	return cratesOnTopOfEachStack, nil
}

// We assume that the number of stacks is less than 10.
// In the input, each stack is made of 3 characters and there is 1 space between each stack.
// Let N be the total number of characters in a line, n the number of stacks:
// Then, N = 3n + (n-1) = 4n - 1 <=> N + 1 = 4n
// Doing modular arithmetic we have: (N + 1) % 4 = 0 <=> N % 4 = 3
func parseStacks(stacksString []string) (map[stackID]*stack, error) {
	if len(stacksString[0])%4 != 3 {
		return nil, fmt.Errorf("invalid length of stacks string")
	}

	numberOfStacks := (len(stacksString[0]) + 1) / 4

	stacks := make(map[stackID]*stack, numberOfStacks)
	for i := 1; i <= numberOfStacks; i++ {
		stacks[stackID(i)] = &stack{id: stackID(i)}
	}

	for _, stackLine := range stacksString {
		parseStackLine(stackLine, stacks)
	}

	for _, v := range stacks {
		util.Reverse(v.crates)
	}

	return stacks, nil
}

func parseStackLine(stackLine string, stacks map[stackID]*stack) {
	for i := 0; i < len(stacks); i++ {
		if stackLine[4*i] == ' ' {
			continue
		}
		if stackLine[4*i] == '[' {
			stacks[stackID(i+1)].crates = append(stacks[stackID(i+1)].crates, crate(stackLine[4*i+1]))
			continue
		}
	}
}

func parseRearrangementProcedure(rearrangementProcedureString []string) ([]step, error) {
	steps := make([]step, 0, len(rearrangementProcedureString))

	for _, stepString := range rearrangementProcedureString {
		step, err := parseStep(stepString)
		if err != nil {
			return nil, fmt.Errorf("could not parse step: %w", err)
		}

		steps = append(steps, step)
	}
	return steps, nil
}

func parseStep(stepString string) (step, error) {
	matches := stepRe.FindStringSubmatch(stepString)
	if len(matches) != 4 {
		return step{}, fmt.Errorf("invalid format of step string: %s", stepString)
	}

	// We know that the step has a valid format so no need to check for errors
	crates, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])

	return step{
		crates: crates,
		from:   stackID(from),
		to:     stackID(to),
	}, nil
}

func cloneStacks(stacks map[stackID]*stack) map[stackID]*stack {
	clonedStacks := make(map[stackID]*stack, len(stacks))
	for id, stackV := range stacks {
		crates := make([]crate, len(stackV.crates))
		copy(crates, stackV.crates)
		clonedStacks[id] = &stack{
			id:     id,
			crates: crates,
		}
	}
	return clonedStacks
}

func runRearrangementProcedureWithCrateMover9000(rearrangementProcedure []step, stacks map[stackID]*stack) (map[stackID]*stack, error) {
	clonedStacks := cloneStacks(stacks)

	for _, step := range rearrangementProcedure {
		for i := 0; i < step.crates; i++ {
			err := moveCrates(step.from, step.to, clonedStacks, 1)
			if err != nil {
				return nil, fmt.Errorf("error moving crate from stack[%d] to stack[%d]", step.from, step.to)
			}
		}
	}
	return clonedStacks, nil
}

func runRearrangementProcedureWithCrateMover9001(rearrangementProcedure []step, stacks map[stackID]*stack) (map[stackID]*stack, error) {
	clonedStacks := cloneStacks(stacks)

	for _, step := range rearrangementProcedure {
		err := moveCrates(step.from, step.to, clonedStacks, step.crates)
		if err != nil {
			return nil, fmt.Errorf("could not move %d crates from stack[%d] to stack[%d]: %w",
				step.crates, step.from, step.to, err,
			)
		}
	}
	return clonedStacks, nil
}

func moveCrates(from, to stackID, stacks map[stackID]*stack, amount int) error {
	fromStack, ok := stacks[from]
	if !ok {
		return fmt.Errorf("invalid from stack ID: %d", from)
	}

	toStack, ok := stacks[to]
	if !ok {
		return fmt.Errorf("invalid to stack ID: %d", from)
	}

	n := len(fromStack.crates) - amount
	if n < 0 {
		return fmt.Errorf("stack[%d] has only %d crate(s)", fromStack.id, len(fromStack.crates))
	}

	topCrates := fromStack.crates[n:]
	fromStack.crates = fromStack.crates[:n]
	toStack.crates = append(toStack.crates, topCrates...)
	return nil
}

func getCratesOnTopOfEachStack(stacks map[stackID]*stack) (string, error) {
	cratesOnTopOfEachStack := make([]crate, 0, len(stacks))
	for id := 1; id <= len(stacks); id++ {
		stack := stacks[stackID(id)]
		n := len(stack.crates) - 1

		if n < 0 {
			return "", fmt.Errorf("stack[%d] is empty", stack.id)
		}

		cratesOnTopOfEachStack = append(cratesOnTopOfEachStack, stack.crates[n])
	}

	return string(cratesOnTopOfEachStack), nil
}
