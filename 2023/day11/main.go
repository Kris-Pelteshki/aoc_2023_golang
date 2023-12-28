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

var galaxySymbol = '#'

type Galaxy struct {
	X, Y int
}

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
	return calcTotalDistancesOfPairs(input, 2)
}

func part2(input string) int {
	return calcTotalDistancesOfPairs(input, 1000000)
}

func calcTotalDistancesOfPairs(input string, expandEmptySpaceByFactorOf int) (total int) {
	universe := getUniverse(input, expandEmptySpaceByFactorOf)

	for i, galaxy := range universe {
		for _, otherGalaxy := range universe[i+1:] {
			dx, dy := galaxy.distanceTo(otherGalaxy)
			total += dx + dy
		}
	}

	return total
}

func buildGrid(input string, expandEmptySpaceByFactorOf int) (*[]string, []int, []int) {
	grid := util.SplitLines(input)

	rowIndexToCoords := make([]int, len(grid))
	colIndexToCoords := make([]int, len(grid[0]))
	emptyRowCount := 0
	emptyColCount := 0

	for y, line := range grid {
		rowIndexToCoords[y] = emptyRowCount*(expandEmptySpaceByFactorOf-1) + y
		if !strings.Contains(line, "#") {
			emptyRowCount++
		}
	}

	for x := range grid[0] {
		containsGalaxy := false
		for y := range grid {
			if grid[y][x] == '#' {
				containsGalaxy = true
				break
			}
		}

		colIndexToCoords[x] = emptyColCount*(expandEmptySpaceByFactorOf-1) + x
		if !containsGalaxy {
			emptyColCount++
		}
	}

	grid = slices.Clip(grid)
	return &grid, rowIndexToCoords, colIndexToCoords
}

// parses the input into galaxies
func getUniverse(input string, expandEmptySpaceByFactorOf int) (galaxies []*Galaxy) {
	grid, rowIndexToCoords, colIndexToCoords := buildGrid(input, expandEmptySpaceByFactorOf)

	for y, line := range *grid {
		for x, char := range line {
			if char == galaxySymbol {
				galaxies = append(galaxies, &Galaxy{colIndexToCoords[x], rowIndexToCoords[y]})
			}
		}
	}

	return galaxies
}
