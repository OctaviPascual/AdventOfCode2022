package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		sensors: []sensor{
			{position: position{x: 2, y: 18}, closestBeacon: position{x: -2, y: 15}},
			{position: position{x: 9, y: 16}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 13, y: 2}, closestBeacon: position{x: 15, y: 3}},
			{position: position{x: 12, y: 14}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 10, y: 20}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 14, y: 17}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 8, y: 7}, closestBeacon: position{x: 2, y: 10}},
			{position: position{x: 2, y: 0}, closestBeacon: position{x: 2, y: 10}},
			{position: position{x: 0, y: 11}, closestBeacon: position{x: 2, y: 10}},
			{position: position{x: 20, y: 14}, closestBeacon: position{x: 25, y: 17}},
			{position: position{x: 17, y: 20}, closestBeacon: position{x: 21, y: 22}},
			{position: position{x: 16, y: 7}, closestBeacon: position{x: 15, y: 3}},
			{position: position{x: 14, y: 3}, closestBeacon: position{x: 15, y: 3}},
			{position: position{x: 20, y: 1}, closestBeacon: position{x: 15, y: 3}},
		},
		maxY: maxY,
	}
	input := `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		sensors: []sensor{
			{position: position{x: 2, y: 18}, closestBeacon: position{x: -2, y: 15}},
			{position: position{x: 9, y: 16}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 13, y: 2}, closestBeacon: position{x: 15, y: 3}},
			{position: position{x: 12, y: 14}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 10, y: 20}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 14, y: 17}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 8, y: 7}, closestBeacon: position{x: 2, y: 10}},
			{position: position{x: 2, y: 0}, closestBeacon: position{x: 2, y: 10}},
			{position: position{x: 0, y: 11}, closestBeacon: position{x: 2, y: 10}},
			{position: position{x: 20, y: 14}, closestBeacon: position{x: 25, y: 17}},
			{position: position{x: 17, y: 20}, closestBeacon: position{x: 21, y: 22}},
			{position: position{x: 16, y: 7}, closestBeacon: position{x: 15, y: 3}},
			{position: position{x: 14, y: 3}, closestBeacon: position{x: 15, y: 3}},
			{position: position{x: 20, y: 1}, closestBeacon: position{x: 15, y: 3}},
		},
		maxY: 20,
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "26", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		sensors: []sensor{
			{position: position{x: 2, y: 18}, closestBeacon: position{x: -2, y: 15}},
			{position: position{x: 9, y: 16}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 13, y: 2}, closestBeacon: position{x: 15, y: 3}},
			{position: position{x: 12, y: 14}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 10, y: 20}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 14, y: 17}, closestBeacon: position{x: 10, y: 16}},
			{position: position{x: 8, y: 7}, closestBeacon: position{x: 2, y: 10}},
			{position: position{x: 2, y: 0}, closestBeacon: position{x: 2, y: 10}},
			{position: position{x: 0, y: 11}, closestBeacon: position{x: 2, y: 10}},
			{position: position{x: 20, y: 14}, closestBeacon: position{x: 25, y: 17}},
			{position: position{x: 17, y: 20}, closestBeacon: position{x: 21, y: 22}},
			{position: position{x: 16, y: 7}, closestBeacon: position{x: 15, y: 3}},
			{position: position{x: 14, y: 3}, closestBeacon: position{x: 15, y: 3}},
			{position: position{x: 20, y: 1}, closestBeacon: position{x: 15, y: 3}},
		},
		maxY: 20,
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "56000011", answer)
}
