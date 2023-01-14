package day08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2022/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	grid [][]tree
}

type tree struct {
	height int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	gridString := strings.Split(input, "\n")

	grid, err := parseGrid(gridString)
	if err != nil {
		return nil, fmt.Errorf("could not parse grid: %w", err)
	}

	return &Day{
		grid: grid,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	return fmt.Sprintf("%d", visibleTrees(d.grid)), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	return fmt.Sprintf("%d", getHighestScenicScore(d.grid)), nil
}

func parseGrid(gridString []string) ([][]tree, error) {
	grid := make([][]tree, 0, len(gridString))
	for _, trees := range gridString {
		row := make([]tree, 0, len(trees))
		for _, treeRune := range trees {
			tree, err := parseTree(treeRune)
			if err != nil {
				return nil, fmt.Errorf("could not parse tree: %w", err)
			}

			row = append(row, tree)
		}
		grid = append(grid, row)
	}

	return grid, nil
}

func parseTree(treeRune rune) (tree, error) {
	height, err := strconv.Atoi(string(treeRune))
	if err != nil {
		return tree{}, fmt.Errorf("could not parse height: %w", err)
	}

	return tree{height: height}, nil
}

func visibleTrees(grid [][]tree) int {
	visibleTrees := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if isVisible(grid, i, j) {
				visibleTrees += 1
			}
		}
	}
	return visibleTrees
}

func isVisible(grid [][]tree, i, j int) bool {
	return isVisibleFromLeft(grid, i, j) ||
		isVisibleFromRight(grid, i, j) ||
		isVisibleFromTop(grid, i, j) ||
		isVisibleFromBottom(grid, i, j)
}

func isVisibleFromLeft(grid [][]tree, i, j int) bool {
	tree := grid[i][j]
	for j > 0 && grid[i][j-1].isShorter(tree) {
		j--
	}
	return j == 0
}

func isVisibleFromRight(grid [][]tree, i, j int) bool {
	tree := grid[i][j]
	for j < len(grid[0])-1 && grid[i][j+1].isShorter(tree) {
		j++
	}
	return j == len(grid[0])-1
}

func isVisibleFromTop(grid [][]tree, i, j int) bool {
	tree := grid[i][j]
	for i > 0 && grid[i-1][j].isShorter(tree) {
		i--
	}
	return i == 0
}

func isVisibleFromBottom(grid [][]tree, i, j int) bool {
	tree := grid[i][j]
	for i < len(grid)-1 && grid[i+1][j].isShorter(tree) {
		i++
	}
	return i == len(grid)-1
}

func getHighestScenicScore(grid [][]tree) int {
	highestScenicScore := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			highestScenicScore = util.Max(highestScenicScore, scenicScore(grid, i, j))
		}
	}
	return highestScenicScore
}

func scenicScore(grid [][]tree, i, j int) int {
	return leftViewingDistance(grid, i, j) *
		rightViewingDistance(grid, i, j) *
		upViewingDistance(grid, i, j) *
		downViewingDistance(grid, i, j)
}

func leftViewingDistance(grid [][]tree, i, j int) int {
	if j == 0 {
		return 0
	}

	tree := grid[i][j]
	viewingDistance := 1
	for j > 1 && grid[i][j-1].isShorter(tree) {
		j--
		viewingDistance++
	}
	return viewingDistance
}

func rightViewingDistance(grid [][]tree, i, j int) int {
	if j == len(grid[0])-1 {
		return 0
	}

	tree := grid[i][j]
	viewingDistance := 1
	for j < len(grid[0])-2 && grid[i][j+1].isShorter(tree) {
		j++
		viewingDistance++
	}
	return viewingDistance
}

func upViewingDistance(grid [][]tree, i, j int) int {
	if i == 0 {
		return 0
	}

	tree := grid[i][j]
	viewingDistance := 1
	for i > 1 && grid[i-1][j].isShorter(tree) {
		i--
		viewingDistance++
	}
	return viewingDistance
}

func downViewingDistance(grid [][]tree, i, j int) int {
	if i == len(grid)-1 {
		return 0
	}

	tree := grid[i][j]
	viewingDistance := 1
	for i < len(grid)-2 && grid[i+1][j].isShorter(tree) {
		i++
		viewingDistance++
	}
	return viewingDistance
}

func (t tree) isShorter(tree tree) bool {
	return t.height < tree.height
}
