package day13

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	packetPairs []packetPair
}

type packetPair struct {
	left  packet
	right packet
}

type packet value

type value struct {
	integer *int
	list    []value
}

type comparisonOutcome string

const (
	rightOrder   comparisonOutcome = "right order"
	wrongOrder   comparisonOutcome = "wrong order"
	keepChecking comparisonOutcome = "keep checking"
)

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	packetPairsString := strings.Split(input, "\n")
	packetPairs, err := parsePacketPairs(packetPairsString)
	if err != nil {
		return nil, fmt.Errorf("could not parse packet pairs: %w", err)
	}

	return &Day{packetPairs: packetPairs}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	sumOfIndices := 0
	for i, packetPair := range d.packetPairs {
		if isPacketPairInRightOrder(packetPair) {
			sumOfIndices += i + 1
		}
	}
	return fmt.Sprintf("%d", sumOfIndices), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	decoderKey := d.getDecoderKey()
	return fmt.Sprintf("%d", decoderKey), nil
}

func parsePacketPairs(packetPairsString []string) ([]packetPair, error) {
	packetPairs := make([]packetPair, 0, (len(packetPairsString)+1)/3)
	for i := 0; i < len(packetPairsString); i += 3 {
		packetPair, err := parsePacketPair(packetPairsString[i], packetPairsString[i+1])
		if err != nil {
			return nil, fmt.Errorf("could not parse packet pair: %w", err)
		}
		packetPairs = append(packetPairs, packetPair)
	}
	return packetPairs, nil
}

func parsePacketPair(leftString, rightString string) (packetPair, error) {
	left, err := parsePacket(leftString)
	if err != nil {
		return packetPair{}, fmt.Errorf("could not parse left packet: %w", err)
	}

	right, err := parsePacket(rightString)
	if err != nil {
		return packetPair{}, fmt.Errorf("could not parse right packet: %w", err)
	}

	return packetPair{left: left, right: right}, nil
}

func parsePacket(packetString string) (packet, error) {
	value, err := parseValue(packetString)
	return packet(value), err
}

func parseValue(valueString string) (value, error) {
	if isInt(valueString) {
		integer, _ := strconv.Atoi(valueString)
		return value{integer: &integer}, nil
	}

	n := len(valueString)
	if valueString[0] != '[' {
		return value{}, fmt.Errorf("invalid value format, must start with [")
	}
	if valueString[n-1] != ']' {
		return value{}, fmt.Errorf("invalid value format, must end with ]")
	}

	if valueString[1:n-1] == "" {
		return value{}, nil
	}

	valuesString := splitByOuterCommas(valueString[1 : n-1])
	var list []value
	for _, v := range valuesString {
		parsedValue, err := parseValue(v)
		if err != nil {
			return value{}, fmt.Errorf("could not parse value: %w", err)
		}

		list = append(list, parsedValue)
	}
	return value{list: list}, nil
}

func splitByOuterCommas(s string) []string {
	var result []string
	start := 0
	commas := outerCommasPositions(s)
	for _, comma := range commas {
		result = append(result, s[start:comma])
		start = comma + 1
	}
	result = append(result, s[start:])
	return result
}

func outerCommasPositions(s string) []int {
	var outerCommas []int
	count := 0
	for i, c := range s {
		if c == '[' {
			count++
		}
		if c == ']' {
			count--
		}
		if count == 0 && c == ',' {
			outerCommas = append(outerCommas, i)
		}
	}
	return outerCommas
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return s != ""
}

func isPacketPairInRightOrder(packetPair packetPair) bool {
	outcome := compare(value(packetPair.left), value(packetPair.right))
	return outcome == rightOrder
}

func compare(leftValue, rightValue value) comparisonOutcome {
	if leftValue.integer != nil && rightValue.integer != nil {
		if *leftValue.integer == *rightValue.integer {
			return keepChecking
		}
		if *leftValue.integer < *rightValue.integer {
			return rightOrder
		}
		return wrongOrder
	}

	if leftValue.integer == nil && rightValue.integer == nil {
		i := 0
		for i = 0; i < len(leftValue.list) && i < len(rightValue.list); i++ {
			outcome := compare(leftValue.list[i], rightValue.list[i])
			if outcome != keepChecking {
				return outcome
			}
		}
		if len(leftValue.list) == len(rightValue.list) {
			return keepChecking
		}
		if i == len(leftValue.list) {
			return rightOrder
		}
		return wrongOrder
	}

	if leftValue.integer != nil {
		return compare(value{list: []value{{integer: leftValue.integer}}}, rightValue)
	}

	if rightValue.integer != nil {
		return compare(leftValue, value{list: []value{{integer: rightValue.integer}}})
	}

	return keepChecking
}

func (d Day) getDecoderKey() int {
	int2 := 2
	int6 := 6
	dividerPacket2 := packet{list: []value{{integer: &int2}}}
	dividerPacket6 := packet{list: []value{{integer: &int6}}}

	packets := make([]packet, 0, len(d.packetPairs)*2+2)
	packets = append(packets, dividerPacket2, dividerPacket6)

	for _, packetPair := range d.packetPairs {
		packets = append(packets, packetPair.left, packetPair.right)
	}

	slices.SortFunc(packets, func(a, b packet) int {
		if isPacketPairInRightOrder(packetPair{left: a, right: b}) {
			return -1
		}
		return 1
	})

	var dividerPacket2Index, dividerPacket6Index int
	for i, packet := range packets {
		if len(packet.list) > 0 && packet.list[0].integer == &int2 {
			dividerPacket2Index = i + 1
		}
		if len(packet.list) > 0 && packet.list[0].integer == &int6 {
			dividerPacket6Index = i + 1
		}
	}
	return dividerPacket2Index * dividerPacket6Index
}
