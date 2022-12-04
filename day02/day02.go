package day02

import (
	"fmt"
	"strings"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	strategyGuide strategyGuide
}

type strategyGuide []round

type round struct {
	column1 string
	column2 string
}

type decryptRound func(r round) (shape, shape, error)

type shape string

const (
	paper    shape = "Paper"
	scissors shape = "Scissors"
	rock     shape = "Rock"
)

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	strategyGuideString := strings.Split(input, "\n")

	strategyGuide, err := parseStrategyGuide(strategyGuideString)
	if err != nil {
		return nil, fmt.Errorf("could not parse strategy guide: %w", err)
	}

	return &Day{strategyGuide: strategyGuide}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	score, err := d.strategyGuide.totalScore(decryptRoundPart1)
	if err != nil {
		return "", fmt.Errorf("could not compute total score: %w", err)
	}

	return fmt.Sprintf("%d", score), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	score, err := d.strategyGuide.totalScore(decryptRoundPart2)
	if err != nil {
		return "", fmt.Errorf("could not compute total score: %w", err)
	}

	return fmt.Sprintf("%d", score), nil
}

func parseStrategyGuide(strategyGuideString []string) (strategyGuide, error) {
	strategyGuide := make([]round, 0, len(strategyGuideString))
	for _, roundString := range strategyGuideString {
		round, err := parseRound(roundString)
		if err != nil {
			return nil, fmt.Errorf("could not parse round: %w", err)
		}

		strategyGuide = append(strategyGuide, round)
	}
	return strategyGuide, nil
}

func parseRound(roundString string) (round, error) {
	columns := strings.Split(roundString, " ")
	if len(columns) != 2 {
		return round{}, fmt.Errorf("invalid number of columns: %d", len(columns))
	}

	return round{
		column1: columns[0],
		column2: columns[1],
	}, nil
}

func (s strategyGuide) totalScore(decryptRound decryptRound) (int, error) {
	totalScore := 0
	for _, round := range s {
		_, score, err := round.score(decryptRound)
		if err != nil {
			return 0, fmt.Errorf("could not compute round score: %w", err)
		}
		totalScore += score
	}
	return totalScore, nil
}

func (r round) score(decryptRound decryptRound) (int, int, error) {
	var player1Score, player2Score int
	shape1, shape2, err := decryptRound(r)
	if err != nil {
		return 0, 0, fmt.Errorf("could not decrypt round: %w", err)
	}

	switch {
	case shape1 == shape2:
		player1Score, player2Score = 3, 3
	case shape1 == paper && shape2 == rock:
		player1Score, player2Score = 6, 0
	case shape1 == scissors && shape2 == paper:
		player1Score, player2Score = 6, 0
	case shape1 == rock && shape2 == scissors:
		player1Score, player2Score = 6, 0
	default:
		player1Score, player2Score = 0, 6
	}
	return player1Score + shape1.score(), player2Score + shape2.score(), nil
}

func (s shape) score() int {
	switch s {
	case paper:
		return 2
	case scissors:
		return 3
	case rock:
		return 1
	}
	return 0
}

func decryptRoundPart1(r round) (shape, shape, error) {
	decryptShape := func(shapeString string) (shape, error) {
		switch shapeString {
		case "A", "X":
			return rock, nil
		case "B", "Y":
			return paper, nil
		case "C", "Z":
			return scissors, nil
		}
		return "", fmt.Errorf("invalid shape: %s", shapeString)
	}

	shape1, err := decryptShape(r.column1)
	if err != nil {
		return "", "", fmt.Errorf("could not decrypt first shape: %w", err)
	}

	shape2, err := decryptShape(r.column2)
	if err != nil {
		return "", "", fmt.Errorf("could not decrypt second shape: %w", err)
	}

	return shape1, shape2, nil
}

func decryptRoundPart2(r round) (shape, shape, error) {
	switch {
	case r.column1 == "A" && r.column2 == "X":
		return rock, scissors, nil
	case r.column1 == "A" && r.column2 == "Y":
		return rock, rock, nil
	case r.column1 == "A" && r.column2 == "Z":
		return rock, paper, nil

	case r.column1 == "B" && r.column2 == "X":
		return paper, rock, nil
	case r.column1 == "B" && r.column2 == "Y":
		return paper, paper, nil
	case r.column1 == "B" && r.column2 == "Z":
		return paper, scissors, nil

	case r.column1 == "C" && r.column2 == "X":
		return scissors, paper, nil
	case r.column1 == "C" && r.column2 == "Y":
		return scissors, scissors, nil
	case r.column1 == "C" && r.column2 == "Z":
		return scissors, rock, nil
	}
	return "", "", fmt.Errorf("invalid shapes: %s %s", r.column1, r.column2)
}
