package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	expected := &Day{
		monkeys: map[string]monkey{
			"root": {
				name: "root",
				operation: &operation{
					left:     "pppw",
					operator: "+",
					right:    "sjmn",
				},
			},
			"dbpl": {
				name:   "dbpl",
				number: 5,
			},
			"cczh": {
				name: "cczh",
				operation: &operation{
					left:     "sllz",
					operator: "+",
					right:    "lgvd",
				},
			},
			"zczc": {
				name:   "zczc",
				number: 2,
			},
			"ptdq": {
				name: "ptdq",
				operation: &operation{
					left:     "humn",
					operator: "-",
					right:    "dvpt",
				},
			},
			"dvpt": {
				name:   "dvpt",
				number: 3,
			},
			"lfqf": {
				name:   "lfqf",
				number: 4,
			},
			"humn": {
				name:   "humn",
				number: 5,
			},
			"ljgn": {
				name:   "ljgn",
				number: 2,
			},
			"sjmn": {
				name: "sjmn",
				operation: &operation{
					left:     "drzm",
					operator: "*",
					right:    "dbpl",
				},
			},
			"sllz": {
				name:   "sllz",
				number: 4,
			},
			"pppw": {
				name: "pppw",
				operation: &operation{
					left:     "cczh",
					operator: "/",
					right:    "lfqf",
				},
			},
			"lgvd": {
				name: "lgvd",
				operation: &operation{
					left:     "ljgn",
					operator: "*",
					right:    "ptdq",
				},
			},
			"drzm": {
				name: "drzm",
				operation: &operation{
					left:     "hmdt",
					operator: "-",
					right:    "zczc",
				},
			},
			"hmdt": {
				name:   "hmdt",
				number: 32,
			},
		},
	}
	input := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`
	actual, err := NewDay(input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestSolvePartOne(t *testing.T) {
	input := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "152", answer)
}

func TestSolvePartTwo(t *testing.T) {
	input := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`
	day, err := NewDay(input)
	require.NoError(t, err)

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "301", answer)
}
