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
	for x := uint8(0); x < grid.Width; x++ {
		entryPoints = append(entryPoints, Beam{x, 0, Down})
		entryPoints = append(entryPoints, Beam{x, grid.Height - 1, Up})
	}
	// horizontal beams
	for y := uint8(0); y < grid.Height; y++ {
		entryPoints = append(entryPoints, Beam{0, y, Right})
		entryPoints = append(entryPoints, Beam{grid.Width - 1, y, Left})
	}

	// Create a channel to parallelize
	energyChan := make(chan int)

	for _, beam := range entryPoints {
		go func(beam Beam) {
			energy := grid.simulateBeam(beam)
			energyChan <- energy
		}(beam)
	}

	for range entryPoints {
		energy := <-energyChan
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

type Direction uint8

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Beam struct {
	X, Y uint8
	Direction
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
	Width, Height uint8
	Cells         [][]rune
}

func parseInput(input string) *Grid {
	lines := strings.Split(input, "\n")
	grid := &Grid{
		Height: uint8(len(lines)),
		Width:  uint8(len(lines[0])),
	}
	grid.Cells = make([][]rune, grid.Height)

	for i, line := range lines {
		grid.Cells[i] = []rune(line)
	}
	return grid
}

func (g *Grid) isInside(x, y uint8) bool {
	return x < g.Width && y < g.Height
}

func (grid *Grid) simulateBeam(b Beam) int {
	beam := &b
	energizedTiles := 0

	visited := make([][][4]bool, grid.Height)
	for i := range visited {
		visited[i] = make([][4]bool, grid.Width)
	}

	queue := arrayqueue.New()
	queue.Enqueue(beam)

	for !queue.Empty() {
		elem, _ := queue.Dequeue()
		beam = elem.(*Beam)

		if !grid.isInside(beam.X, beam.Y) {
			continue
		}

		visitedPos := visited[beam.Y][beam.X]
		if visitedPos[beam.Direction] {
			continue
		}
		visited[beam.Y][beam.X][beam.Direction] = true

		energizedTiles++
		for _, b := range visitedPos {
			if b {
				energizedTiles--
				break
			}
		}

		switch grid.Cells[beam.Y][beam.X] {
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
			if beam.Direction == Right || beam.Direction == Left {
				beam.Direction = Down
				newBeam := &Beam{beam.X, beam.Y - 1, Up}
				queue.Enqueue(newBeam)
			}
		case horizontalSplitter:
			if beam.Direction == Up || beam.Direction == Down {
				newBeam := &Beam{beam.X - 1, beam.Y, Left}
				beam.Direction = Right
				queue.Enqueue(newBeam)
			}
		}

		beam.move()
		queue.Enqueue(beam)
	}

	return energizedTiles
}
