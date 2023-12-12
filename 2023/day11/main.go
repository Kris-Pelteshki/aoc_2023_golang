package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/util"
	"github.com/Kris-Pelteshki/aoc_2023/util/maths"
)

type Point struct {
	X, Y int
}
type Id = int

type Galaxy struct {
	Point
	ID Id
}

type Universe map[Id]*Galaxy

type GalaxyPairs map[Id][]*Galaxy

func (g *Galaxy) distanceTo(other *Galaxy) (dx, dy int) {
	dx = maths.Abs(other.X - g.X)
	dy = maths.Abs(other.Y - g.Y)
	return dx, dy
}

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	start := time.Now()

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}

	fmt.Println("Time:", time.Since(start))
}

func part1(input string) (total int) {
	universe := getUniverse(input)
	galaxyPairs := make(GalaxyPairs)
	visited := make(map[Id]bool)

	// adjavency list
	for _, galaxy := range universe {
		galaxyPairs[galaxy.ID] = make([]*Galaxy, 0)
		visited[galaxy.ID] = true
		for _, otherGalaxy := range universe {
			if !visited[otherGalaxy.ID] {
				galaxyPairs[galaxy.ID] = append(galaxyPairs[galaxy.ID], otherGalaxy)
			}
		}
	}

	for id, galaxyList := range galaxyPairs {
		node := universe[id]
		for _, otherGalaxy := range galaxyList {
			dx, dy := node.distanceTo(otherGalaxy)
			total += dx + dy
		}
	}

	return total
}

func part2(input string) int {
	return 0
}

func buildGrid(input string) *[]string {
	grid := util.SplitLines(input)

	insertCount := 0
	for y, line := range grid {
		if !strings.Contains(line, "#") {
			grid = slices.Insert(grid, y+insertCount, line)
			insertCount++
		}
	}

	insertCount = 0
	for x := range grid[0] {
		col := ""
		for y := range grid {
			col += string(grid[y][x+insertCount])
		}
		if !strings.Contains(col, "#") {
			for y := range grid {
				grid[y] = grid[y][:x+insertCount] + "." + grid[y][x+insertCount:]
			}
			insertCount++
		}
	}

	clipped := slices.Clip(grid)
	return &clipped
}

// parses the input into galaxies
func getUniverse(input string) Universe {
	galaxies := make(Universe)
	grid := buildGrid(input)
	id := 0

	for y, line := range *grid {
		for x, char := range line {
			if char == '#' {
				galaxies[id] = &Galaxy{Point{x, y}, id}
				id++
			}
		}
	}

	return galaxies
}
