package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2022/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	valves map[string]valve
}

var (
	// Regex matching a line of the form "Valve AA has flow rate=0; tunnels lead to valves DD, II, BB"
	valveRe = regexp.MustCompile(`^Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.+)$`)
)

type valve struct {
	label    string
	flowRate int
	tunnels  []string
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	valvesString := strings.Split(input, "\n")

	valves, err := parseValves(valvesString)
	if err != nil {
		return nil, fmt.Errorf("could not parse valves: %w", err)
	}

	return &Day{valves: valves}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	return "", nil
	//opened := util.NewSet[string]()
	//opened := "000000"
	//return fmt.Sprintf("%d", d.rec2(30, d.valves["AA"], opened)), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	return "", nil
}

func parseValves(valvesString []string) (map[string]valve, error) {
	valves := make(map[string]valve, len(valvesString))
	for _, valveString := range valvesString {
		valve, err := parseValve(valveString)
		if err != nil {
			return nil, fmt.Errorf("coud not parse valve: %w", err)
		}
		valves[valve.label] = valve
	}
	return valves, nil
}

func parseValve(valveString string) (valve, error) {
	matches := valveRe.FindStringSubmatch(valveString)
	if len(matches) != 4 {
		return valve{}, fmt.Errorf("invalid valve format: %s", valveString)
	}

	label := matches[1]
	flowRate, _ := strconv.Atoi(matches[2])
	tunnels := strings.Split(matches[3], ", ")

	return valve{
		label:    label,
		flowRate: flowRate,
		tunnels:  tunnels,
	}, nil
}

func (d Day) rec(remainingMinutes int, current valve, opened util.Set[string]) int {
	if len(opened) == 6 {
		return 0
	}
	if remainingMinutes == 0 {
		return 0
	}

	// don't open valve and take tunnel
	tmp2 := 0
	for _, tunnel := range current.tunnels {
		maxPressure := d.rec(remainingMinutes-1, d.valves[tunnel], opened)
		tmp2 = max(tmp2, maxPressure)
	}

	tmp := 0
	// open valve
	if current.flowRate > 0 && !opened.Contains(current.label) {
		opened.Add(current.label)
		for _, tunnel := range current.tunnels {
			maxPressure := d.rec(remainingMinutes-1, d.valves[tunnel], opened)
			tmp = max(tmp, maxPressure)
		}
		opened.Remove(current.label)
		tmp += (remainingMinutes - 1) * current.flowRate
	}

	return max(tmp, tmp2)
}

func isOpened(opened, label string) bool {
	if label == "BB" {
		return opened[0] == '1'
	}
	if label == "CC" {
		return opened[1] == '1'
	}
	if label == "DD" {
		return opened[2] == '1'
	}
	if label == "EE" {
		return opened[3] == '1'
	}
	if label == "HH" {
		return opened[4] == '1'
	}
	if label == "JJ" {
		return opened[5] == '1'
	}
	return false
}

func addToOpened(opened, label string) string {
	if label == "BB" {
		return "1" + opened[1:]
	}
	if label == "CC" {
		return opened[:1] + "1" + opened[2:]
	}
	if label == "DD" {
		return opened[:2] + "1" + opened[3:]
	}
	if label == "EE" {
		return opened[:3] + "1" + opened[4:]
	}
	if label == "HH" {
		return opened[:4] + "1" + opened[5:]
	}
	if label == "JJ" {
		return opened[:5] + "1"
	}
	return ""
}

func (d Day) rec2(remainingMinutes int, current valve, opened string) int {
	if opened == "111111" {
		return 0
	}
	if remainingMinutes == 0 {
		return 0
	}

	// don't open valve and take tunnel
	tmp2 := 0
	for _, tunnel := range current.tunnels {
		maxPressure := d.rec2(remainingMinutes-1, d.valves[tunnel], opened)
		tmp2 = max(tmp2, maxPressure)
	}

	tmp := 0
	// open valve
	if current.flowRate > 0 && !isOpened(opened, current.label) {
		opened = addToOpened(opened, current.label)
		for _, tunnel := range current.tunnels {
			maxPressure := d.rec2(remainingMinutes-1, d.valves[tunnel], opened)
			tmp = max(tmp, maxPressure)
		}
		tmp += (remainingMinutes - 1) * current.flowRate
	}

	return max(tmp, tmp2)
}
