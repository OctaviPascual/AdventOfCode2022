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

const (
	totalMinutesPartOne = 24
	totalMinutesPartTwo = 32
)

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
		maxGeodes := 0
		explore(newState(blueprint), totalMinutesPartOne, &maxGeodes)
		qualityLevelsSum += blueprint.qualityLevel(maxGeodes)
	}
	return fmt.Sprintf("%d", qualityLevelsSum), nil

}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	maxGeodesBlueprint1 := 0
	explore(newState(d.blueprints[0]), totalMinutesPartTwo, &maxGeodesBlueprint1)

	maxGeodesBlueprint2 := 0
	explore(newState(d.blueprints[1]), totalMinutesPartTwo, &maxGeodesBlueprint2)

	maxGeodesBlueprint3 := 0
	explore(newState(d.blueprints[2]), totalMinutesPartTwo, &maxGeodesBlueprint3)

	return fmt.Sprintf("%d", maxGeodesBlueprint1*maxGeodesBlueprint2*maxGeodesBlueprint3), nil
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

	minute int

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
	oreRobot      action = "oreRobot"
	clayRobot     action = "clayRobot"
	obsidianRobot action = "obsidianRobot"
	geodeRobot    action = "geodeRobot"
	noRobot       action = "noRobot"
)

func newState(blueprint blueprint) state {
	return state{
		blueprint: blueprint,
		minute:    1,
		oreRobots: 1,
	}
}

func (b blueprint) qualityLevel(geodes int) int {
	return b.ID * geodes
}

func explore(state state, totalMinutes int, maxGeodes *int) {
	if state.minute > totalMinutes {
		if state.geodes > *maxGeodes {
			*maxGeodes = state.geodes
		}
		return
	}

	if maxBoundOfGeodes(state, totalMinutes) <= *maxGeodes {
		return
	}

	actions := getActions(state, totalMinutes)

	state2 := collectResources(state)

	for _, action := range actions {
		state3 := executeAction(state2, action)
		explore(state3, totalMinutes, maxGeodes)
	}
}

func getActions(s state, totalMinutes int) []action {
	if s.minute == totalMinutes {
		// Last action doesn't matter so let's always do this one.
		return []action{noRobot}
	}

	actions := []action{noRobot}

	// If we build more ore robots than maxOre we will be wasting resources on each minute.
	maxOre := max(s.blueprint.oreRobotCost.ore, s.blueprint.clayRobotCost.ore, s.blueprint.obsidianRobotCost.ore, s.blueprint.geodeRobotCost.ore)
	if s.ore >= s.blueprint.oreRobotCost.ore && s.oreRobots < maxOre {
		actions = append(actions, oreRobot)
	}

	// If we build more clay robots than maxClay we will be wasting resources on each minute.
	maxClay := s.blueprint.obsidianRobotCost.clay
	if s.ore >= s.blueprint.clayRobotCost.ore && s.clayRobots < maxClay {
		actions = append(actions, clayRobot)
	}

	// If we build more obsidian robots than maxObsidian we will be wasting resources on each minute.
	maxObsidian := s.blueprint.geodeRobotCost.obsidian
	if s.ore >= s.blueprint.obsidianRobotCost.ore && s.clay >= s.blueprint.obsidianRobotCost.clay && s.obsidianRobots < maxObsidian {
		actions = append(actions, obsidianRobot)
	}

	if s.ore >= s.blueprint.geodeRobotCost.ore && s.obsidian >= s.blueprint.geodeRobotCost.obsidian {
		// Building a geode robot will always provide more geodes than performing any other action.
		return []action{geodeRobot}
	}

	// If it's possible to build any robot, it's better to build one than none.
	if len(actions) == 4 {
		return []action{oreRobot, clayRobot, obsidianRobot}
	}

	return actions
}

func collectResources(s state) state {
	return state{
		blueprint: s.blueprint,

		minute: s.minute,

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

		minute: s.minute + 1,

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

func maxBoundOfGeodes(s state, totalMinutes int) int {
	// If we have no obsidian robots we are minutes away from building geode robots.
	// Note that this assumes that geode robots cost obsidian.
	if s.obsidianRobots == 0 {
		maxBound, geodeRobots := 0, 0
		// We won't be able to start building a geode robot at least until 5 minutes later.
		// Note that this value of 5 is empirical from the input we have, with lower obsidian costs it could be less.
		for i := s.minute + 5; i <= totalMinutes; i++ {
			maxBound += geodeRobots
			geodeRobots++
		}
		return maxBound
	}

	maxBound := s.geodes
	geodeRobots := s.geodeRobots
	canBuildGeodeRobotThisMinute := s.ore >= s.blueprint.geodeRobotCost.ore && s.obsidian >= s.blueprint.geodeRobotCost.obsidian

	for i := s.minute; i <= totalMinutes; i++ {
		maxBound += geodeRobots
		if i > s.minute || canBuildGeodeRobotThisMinute {
			geodeRobots++
		}
	}
	return maxBound
}
