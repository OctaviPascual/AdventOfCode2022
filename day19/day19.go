package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	blueprints []blueprint
}

var (
	// Regex matching a blueprint such as "Blueprint 12: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 7 clay. Each geode robot costs 4 ore and 20 obsidian."
	blueprintRe = regexp.MustCompile(`^Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.$`)
)

type blueprint struct {
	ID                int
	oreRobotCost      robotCost
	clayRobotCost     robotCost
	obsidianRobotCost robotCost
	geodeRobotCost    robotCost
}

type robotCost struct {
	ore      int
	clay     int
	obsidian int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	blueprintsString := strings.Split(input, "\n")

	blueprints, err := parseBlueprints(blueprintsString)
	if err != nil {
		return nil, fmt.Errorf("could not parse cubes: %w", err)
	}

	return &Day{blueprints: blueprints}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	qualityLevelsSum := 0
	for _, blueprint := range d.blueprints {
		state := state{
			blueprint: blueprint,
			oreRobots: 1,
		}
		maxGeodes := 0
		rec(state, 1, &maxGeodes)
		qualityLevelsSum += blueprint.qualityLevel(maxGeodes)
	}

	return fmt.Sprintf("%d", qualityLevelsSum), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	return "", nil
}

func parseBlueprints(blueprintsString []string) ([]blueprint, error) {
	blueprints := make([]blueprint, 0, len(blueprintsString))
	for _, blueprintString := range blueprintsString {
		blueprint, err := parseBlueprint(blueprintString)
		if err != nil {
			return nil, fmt.Errorf("coud not parse blueprint: %w", err)
		}
		blueprints = append(blueprints, blueprint)
	}
	return blueprints, nil
}

func parseBlueprint(blueprintString string) (blueprint, error) {
	matches := blueprintRe.FindStringSubmatch(blueprintString)
	if len(matches) != 8 {
		return blueprint{}, fmt.Errorf("invalid blueprint format: %s", blueprintString)
	}

	id, _ := strconv.Atoi(matches[1])
	oreRobotOreCost, _ := strconv.Atoi(matches[2])
	clayRobotOreCost, _ := strconv.Atoi(matches[3])
	obsidianRobotOreCost, _ := strconv.Atoi(matches[4])
	obsidianRobotClayCost, _ := strconv.Atoi(matches[5])
	geodeRobotOreCost, _ := strconv.Atoi(matches[6])
	geodeRobotObsidianCost, _ := strconv.Atoi(matches[7])

	return blueprint{
		ID:                id,
		oreRobotCost:      robotCost{ore: oreRobotOreCost},
		clayRobotCost:     robotCost{ore: clayRobotOreCost},
		obsidianRobotCost: robotCost{ore: obsidianRobotOreCost, clay: obsidianRobotClayCost},
		geodeRobotCost:    robotCost{ore: geodeRobotOreCost, obsidian: geodeRobotObsidianCost},
	}, nil
}

type state struct {
	blueprint blueprint

	ore      int
	clay     int
	obsidian int

	oreRobots      int
	clayRobots     int
	obsidianRobots int
	geodeRobots    int

	geodes int
}

type action string

const (
	totalMinutes = 24

	oreRobot      action = "oreRobot"
	clayRobot     action = "clayRobot"
	obsidianRobot action = "obsidianRobot"
	geodeRobot    action = "geodeRobot"
	noRobot       action = "noRobot"
)

func (b blueprint) qualityLevel(geodes int) int {
	return b.ID * geodes
}

func rec(state state, minute int, maxGeodes *int) {
	if minute > totalMinutes {
		if state.geodes > *maxGeodes {
			*maxGeodes = state.geodes
		}
		return
	}

	if maxBoundOfGeodes(minute, state.geodes, state.geodeRobots) <= *maxGeodes {
		return
	}

	actions := getActions(state, minute)

	state2 := collectResources(state)

	for _, action := range actions {
		state3 := executeAction(state2, action)
		rec(state3, minute+1, maxGeodes)
	}
}

func getActions(state state, minute int) []action {
	if minute == totalMinutes {
		// Last action doesn't matter so let's always do this one.
		return []action{noRobot}
	}

	possibleActions := []action{noRobot}

	// If we build more ore robots than maxOre we will be wasting resources on each minute.
	maxOre := max(state.blueprint.oreRobotCost.ore, state.blueprint.clayRobotCost.ore, state.blueprint.obsidianRobotCost.ore, state.blueprint.geodeRobotCost.ore)
	if state.ore >= state.blueprint.oreRobotCost.ore && state.oreRobots < maxOre {
		possibleActions = append(possibleActions, oreRobot)
	}

	// If we build more clay robots than maxClay we will be wasting resources on each minute.
	maxClay := state.blueprint.obsidianRobotCost.clay
	if state.ore >= state.blueprint.clayRobotCost.ore && state.clayRobots < maxClay {
		possibleActions = append(possibleActions, clayRobot)
	}

	if state.ore >= state.blueprint.obsidianRobotCost.ore && state.clay >= state.blueprint.obsidianRobotCost.clay {
		possibleActions = append(possibleActions, obsidianRobot)
	}

	if state.ore >= state.blueprint.geodeRobotCost.ore && state.obsidian >= state.blueprint.geodeRobotCost.obsidian {
		// Building a geode robot will always provide more geodes than performing any other action.
		return []action{geodeRobot}
	}
	return possibleActions
}

func collectResources(s state) state {
	return state{
		blueprint: s.blueprint,

		ore:      s.ore + s.oreRobots,
		clay:     s.clay + s.clayRobots,
		obsidian: s.obsidian + s.obsidianRobots,
		geodes:   s.geodes + s.geodeRobots,

		oreRobots:      s.oreRobots,
		clayRobots:     s.clayRobots,
		obsidianRobots: s.obsidianRobots,
		geodeRobots:    s.geodeRobots,
	}
}

func executeAction(s state, action action) state {
	newState := state{
		blueprint: s.blueprint,

		ore:      s.ore,
		clay:     s.clay,
		obsidian: s.obsidian,
		geodes:   s.geodes,

		oreRobots:      s.oreRobots,
		clayRobots:     s.clayRobots,
		obsidianRobots: s.obsidianRobots,
		geodeRobots:    s.geodeRobots,
	}
	switch action {
	case oreRobot:
		newState.ore -= s.blueprint.oreRobotCost.ore
		newState.oreRobots++
	case clayRobot:
		newState.ore -= s.blueprint.clayRobotCost.ore
		newState.clayRobots++
	case obsidianRobot:
		newState.ore -= s.blueprint.obsidianRobotCost.ore
		newState.clay -= s.blueprint.obsidianRobotCost.clay
		newState.obsidianRobots++
	case geodeRobot:
		newState.ore -= s.blueprint.geodeRobotCost.ore
		newState.obsidian -= s.blueprint.geodeRobotCost.obsidian
		newState.geodeRobots++
	}
	return newState
}

func maxBoundOfGeodes(minute, geodes, geodeRobots int) int {
	maxBound := geodes
	for i := minute; i <= totalMinutes; i++ {
		maxBound += geodeRobots
		geodeRobots++
	}
	return maxBound
}
