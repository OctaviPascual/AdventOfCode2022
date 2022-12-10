package day06

import (
	"fmt"
)

const (
	startOfPacketMarkerSize  = 4
	startOfMessageMarkerSize = 14
)

// Day holds the data needed to solve part one and part two
type Day struct {
	datastreamBuffer datastreamBuffer
}

type datastreamBuffer string

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	return &Day{
		datastreamBuffer: datastreamBuffer(input),
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	firstStartOfPacketMarker := d.datastreamBuffer.getFirstStartOfPacketMarker()
	if firstStartOfPacketMarker == -1 {
		return "", fmt.Errorf("could not find start of packet marker")
	}

	return fmt.Sprintf("%d", firstStartOfPacketMarker), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	firstStartOfMessageMarker := d.datastreamBuffer.getFirstStartOfMessageMarker()
	if firstStartOfMessageMarker == -1 {
		return "", fmt.Errorf("could not find start of message marker")
	}

	return fmt.Sprintf("%d", firstStartOfMessageMarker), nil
}

func (d datastreamBuffer) getFirstStartOfPacketMarker() int {
	return d.getFirstStartOfMarker(startOfPacketMarkerSize)
}

func (d datastreamBuffer) getFirstStartOfMessageMarker() int {
	return d.getFirstStartOfMarker(startOfMessageMarkerSize)
}

func (d datastreamBuffer) getFirstStartOfMarker(size int) int {
	for i := 0; i+size < len(d); i++ {
		if isStartOfPacketMarker([]rune(d[i : i+size])...) {
			return i + size
		}
	}
	return -1
}

func isStartOfPacketMarker(characters ...rune) bool {
	charactersSet := make(map[rune]struct{})
	for _, character := range characters {
		charactersSet[character] = struct{}{}
	}
	return len(charactersSet) == len(characters)
}
