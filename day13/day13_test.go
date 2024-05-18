package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	p := func(v int) *int { return &v }
	expected := &Day{
		packetPairs: []packetPair{
			{
				left:  packet{list: []value{{integer: p(1)}, {integer: p(1)}, {integer: p(3)}, {integer: p(1)}, {integer: p(1)}}},
				right: packet{list: []value{{integer: p(1)}, {integer: p(1)}, {integer: p(5)}, {integer: p(1)}, {integer: p(1)}}},
			},
			{
				left:  packet{list: []value{{list: []value{{integer: p(1)}}}, {list: []value{{integer: p(2)}, {integer: p(3)}, {integer: p(4)}}}}},
				right: packet{list: []value{{list: []value{{integer: p(1)}}}, {integer: p(4)}}},
			},
			{
				left:  packet{list: []value{{integer: p(9)}}},
				right: packet{list: []value{{list: []value{{integer: p(8)}, {integer: p(7)}, {integer: p(6)}}}}},
			},
			{
				left:  packet{list: []value{{list: []value{{integer: p(4)}, {integer: p(4)}}}, {integer: p(4)}, {integer: p(4)}}},
				right: packet{list: []value{{list: []value{{integer: p(4)}, {integer: p(4)}}}, {integer: p(4)}, {integer: p(4)}, {integer: p(4)}}},
			},
			{
				left:  packet{list: []value{{integer: p(7)}, {integer: p(7)}, {integer: p(7)}, {integer: p(7)}}},
				right: packet{list: []value{{integer: p(7)}, {integer: p(7)}, {integer: p(7)}}},
			},
			{
				left:  packet{},
				right: packet{list: []value{{integer: p(3)}}},
			},
			{
				left:  packet{list: []value{{list: []value{{}}}}},
				right: packet{list: []value{{}}},
			},
			{
				left:  packet{list: []value{{integer: p(1)}, {list: []value{{integer: p(2)}, {list: []value{{integer: p(3)}, {list: []value{{integer: p(4)}, {list: []value{{integer: p(5)}, {integer: p(6)}, {integer: p(7)}}}}}}}}}, {integer: p(8)}, {integer: p(9)}}},
				right: packet{list: []value{{integer: p(1)}, {list: []value{{integer: p(2)}, {list: []value{{integer: p(3)}, {list: []value{{integer: p(4)}, {list: []value{{integer: p(5)}, {integer: p(6)}, {integer: p(0)}}}}}}}}}, {integer: p(8)}, {integer: p(9)}}},
			},
		},
	}
	input := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	p := func(v int) *int { return &v }
	day := &Day{
		packetPairs: []packetPair{
			{
				left:  packet{list: []value{{integer: p(1)}, {integer: p(1)}, {integer: p(3)}, {integer: p(1)}, {integer: p(1)}}},
				right: packet{list: []value{{integer: p(1)}, {integer: p(1)}, {integer: p(5)}, {integer: p(1)}, {integer: p(1)}}},
			},
			{
				left:  packet{list: []value{{list: []value{{integer: p(1)}}}, {list: []value{{integer: p(2)}, {integer: p(3)}, {integer: p(4)}}}}},
				right: packet{list: []value{{list: []value{{integer: p(1)}}}, {integer: p(4)}}},
			},
			{
				left:  packet{list: []value{{integer: p(9)}}},
				right: packet{list: []value{{list: []value{{integer: p(8)}, {integer: p(7)}, {integer: p(6)}}}}},
			},
			{
				left:  packet{list: []value{{list: []value{{integer: p(4)}, {integer: p(4)}}}, {integer: p(4)}, {integer: p(4)}}},
				right: packet{list: []value{{list: []value{{integer: p(4)}, {integer: p(4)}}}, {integer: p(4)}, {integer: p(4)}, {integer: p(4)}}},
			},
			{
				left:  packet{list: []value{{integer: p(7)}, {integer: p(7)}, {integer: p(7)}, {integer: p(7)}}},
				right: packet{list: []value{{integer: p(7)}, {integer: p(7)}, {integer: p(7)}}},
			},
			{
				left:  packet{},
				right: packet{list: []value{{integer: p(3)}}},
			},
			{
				left:  packet{list: []value{{list: []value{{}}}}},
				right: packet{list: []value{{}}},
			},
			{
				left:  packet{list: []value{{integer: p(1)}, {list: []value{{integer: p(2)}, {list: []value{{integer: p(3)}, {list: []value{{integer: p(4)}, {list: []value{{integer: p(5)}, {integer: p(6)}, {integer: p(7)}}}}}}}}}, {integer: p(8)}, {integer: p(9)}}},
				right: packet{list: []value{{integer: p(1)}, {list: []value{{integer: p(2)}, {list: []value{{integer: p(3)}, {list: []value{{integer: p(4)}, {list: []value{{integer: p(5)}, {integer: p(6)}, {integer: p(0)}}}}}}}}}, {integer: p(8)}, {integer: p(9)}}},
			},
		},
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "13", answer)
}

