package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/util"
)

type Grid struct {
	rows []string
}

func (g *Grid) isEqualRow(row1, row2 int) bool {
	return g.rows[row1] == g.rows[row2]
}

func (g *Grid) isEqualCol(col1, col2 int) bool {
	for i := 0; i < len(g.rows); i++ {
		if g.rows[i][col1] != g.rows[i][col2] {
			return false
		}
	}
	return true
}

func (g *Grid) isMirror(isEqual func(int, int) bool, length, idx int) bool {
	lastIdx := length - 1
	minIdx := idx - min(lastIdx-idx-1, idx)

	for i := minIdx; i <= idx; i++ {
		if !isEqual(i, getReflectedIndex(idx, i)) {
			return false
		}
	}

	return true
}

func (g *Grid) isMirrorAtRow(rowIdx int) bool {
	return g.isMirror(g.isEqualRow, len(g.rows), rowIdx)
}

func (g *Grid) isMirrorAtCol(colIdx int) bool {
	return g.isMirror(g.isEqualCol, len(g.rows[0]), colIdx)
}

func getReflectedIndex(inflectionPoint int, index int) int {
	return 2*inflectionPoint - index + 1
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
	grids := parseInput(input)
	rowsAboveMirror := 0
	colsBeforeMirror := 0

	for _, grid := range grids {
		for rowIdx := 0; rowIdx < len(grid.rows)-1; rowIdx++ {
			if grid.isMirrorAtRow(rowIdx) {
				rowsAboveMirror += rowIdx + 1
				break
			}
		}

		for colIdx := 0; colIdx < len(grid.rows[0])-1; colIdx++ {
			if grid.isMirrorAtCol(colIdx) {
				colsBeforeMirror += colIdx + 1
				break
			}
		}
	}

	return (rowsAboveMirror * 100) + colsBeforeMirror
}

func part2(input string) int {
	return 0
}

func parseInput(input string) []*Grid {
	grids := []*Grid{}

	for _, gridInput := range strings.Split(input, "\n\n") {
		grids = append(grids, &Grid{rows: strings.Split(gridInput, "\n")})
	}
	return grids
}
