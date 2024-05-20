package day15

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2022/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	sensors []sensor
	maxY    int
}

const maxY = 4_000_000

var (
	// Regex matching a line of the form "Sensor at x=2, y=18: closest beacon is at x=-2, y=15"
	sensorRe = regexp.MustCompile(`^Sensor at x=(|-?\d+), y=(|-?\d+): closest beacon is at x=(|-?\d+), y=(|-?\d+)$`)

	// emptyInterval represents an empty interval
	emptyInterval = interval{}
)

type sensor struct {
	position      position
	closestBeacon position
}

type position struct {
	x, y int
}

type interval struct {
	start int
	end   int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	sensorsString := strings.Split(input, "\n")
	sensors, err := parseSensors(sensorsString)
	if err != nil {
		return nil, fmt.Errorf("could not parse sensors: %w", err)
	}

	return &Day{sensors: sensors, maxY: maxY}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	intervals := d.noBeaconIntervals(d.maxY / 2)
	noBeaconCount := d.getNoBeaconCount(intervals[0], d.maxY/2)
	return fmt.Sprintf("%d", noBeaconCount), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	distressBeacon, err := d.getDistressBeacon()
	if err != nil {
		return "", fmt.Errorf("could not get distress beacon: %w", err)
	}

	return fmt.Sprintf("%d", distressBeacon.tuningFrequency()), nil
}

func parseSensors(sensorsString []string) ([]sensor, error) {
	sensors := make([]sensor, 0, len(sensorsString))
	for _, sensorString := range sensorsString {
		sensor, err := parseSensor(sensorString)
		if err != nil {
			return nil, fmt.Errorf("could not parse sensor: %w", err)
		}
		sensors = append(sensors, sensor)
	}
	return sensors, nil
}

func parseSensor(sensorString string) (sensor, error) {
	matches := sensorRe.FindStringSubmatch(sensorString)
	if len(matches) != 5 {
		return sensor{}, fmt.Errorf("invalid sensor format: %s", sensorString)
	}

	position, err := parsePosition(matches[1], matches[2])
	if err != nil {
		return sensor{}, fmt.Errorf("coud not parse sensor position: %w", err)
	}

	closestBeacon, err := parsePosition(matches[3], matches[4])
	if err != nil {
		return sensor{}, fmt.Errorf("coud not parse closest beacon position: %w", err)
	}

	return sensor{position: position, closestBeacon: closestBeacon}, nil
}

func parsePosition(xString, yString string) (position, error) {
	x, err := parseSignedInt(xString)
	if err != nil {
		return position{}, fmt.Errorf("coud not parse x: %w", err)
	}

	y, err := parseSignedInt(yString)
	if err != nil {
		return position{}, fmt.Errorf("coud not parse y: %w", err)
	}

	return position{x: x, y: y}, nil
}

func parseSignedInt(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("empty integer")
	}

	if s[0] == '-' {
		i, err := strconv.Atoi(s[1:])
		if err != nil {
			return 0, err
		}
		return -i, nil
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (d Day) getDistressBeacon() (position, error) {
	for y := 0; y <= d.maxY; y++ {
		intervals := d.noBeaconIntervals(y)
		if len(intervals) == 2 {
			if intervals[0].end+2 != intervals[1].start {
				return position{}, fmt.Errorf("more than one position for distress beacon")
			}
			return position{x: intervals[0].end + 1, y: y}, nil
		}
	}
	return position{}, fmt.Errorf("no possible position found for distress beacon")
}

func (d Day) noBeaconIntervals(yRow int) []interval {
	intervals := make([]interval, 0, len(d.sensors))
	for _, sensor := range d.sensors {
		interval := noBeaconPositions(sensor, yRow)
		if interval != emptyInterval {
			intervals = append(intervals, interval)
		}
	}

	slices.SortFunc(intervals, func(a, b interval) int {
		return cmp.Compare(a.start, b.start)
	})

	union := []interval{intervals[0]}
	for _, interval := range intervals[1:] {
		current := &union[len(union)-1]
		if current.end+1 == interval.start {
			current.end = interval.end
			continue
		}
		if current.end < interval.start {
			union = append(union, interval)
			continue
		}
		if current.end >= interval.start {
			current.end = max(interval.end, current.end)
		}
	}
	return union
}

func (d Day) getNoBeaconCount(interval interval, y int) int {
	noBeaconCount := abs(interval.end-interval.start) + 1
	visitedBeacons := util.NewSet[position]()
	for _, sensor := range d.sensors {
		if sensor.closestBeacon.y != y || visitedBeacons.Contains(sensor.closestBeacon) {
			continue
		}
		if interval.start <= sensor.closestBeacon.x && sensor.closestBeacon.x <= interval.end {
			noBeaconCount--
			visitedBeacons.Add(sensor.closestBeacon)
		}
	}
	return noBeaconCount
}

// noBeaconPositions returns the interval for which there can't be beacons according to a sensor
func noBeaconPositions(sensor sensor, y int) interval {
	d := sensor.position.manhattanDistance(sensor.closestBeacon)
	start := sensor.position.x - (d - abs(y-sensor.position.y))
	end := sensor.position.x + (d - abs(y-sensor.position.y))
	if start > end {
		return emptyInterval
	}
	return interval{start: start, end: end}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (p position) tuningFrequency() int {
	return p.x*maxY + p.y
}

func (p position) manhattanDistance(position position) int {
	minX := min(p.x, position.x)
	maxX := max(p.x, position.x)
	minY := min(p.y, position.y)
	maxY := max(p.y, position.y)
	return maxX - minX + maxY - minY
}