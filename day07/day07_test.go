package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDay(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	// I tried a lot of time trying to test this: it's tricky as pointers are involved,
	// hence I ended up manually checking that the parsing works.
	_, err := NewDay(input)
	require.NoError(t, err)
}

func TestSolvePartOne(t *testing.T) {
	day := &Day{
		fileSystem: getFileSystem(),
	}

	answer, err := day.SolvePartOne()
	require.NoError(t, err)

	assert.Equal(t, "95437", answer)
}

func TestSolvePartTwo(t *testing.T) {
	day := &Day{
		fileSystem: getFileSystem(),
	}

	answer, err := day.SolvePartTwo()
	require.NoError(t, err)

	assert.Equal(t, "24933642", answer)
}

func getFileSystem() *Node {
	fileI := Node{name: "i", size: 584}
	directoryE := Node{
		name:     "e",
		children: []*Node{&fileI},
	}
	fileF := Node{name: "f", size: 29116}
	fileG := Node{name: "g", size: 2557}
	fileH := Node{name: "h.lst", size: 62596}
	directoryA := Node{
		name:     "a",
		children: []*Node{&directoryE, &fileF, &fileG, &fileH},
	}
	fileB := Node{name: "b.txt", size: 14848514}
	fileC := Node{name: "c.dat", size: 8504156}
	fileJ := Node{name: "j", size: 4060174}
	fileD1 := Node{name: "d.log", size: 8033020}
	fileD2 := Node{name: "d.ext", size: 5626152}
	fileK := Node{name: "K", size: 7214296}
	directoryD := Node{
		name:     "d",
		children: []*Node{&fileJ, &fileD1, &fileD2, &fileK},
	}
	root := Node{
		name:     "/",
		children: []*Node{&directoryA, &fileB, &fileC, &directoryD},
	}

	fileI.parent = &directoryE
	directoryE.parent = &directoryA
	fileF.parent = &directoryA
	fileG.parent = &directoryA
	fileH.parent = &directoryA
	directoryA.parent = &root
	fileB.parent = &root
	fileC.parent = &root
	fileJ.parent = &directoryD
	fileD1.parent = &directoryD
	fileD2.parent = &directoryD
	fileK.parent = &directoryD
	directoryD.parent = &root

	return &root
}