func TestSolvePartTwo(t *testing.T) {
	p := func(v int) *int { return &v }
	day := &Day{
		packetPairs: []packetPair{
			{
				left:  packet{list: []value{{integer: p(1)}, {integer: p(1)}, {integer: p(3)}, {integer: p(1)}, {integer: p(1)}}},
				right: packet{list: []value{{integer: p(1)}, {integer: p(1)}, {integer: p(5)}, {integer: p(1)}, {integer: p(1)}}},
			},
			{
				left:  packet{list: []value{{list: []value{{integer: p(1)}}}, {list: []value{{integer: p(2)}, {integer: p(3)}, {integer: p(4)}}}}},
				right: packet{list: []value{{list: []value{{integer: p(1)}}}, {integer: p(4)}}},
			},
			{
				left:  packet{list: []value{{integer: p(9)}}},
				right: packet{list: []value{{list: []value{{integer: p(8)}, {integer: p(7)}, {integer: p(6)}}}}},
			},
			{
				left:  packet{list: []value{{list: []value{{integer: p(4)}, {integer: p(4)}}}, {integer: p(4)}, {integer: p(4)}}},
				right: packet{list: []value{{list: []value{{integer: p(4)}, {integer: p(4)}}}, {integer: p(4)}, {integer: p(4)}, {integer: p(4)}}},
			},
			{
				left:  packet{list: []value{{integer: p(7)}, {integer: p(7)}, {integer: p(7)}, {integer: p(7)}}},
				right: packet{list: []value{{integer: p(7)}, {integer: p(7)}, {integer: p(7)}}},
			},
			{
				left:  packet{},
				right: packet{list: []value{{integer: p(3)}}},
			},
			{
				left:  packet{list: []value{{list: []value{{}}}}},
				right: packet{list: []value{{}}},
			},
			{
				left:  packet{list: []value{{integer: p(1)}, {list: []value{{integer: p(2)}, {list: []value{{integer: p(3)}, {list: []value{{integer: p(4)}, {list: []value{{integer: p(5)}, {integer: p(6)}, {integer: p(7)}}}}}}}}}, {integer: p(8)}, {integer: p(9)}}},
				right: packet{list: []value{{integer: p(1)}, {list: []value{{integer: p(2)}, {list: []value{{integer: p(3)}, {list: []value{{integer: p(4)}, {list: []value{{integer: p(5)}, {integer: p(6)}, {integer: p(0)}}}}}}}}}, {integer: p(8)}, {integer: p(9)}}},
			},
		},
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "140", answer)
}

func TestIsPacketPairInRightOrder(t *testing.T) {
	packetPair := packetPair{
		left:  packet{},
		right: packet{},
	}
	assert.False(t, isPacketPairInRightOrder(packetPair))
}
