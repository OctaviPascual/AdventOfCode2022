package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		valves: map[string]valve{
			"AA": {label: "AA", flowRate: 0, tunnels: []string{"DD", "II", "BB"}},
			"BB": {label: "BB", flowRate: 13, tunnels: []string{"CC", "AA"}},
			"CC": {label: "CC", flowRate: 2, tunnels: []string{"DD", "BB"}},
			"DD": {label: "DD", flowRate: 20, tunnels: []string{"CC", "AA", "EE"}},
			"EE": {label: "EE", flowRate: 3, tunnels: []string{"FF", "DD"}},
			"FF": {label: "FF", flowRate: 0, tunnels: []string{"EE", "GG"}},
			"GG": {label: "GG", flowRate: 0, tunnels: []string{"FF", "HH"}},
			"HH": {label: "HH", flowRate: 22, tunnels: []string{"GG"}},
			"II": {label: "II", flowRate: 0, tunnels: []string{"AA", "JJ"}},
			"JJ": {label: "JJ", flowRate: 21, tunnels: []string{"II"}},
		},
	}
	input := `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		valves: map[string]valve{
			"AA": {label: "AA", flowRate: 0, tunnels: []string{"DD", "II", "BB"}},
			"BB": {label: "BB", flowRate: 13, tunnels: []string{"CC", "AA"}},
			"CC": {label: "CC", flowRate: 2, tunnels: []string{"DD", "BB"}},
			"DD": {label: "DD", flowRate: 20, tunnels: []string{"CC", "AA", "EE"}},
			"EE": {label: "EE", flowRate: 3, tunnels: []string{"FF", "DD"}},
			"FF": {label: "FF", flowRate: 0, tunnels: []string{"EE", "GG"}},
			"GG": {label: "GG", flowRate: 0, tunnels: []string{"FF", "HH"}},
			"HH": {label: "HH", flowRate: 22, tunnels: []string{"GG"}},
			"II": {label: "II", flowRate: 0, tunnels: []string{"AA", "JJ"}},
			"JJ": {label: "JJ", flowRate: 21, tunnels: []string{"II"}},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "", answer)
}
