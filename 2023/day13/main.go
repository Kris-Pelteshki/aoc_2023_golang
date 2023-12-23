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

func (g *Grid) countDiffBetweenRows(row1, row2 int) (count int) {
	for i := 0; i < len(g.rows[row1]); i++ {
		if g.rows[row1][i] != g.rows[row2][i] {
			count++
		}
	}
	return count
}

func (g *Grid) countDiffBetweenCols(col1, col2 int) (count int) {
	for i := 0; i < len(g.rows); i++ {
		if g.rows[i][col1] != g.rows[i][col2] {
			count++
		}
	}
	return count
}

func (g *Grid) isMirrorRows(idx int, maxDiff int) (diffs int) {
	return g.isMirror(g.countDiffBetweenRows, len(g.rows), idx, maxDiff)
}

func (g *Grid) isMirrorCols(idx int, maxDiff int) (diffs int) {
	return g.isMirror(g.countDiffBetweenCols, len(g.rows[0]), idx, maxDiff)
}

func (g *Grid) isMirror(
	calcDiff func(int, int) int,
	length int,
	idx int,
	maxDiff int,
) (diffs int) {
	lastIdx := length - 1
	minIdx := idx - min(lastIdx-idx-1, idx)

	for i := minIdx; i <= idx; i++ {
		diffs += calcDiff(i, getReflectedIndex(idx, i))
		if diffs > maxDiff {
			break
		}
	}
	return diffs
}

func (g *Grid) findMirrorRow(maxDiff int) int {
	for rowIdx := 0; rowIdx < len(g.rows)-1; rowIdx++ {
		if g.isMirrorRows(rowIdx, maxDiff) == maxDiff {
			return rowIdx + 1
		}
	}
	return 0
}

func (g *Grid) findMirrorCol(maxDiff int) int {
	for colIdx := 0; colIdx < len(g.rows[0])-1; colIdx++ {
		if g.isMirrorCols(colIdx, maxDiff) == maxDiff {
			return colIdx + 1
		}
	}
	return 0
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
		rowsAboveMirror += grid.findMirrorRow(0)
		colsBeforeMirror += grid.findMirrorCol(0)
	}
	return (rowsAboveMirror * 100) + colsBeforeMirror
}

func part2(input string) (total int) {
	grids := parseInput(input)
	rowsAboveMirror := 0
	colsBeforeMirror := 0

	for _, grid := range grids {
		rowsAboveMirror += grid.findMirrorRow(1)
		colsBeforeMirror += grid.findMirrorCol(1)
	}
	return (rowsAboveMirror * 100) + colsBeforeMirror
}

func parseInput(input string) []*Grid {
	grids := []*Grid{}

	for _, gridInput := range strings.Split(input, "\n\n") {
		grids = append(grids, &Grid{rows: strings.Split(gridInput, "\n")})
	}
	return grids
}
