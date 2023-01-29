package day10

import (
	"fmt"
	"strconv"
	"strings"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	program []instruction
}

type instruction interface {
	cycles() int
}

type noopInstruction struct{}

func (i noopInstruction) cycles() int {
	return 1
}

type addXInstruction struct {
	value int
}

func (i addXInstruction) cycles() int {
	return 2
}

type cpu struct {
	registerX int
}

type execution map[int]int

type crt struct {
	sprite int
	pixels [6][40]rune
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	programString := strings.Split(input, "\n")

	program, err := parseProgram(programString)
	if err != nil {
		return nil, fmt.Errorf("could not parse program: %w", err)
	}

	return &Day{
		program: program,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	var cycles = []int{20, 60, 100, 140, 180, 220}

	cpu := newCpu()
	execution := cpu.execute(d.program)

	sumSignalStrengths := 0
	for _, cycle := range cycles {
		sumSignalStrengths += execution.signalStrength(cycle)
	}

	return fmt.Sprintf("%d", sumSignalStrengths), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	cpu := newCpu()
	execution := cpu.execute(d.program)

	crt := newCrt()
	return crt.render(execution), nil
}

func parseProgram(programString []string) ([]instruction, error) {
	program := make([]instruction, 0, len(programString))
	for _, instructionString := range programString {
		motion, err := parseInstruction(instructionString)
		if err != nil {
			return nil, fmt.Errorf("could not parse motion: %w", err)
		}

		program = append(program, motion)
	}
	return program, nil
}

func parseInstruction(instructionString string) (instruction, error) {
	switch {

	case instructionString == "noop":
		return noopInstruction{}, nil

	case strings.HasPrefix(instructionString, "addx"):
		splitAddXInstruction := strings.Split(instructionString, " ")
		if len(splitAddXInstruction) != 2 {
			return nil, fmt.Errorf("invalid format of addx: %s", instructionString)
		}

		value, err := strconv.Atoi(splitAddXInstruction[1])
		if err != nil {
			return nil, fmt.Errorf("invalid value: %w", err)
		}
		return addXInstruction{value: value}, nil
	}
	return nil, fmt.Errorf("unknown instruction: %s", instructionString)
}

func newCpu() *cpu {
	return &cpu{
		registerX: 1,
	}
}

func (c cpu) execute(program []instruction) execution {
	execution := make(execution)
	cycle := 1
	for _, instruction := range program {
		for i := 0; i < instruction.cycles(); i++ {
			execution[cycle] = c.registerX
			cycle++
		}
		if addXInstruction, ok := instruction.(addXInstruction); ok {
			c.registerX += addXInstruction.value
		}
	}
	return execution
}

func (e execution) signalStrength(cycle int) int {
	return e[cycle] * cycle
}

func newCrt() *crt {
	var pixels [6][40]rune
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			pixels[i][j] = '.'
		}
	}
	return &crt{
		pixels: pixels,
	}
}

func (c crt) render(execution execution) string {
	cycle := 1
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			spritePosition := execution[cycle]
			if spritePosition-1 <= j && j <= spritePosition+1 {
				c.pixels[i][j] = '#'
			}
			cycle++
		}
	}

	var s strings.Builder
	s.WriteRune('\n')
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			s.WriteRune(c.pixels[i][j])
		}
		s.WriteRune('\n')
	}
	return s.String()
}
