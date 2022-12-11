package day07

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	totalDiskSpace    = 70_000_000
	unusedSpaceNeeded = 30_000_000
)

var (
	// Regex matching a file in terminal output
	fileRe = regexp.MustCompile(`(\d+) (.+)`)
)

// Day holds the data needed to solve part one and part two
type Day struct {
	fileSystem *Node
}

type Node struct {
	name     string
	size     int
	parent   *Node
	children []*Node
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	terminalOutput := strings.Split(input, "\n")

	root, err := parseTerminalOutput(terminalOutput)
	if err != nil {
		return nil, fmt.Errorf("could not parse terminal output: %w", err)
	}

	return &Day{
		fileSystem: root,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	sizeByDirectory := make(map[string]int)
	computeSizeByDirectory(d.fileSystem, "", sizeByDirectory)

	total := 0
	for _, v := range sizeByDirectory {
		if v <= 100000 {
			total += v
		}
	}

	return fmt.Sprintf("%d", total), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	sizeByDirectory := make(map[string]int)
	computeSizeByDirectory(d.fileSystem, "", sizeByDirectory)

	keys := make([]string, 0, len(sizeByDirectory))
	for key := range sizeByDirectory {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return sizeByDirectory[keys[i]] < sizeByDirectory[keys[j]]
	})

	totalUsedSpace := sizeByDirectory["/"]
	unusedSpace := totalDiskSpace - totalUsedSpace
	if unusedSpace >= unusedSpaceNeeded {
		return "", fmt.Errorf("no directory needs to be deleted to run the update")
	}

	for _, directory := range keys {
		if unusedSpace+sizeByDirectory[directory] >= unusedSpaceNeeded {
			return fmt.Sprintf("%d", sizeByDirectory[directory]), nil
		}
	}

	return "", fmt.Errorf("can't run the update as there is not enough space")
}

func parseTerminalOutput(terminalOutput []string) (*Node, error) {
	commandStart := 1
	commandEnd := 1

	root := Node{
		name: "/",
	}

	var currentNode *Node
	currentNode = &root
	for commandEnd < len(terminalOutput) {
		commandEnd++
		for commandEnd < len(terminalOutput) && !isCommand(terminalOutput[commandEnd]) {
			commandEnd++
		}
		var err error
		currentNode, err = parseCommand(terminalOutput[commandStart:commandEnd], currentNode)
		if err != nil {
			return nil, fmt.Errorf("could not parse command: %w", err)
		}
		commandStart = commandEnd
	}
	return &root, nil
}

func isCommand(line string) bool {
	if len(line) == 0 {
		return false
	}
	return line[0] == '$'
}

func parseCommand(lines []string, currentNode *Node) (*Node, error) {
	if lines[0] == "$ ls" {
		return parseListCommand(lines, currentNode)
	}
	return parseChangeDirectoryCommand(lines[0], currentNode)
}

func parseListCommand(lines []string, currentNode *Node) (*Node, error) {
	for _, line := range lines[1:] {
		var node Node
		if strings.HasPrefix(line, "dir") {
			node = Node{
				name:   line[4:],
				parent: currentNode,
			}
		} else {
			matches := fileRe.FindStringSubmatch(line)
			if len(matches) != 3 {
				return nil, fmt.Errorf("invalid list file output: %s", line)
			}

			size, _ := strconv.Atoi(matches[1])
			name := matches[2]
			node = Node{
				name:   name,
				size:   size,
				parent: currentNode,
			}
		}
		currentNode.children = append(currentNode.children, &node)
	}
	return currentNode, nil
}

func parseChangeDirectoryCommand(line string, currentNode *Node) (*Node, error) {
	if line == "$ cd .." {
		return currentNode.parent, nil
	}
	name := line[5:]
	for _, child := range currentNode.children {
		if child.name == name {
			return child, nil
		}
	}
	return nil, fmt.Errorf("could not find %s directory in current directory %s", name, currentNode.name)
}

func computeSizeByDirectory(n *Node, path string, sizeByDirectory map[string]int) int {
	absolutePath := path + n.name
	if size, ok := sizeByDirectory[absolutePath]; ok {
		return size
	}

	total := 0
	for _, child := range n.children {
		if !child.isDirectory() {
			total += child.size
		}
		total += computeSizeByDirectory(child, absolutePath, sizeByDirectory)
	}

	sizeByDirectory[absolutePath] = total
	return total
}

func (n Node) isDirectory() bool {
	return n.size == 0
}
