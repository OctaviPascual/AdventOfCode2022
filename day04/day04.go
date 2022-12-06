package day04

import (
	"fmt"
	"strconv"
	"strings"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	assignments []assignment
}

type assignment struct {
	firstElf, secondElf sectionIDRange
}

type sectionID uint

type sectionIDRange struct {
	start, end sectionID
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	assignmentsString := strings.Split(input, "\n")

	assignments, err := parseAssignments(assignmentsString)
	if err != nil {
		return nil, fmt.Errorf("could not parse assignments: %w", err)
	}

	return &Day{
		assignments: assignments,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	fullyContainedAssignments := getFullyContainedAssignments(d.assignments)

	return fmt.Sprintf("%d", fullyContainedAssignments), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	overlappingAssignments := getOverlappingAssignments(d.assignments)

	return fmt.Sprintf("%d", overlappingAssignments), nil
}

func parseAssignments(assignmentsString []string) ([]assignment, error) {
	assignments := make([]assignment, 0, len(assignmentsString))
	for _, assignmentString := range assignmentsString {
		assignment, err := parseAssignment(assignmentString)
		if err != nil {
			return nil, fmt.Errorf("could not parse assignment: %w", err)
		}

		assignments = append(assignments, assignment)
	}
	return assignments, nil
}

func parseAssignment(assignmentString string) (assignment, error) {
	splitAssignment := strings.Split(assignmentString, ",")
	if len(splitAssignment) != 2 {
		return assignment{}, fmt.Errorf("invalid number of pairs: %d", len(splitAssignment))
	}

	firstElf, err := parseSectionIDRange(splitAssignment[0])
	if err != nil {
		return assignment{}, fmt.Errorf("could not parse first elf assignment (%s): %w", splitAssignment[0], err)
	}

	secondElf, err := parseSectionIDRange(splitAssignment[1])
	if err != nil {
		return assignment{}, fmt.Errorf("could not parse second elf assignment (%s): %w", splitAssignment[1], err)
	}

	return assignment{
		firstElf:  firstElf,
		secondElf: secondElf,
	}, nil
}

func parseSectionIDRange(sectionIDRangeString string) (sectionIDRange, error) {
	splitSectionIDRange := strings.Split(sectionIDRangeString, "-")
	if len(splitSectionIDRange) != 2 {
		return sectionIDRange{}, fmt.Errorf("invalid range format: %s", sectionIDRangeString)
	}

	start, err := strconv.Atoi(splitSectionIDRange[0])
	if err != nil {
		return sectionIDRange{}, fmt.Errorf("could not parse start range: %w", err)
	}

	end, err := strconv.Atoi(splitSectionIDRange[1])
	if err != nil {
		return sectionIDRange{}, fmt.Errorf("could not parse end range: %w", err)
	}

	return sectionIDRange{
		start: sectionID(start),
		end:   sectionID(end),
	}, nil
}

func getFullyContainedAssignments(assignments []assignment) int {
	fullyContainedAssignments := 0
	for _, a := range assignments {
		if a.firstElf.fullyContains(a.secondElf) || a.secondElf.fullyContains(a.firstElf) {
			fullyContainedAssignments++
		}
	}
	return fullyContainedAssignments
}

func getOverlappingAssignments(assignments []assignment) int {
	overlappingAssignments := 0
	for _, a := range assignments {
		if a.firstElf.overlaps(a.secondElf) || a.secondElf.overlaps(a.firstElf) {
			overlappingAssignments++
		}
	}
	return overlappingAssignments
}

func (s sectionIDRange) fullyContains(sectionIDRange sectionIDRange) bool {
	return s.start <= sectionIDRange.start && s.end >= sectionIDRange.end
}

func (s sectionIDRange) overlaps(sectionIDRange sectionIDRange) bool {
	return (sectionIDRange.start >= s.start && sectionIDRange.start <= s.end) ||
		(sectionIDRange.end >= s.start && sectionIDRange.end <= s.end)
}
