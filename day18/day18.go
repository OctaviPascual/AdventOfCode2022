package day18

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/OctaviPascual/AdventOfCode2022/util"
)

// Day holds the data needed to solve part one and part two
type Day struct {
	cubes []cube
}

var (
	// Regex matching a 3D position such as "1,22,333"
	positionRe = regexp.MustCompile(`^(\d+),(\d+),(\d+)$`)
)

type cube struct {
	x, y, z int
}

type lavaDroplet struct {
	surface int
	cubes   util.Set[cube]
}

type boundingBox struct {
	xMin, yMin, xMax, yMax, zMin, zMax int
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	cubesString := strings.Split(input, "\n")

	cubes, err := parseCubes(cubesString)
	if err != nil {
		return nil, fmt.Errorf("could not parse cubes: %w", err)
	}

	return &Day{
		cubes: cubes,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	lavaDroplet := newLavaDroplet(d.cubes)
	return fmt.Sprintf("%d", lavaDroplet.getSurface()), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	lavaDroplet := newLavaDroplet(d.cubes)
	return fmt.Sprintf("%d", lavaDroplet.getExteriorSurface()), nil
}

func parseCubes(cubesString []string) ([]cube, error) {
	cubes := make([]cube, 0, len(cubesString))
	for _, cubeString := range cubesString {
		cube, err := parseCube(cubeString)
		if err != nil {
			return nil, fmt.Errorf("coud not parse cube: %w", err)
		}
		cubes = append(cubes, cube)
	}
	return cubes, nil
}

func parseCube(cubeString string) (cube, error) {
	matches := positionRe.FindStringSubmatch(cubeString)
	if len(matches) != 4 {
		return cube{}, fmt.Errorf("invalid cube format: %s", cubeString)
	}

	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	z, _ := strconv.Atoi(matches[3])

	return cube{x: x, y: y, z: z}, nil
}

func (c cube) adjacentCubes() []cube {
	return []cube{
		{c.x - 1, c.y, c.z},
		{c.x + 1, c.y, c.z},
		{c.x, c.y - 1, c.z},
		{c.x, c.y + 1, c.z},
		{c.x, c.y, c.z - 1},
		{c.x, c.y, c.z + 1},
	}
}

func (c cube) isInside(boundingBox boundingBox) bool {
	return c.x >= boundingBox.xMin && c.x <= boundingBox.xMax &&
		c.y >= boundingBox.yMin && c.y <= boundingBox.yMax &&
		c.z >= boundingBox.zMin && c.z <= boundingBox.zMax
}

func newLavaDroplet(cubes []cube) lavaDroplet {
	lavaDroplet := lavaDroplet{cubes: util.NewSet[cube]()}
	for _, cube := range cubes {
		lavaDroplet.addCube(cube)
	}
	return lavaDroplet
}

func (l *lavaDroplet) addCube(cube cube) {
	// When two cubes are adjacent two sides of the droplet disappear.
	l.surface += 6 - 2*l.numberOfAdjacentCubes(cube)
	l.cubes.Add(cube)
}

func (l *lavaDroplet) numberOfAdjacentCubes(cube cube) int {
	adjacentCubes := 0
	for _, c := range cube.adjacentCubes() {
		if l.cubes.Contains(c) {
			adjacentCubes++
		}
	}
	return adjacentCubes
}

func (l *lavaDroplet) getSurface() int {
	return l.surface
}

func (l *lavaDroplet) getExteriorSurface() int {
	boundingBox := getBoundingBox(l.cubes.Members())

	currentCube := cube{boundingBox.xMin, boundingBox.yMin, boundingBox.zMin}
	steam, visited := util.NewSet[cube](), util.NewSet[cube]()
	coverWithSteam(currentCube, l, steam, visited, boundingBox)

	exteriorSurface := 0
	for _, c := range steam.Members() {
		exteriorSurface += l.numberOfAdjacentCubes(c)
	}
	return exteriorSurface
}

func getBoundingBox(cubes []cube) boundingBox {
	xMin, xMax, yMin, yMax, zMin, zMax := cubes[0].x, cubes[0].x, cubes[0].y, cubes[0].y, cubes[0].z, cubes[0].z
	for _, cube := range cubes {
		xMin = min(xMin, cube.x)
		xMax = max(xMax, cube.x)

		yMin = min(yMin, cube.y)
		yMax = max(yMax, cube.y)

		zMin = min(zMin, cube.z)
		zMax = max(zMax, cube.z)
	}
	return boundingBox{
		xMin: xMin - 1,
		xMax: xMax + 1,
		yMin: yMin - 1,
		yMax: yMax + 1,
		zMin: zMin - 1,
		zMax: zMax + 1,
	}
}

func coverWithSteam(currentCube cube, lavaDroplet *lavaDroplet, steam, visited util.Set[cube], boundingBox boundingBox) {
	if !currentCube.isInside(boundingBox) || lavaDroplet.cubes.Contains(currentCube) || visited.Contains(currentCube) {
		return
	}
	visited.Add(currentCube)

	for _, c := range currentCube.adjacentCubes() {
		if lavaDroplet.cubes.Contains(c) {
			steam.Add(currentCube)
		}
		coverWithSteam(c, lavaDroplet, steam, visited, boundingBox)
	}
}
