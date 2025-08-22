package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		blueprints: []blueprint{
			{
				ID:                1,
				oreRobotCost:      robotCost{ore: 4},
				clayRobotCost:     robotCost{ore: 2},
				obsidianRobotCost: robotCost{ore: 3, clay: 14},
				geodeRobotCost:    robotCost{ore: 2, obsidian: 7},
			},
			{
				ID:                2,
				oreRobotCost:      robotCost{ore: 2},
				clayRobotCost:     robotCost{ore: 3},
				obsidianRobotCost: robotCost{ore: 3, clay: 8},
				geodeRobotCost:    robotCost{ore: 3, obsidian: 12},
			},
		},
	}
	input := `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		blueprints: []blueprint{
			{
				ID:                1,
				oreRobotCost:      robotCost{ore: 4},
				clayRobotCost:     robotCost{ore: 2},
				obsidianRobotCost: robotCost{ore: 3, clay: 14},
				geodeRobotCost:    robotCost{ore: 2, obsidian: 7},
			},
			{
				ID:                2,
				oreRobotCost:      robotCost{ore: 2},
				clayRobotCost:     robotCost{ore: 3},
				obsidianRobotCost: robotCost{ore: 3, clay: 8},
				geodeRobotCost:    robotCost{ore: 3, obsidian: 12},
			},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "33", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "", answer)
}
