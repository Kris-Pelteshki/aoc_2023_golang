package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/util"
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

type rock byte
type platform [][]rock
type direction int

const (
	rounded rock = 'O'
	cube    rock = '#'
	empty   rock = '.'
)

const (
	north direction = iota
	south
	east
	west
)

func part1(input string) int {
	platform := parseInput(input)
	platform.tilt(north)
	return platform.load()
}

// TODO:
// learn about cycle detection algorithms
func part2(input string) int {
	platform := parseInput(input)
	for i := 0; i < 1000; i++ {
		platform.cycle()
	}
	return platform.load()
}

func parseInput(input string) (p platform) {
	for _, line := range util.SplitLines(input) {
		row := make([]rock, len(line))
		for i := range line {
			switch line[i] {
			case 'O':
				row[i] = rounded
			case '#':
				row[i] = cube
			case '.':
				row[i] = empty
			}
		}
		p = append(p, row)
	}
	return p
}

func (p *platform) cycle() {
	p.tilt(north)
	p.tilt(west)
	p.tilt(south)
	p.tilt(east)
}

// Tilt the platform in the given direction and return the new platform
func (p *platform) tilt(dir direction) {
	platform := *p
	switch dir {
	case north:
		// Use a sliding window from top to bottom
		// use 2 pointers for a sliding window
		// pointer 1 will keep track of the last known empty position
		// pointer 2 will keep track of the current position
		// when pointer 2 finds a rounded rock, swap it with the empty spot at pointer 1
		// then increment pointer 1
		// this will move all the rocks to the top of the platform
		// if pointer 2 encounters a cube, move pointer 1 to the first empty spot above pointer 2
		// then increment pointer 2
		for j := 0; j < len(platform[0]); j++ {
			empty := 0

			// pointer 2
			for i := 0; i < len(platform); i++ {
				if platform[i][j] == rounded {
					platform[empty][j], platform[i][j] = platform[i][j], platform[empty][j]
					empty++
				} else if platform[i][j] == cube {
					empty = i + 1
				}
			}
		}

	case south:
		for j := 0; j < len(platform[0]); j++ {
			empty := len(platform) - 1

			for i := len(platform) - 1; i >= 0; i-- {
				if platform[i][j] == rounded {
					platform[empty][j], platform[i][j] = platform[i][j], platform[empty][j]
					empty--
				} else if platform[i][j] == cube {
					empty = i - 1
				}
			}
		}

	case east:
		for i := 0; i < len(platform); i++ {
			empty := len(platform[i]) - 1

			for j := len(platform[i]) - 1; j >= 0; j-- {
				if platform[i][j] == rounded {
					platform[i][empty], platform[i][j] = platform[i][j], platform[i][empty]
					empty--
				} else if platform[i][j] == cube {
					empty = j - 1
				}
			}
		}

	case west:
		for i := 0; i < len(platform); i++ {
			empty := 0

			for j := 0; j < len(platform[i]); j++ {
				if platform[i][j] == rounded {
					platform[i][empty], platform[i][j] = platform[i][j], platform[i][empty]
					empty++
				} else if platform[i][j] == cube {
					empty = j + 1
				}
			}
		}
	}
}

// Calculate the total load on the north support beams
func (p *platform) load() (load int) {
	platform := *p

	for i := range platform {
		for j := range platform[i] {
			if platform[i][j] == rounded {
				// Add the number of rows from the rock to the south edge
				load += len(platform) - i
			}
		}
	}
	return load
}

// Print the platform as a string
func (platform *platform) String() string {
	p := *platform
	var sb strings.Builder

	sb.WriteByte('\n')
	for i := range p {
		for j := range p[i] {
			sb.WriteByte(byte(p[i][j]))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
