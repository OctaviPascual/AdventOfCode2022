package main

import (
	"bufio"
	"fmt"
	"os"
)

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

type blueprint struct {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	state := state{
		blueprint: blueprint{
			oreRobotCost:      robotCost{ore: 2},
			clayRobotCost:     robotCost{ore: 3},
			obsidianRobotCost: robotCost{ore: 3, clay: 8},
			geodeRobotCost:    robotCost{ore: 3, obsidian: 12},
		},
		oreRobots: 1,
	}
	for i := 1; i <= 24; i++ {
		fmt.Printf("\n== Minute %d ==\n", i)
		robot := buildRobot(reader, &state)
		collectResources(&state)
		addRobot(&state, robot)
	}
	fmt.Printf("You ended up with %d geodes\n", state.geodes)
}

type robot rune

const (
	oreRobot      robot = 'o'
	clayRobot     robot = 'c'
	obsidianRobot robot = 'b'
	geodeRobot    robot = 'g'
	noRobot       robot = 'n'
)

func buildRobot(reader *bufio.Reader, state *state) robot {
	for {
		fmt.Println("Which robot do you want to build?")
		char, _, _ := reader.ReadRune()
		// discard CR
		_, _, _ = reader.ReadRune()

		switch char {
		case rune(oreRobot):
			if state.ore >= state.blueprint.oreRobotCost.ore {
				state.ore -= state.blueprint.oreRobotCost.ore
				return oreRobot
			}
		case rune(clayRobot):
			if state.ore >= state.blueprint.clayRobotCost.ore {
				state.ore -= state.blueprint.clayRobotCost.ore
				return clayRobot
			}
		case rune(obsidianRobot):
			if state.ore >= state.blueprint.obsidianRobotCost.ore && state.clay >= state.blueprint.obsidianRobotCost.clay {
				state.ore -= state.blueprint.obsidianRobotCost.ore
				state.clay -= state.blueprint.obsidianRobotCost.clay
				return obsidianRobot
			}
		case rune(geodeRobot):
			if state.ore >= state.blueprint.geodeRobotCost.ore && state.obsidian >= state.blueprint.geodeRobotCost.obsidian {
				state.ore -= state.blueprint.geodeRobotCost.ore
				state.obsidian -= state.blueprint.geodeRobotCost.obsidian
				return geodeRobot
			}
		case rune(noRobot):
			return noRobot
		default:
			fmt.Printf("This robot is unknown: %c\n", char)
		}
		fmt.Println("Not enough resources to build the robot, choose another action")
	}
}

func collectResources(state *state) {
	state.ore += state.oreRobots
	fmt.Printf("%d ore-collecting robots collect %d ore; you now have %d ore.\n", state.oreRobots, state.oreRobots, state.ore)

	state.clay += state.clayRobots
	fmt.Printf("%d clay-collecting robots collect %d clay; you now have %d clay.\n", state.clayRobots, state.clayRobots, state.clay)

	state.obsidian += state.obsidianRobots
	fmt.Printf("%d obsidian-collecting robots collect %d obsidian; you now have %d obsidian.\n", state.obsidianRobots, state.obsidianRobots, state.obsidian)

	state.geodes += state.geodeRobots
	fmt.Printf("%d geode-cracking robots collect %d geodes; you now have %d geodes.\n", state.geodeRobots, state.geodeRobots, state.geodes)
}

func addRobot(state *state, robot robot) {
	switch robot {
	case oreRobot:
		state.oreRobots++
		fmt.Printf("The new ore-robot is ready; you now have %d of them.\n", state.oreRobots)
	case clayRobot:
		state.clayRobots++
		fmt.Printf("The new clay-robot is ready; you now have %d of them.\n", state.clayRobots)
	case obsidianRobot:
		state.obsidianRobots++
		fmt.Printf("The new obsidian-robot is ready; you now have %d of them.\n", state.obsidianRobots)
	case geodeRobot:
		state.geodeRobots++
		fmt.Printf("The new geode-robot is ready; you now have %d of them.\n", state.geodeRobots)
	}
}
