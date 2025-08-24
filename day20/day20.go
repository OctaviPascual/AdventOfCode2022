package day20

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	encryptedFile []int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	lines := strings.Split(input, "\n")

	encryptedFile, err := parseEncryptedFile(lines)
	if err != nil {
		return nil, fmt.Errorf("could not parse encrypted file: %w", err)
	}

	return &Day{encryptedFile: encryptedFile}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	f := newFile(d.encryptedFile)
	f.mix()
	return fmt.Sprintf("%d", f.getGrooveCoordinates()), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	const decryptionKey = 811589153

	applyDecryptionKey(d.encryptedFile, decryptionKey)

	f := newFile(d.encryptedFile)
	for i := 0; i < 10; i++ {
		f.mix()
	}

	return fmt.Sprintf("%d", f.getGrooveCoordinates()), nil
}

func parseEncryptedFile(lines []string) ([]int, error) {
	numbers := make([]int, 0, len(lines))

	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("could not parse number %s: %w", line, err)
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func applyDecryptionKey(encryptedFile []int, decryptionKey int) {
	for i := range encryptedFile {
		encryptedFile[i] *= decryptionKey
	}
}

type file struct {
	index []*ring.Ring
	ring  *ring.Ring
}

func newFile(encryptedFile []int) file {
	index := make([]*ring.Ring, len(encryptedFile))

	r := ring.New(len(encryptedFile))

	for i, number := range encryptedFile {
		index[i] = r
		r.Value = number
		r = r.Next()
	}

	return file{index: index, ring: r}
}

func (f file) mix() {
	n := len(f.index)

	for i := range f.index {
		positionsToMove := f.index[i].Value.(int)
		previousElement := f.index[i].Prev()
		removedElement := previousElement.Unlink(1)
		previousElement.Move(positionsToMove % (n - 1)).Link(removedElement)
	}
}

func (f file) getGrooveCoordinates() int {
	r := f.ring
	for r.Value.(int) != 0 {
		r = r.Next()
	}
	return r.Move(1_000).Value.(int) + r.Move(2_000).Value.(int) + r.Move(3_000).Value.(int)
}
