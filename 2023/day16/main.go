package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/util"
	"github.com/emirpasic/gods/queues/arrayqueue"
)

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

func part1(input string) int {
	grid := parseInput(input)
	total := grid.simulateBeam(Beam{0, 0, Right})
	return total
}

func part2(input string) int {
	grid := parseInput(input)
	maxEnergy := 0
	entryPoints := []Beam{}

	// vertical beams
	for x := 0; x < grid.Width; x++ {
		entryPoints = append(entryPoints, Beam{x, 0, Down})
		entryPoints = append(entryPoints, Beam{x, grid.Height - 1, Up})
	}
	// horizontal beams
	for y := 0; y < grid.Height; y++ {
		entryPoints = append(entryPoints, Beam{0, y, Right})
		entryPoints = append(entryPoints, Beam{grid.Width - 1, y, Left})
	}

	for _, beam := range entryPoints {
		energy := grid.simulateBeam(beam)
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}

	return maxEnergy
}

const (
	emptySpace         = '.'
	backMirror         = '/'
	forwardMirror      = '\\'
	verticalSplitter   = '|'
	horizontalSplitter = '-'
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Beam struct {
	X, Y int
	Direction
}

func (beam *Beam) coords() [2]int {
	return [2]int{beam.X, beam.Y}
}

func (beam *Beam) move() {
	switch beam.Direction {
	case Up:
		beam.Y--
	case Right:
		beam.X++
	case Down:
		beam.Y++
	case Left:
		beam.X--
	}
}

type Grid struct {
	Width, Height int
	Cells         [][]rune
}

func (g *Grid) isInside(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func (g *Grid) simulateBeam(b Beam) int {
	beam := &b
	energizedTiles := make(map[[2]int]bool)
	visitedSplitters := make(map[[2]int]bool)

	queue := arrayqueue.New()
	queue.Enqueue(beam)

loop:
	for !queue.Empty() {
		elem, _ := queue.Dequeue()
		beam = elem.(*Beam)
		tile := g.Cells[beam.Y][beam.X]

		for tile == emptySpace {
			energizedTiles[beam.coords()] = true
			beam.move()
			if !g.isInside(beam.X, beam.Y) {
				continue loop
			}
			tile = g.Cells[beam.Y][beam.X]
		}

		coords := beam.coords()

		switch tile {
		case backMirror:
			switch beam.Direction {
			case Up:
				beam.Direction = Right
			case Right:
				beam.Direction = Up
			case Down:
				beam.Direction = Left
			default:
				beam.Direction = Down
			}
		case forwardMirror:
			switch beam.Direction {
			case Up:
				beam.Direction = Left
			case Right:
				beam.Direction = Down
			case Down:
				beam.Direction = Right
			default:
				beam.Direction = Up
			}
		case verticalSplitter:
			// Split the beam
			if beam.Direction == Right || beam.Direction == Left {
				if visitedSplitters[coords] {
					continue loop
				}
				visitedSplitters[coords] = true
				beam.Direction = Down
				newBeam := &Beam{beam.X, beam.Y - 1, Up}
				if g.isInside(newBeam.X, newBeam.Y) {
					queue.Enqueue(newBeam)
				}
			}
		case horizontalSplitter:
			// Split the beam
			if beam.Direction == Up || beam.Direction == Down {
				if visitedSplitters[coords] {
					continue loop
				}
				visitedSplitters[coords] = true
				newBeam := &Beam{beam.X - 1, beam.Y, Left}
				beam.Direction = Right
				if g.isInside(newBeam.X, newBeam.Y) {
					queue.Enqueue(newBeam)
				}
			}
		}

		energizedTiles[coords] = true
		beam.move()
		if g.isInside(beam.X, beam.Y) {
			queue.Enqueue(beam)
		}
	}

	return len(energizedTiles)
}

func parseInput(input string) *Grid {
	grid := &Grid{}
	lines := strings.Split(input, "\n")
	grid.Height = len(lines)
	grid.Width = len(lines[0])
	grid.Cells = make([][]rune, grid.Height)

	for i, line := range lines {
		grid.Cells[i] = make([]rune, grid.Width)
		for j, r := range line {
			grid.Cells[i][j] = r
		}
	}
	return grid
}
