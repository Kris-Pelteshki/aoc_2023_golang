package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/util"
	"github.com/Kris-Pelteshki/aoc_2023/util/maths"
)

type DirectionTuple = [2]string
type LocationLookup = map[string]DirectionTuple

var Directions = map[string]int{
	"L": 0,
	"R": 1,
}

type PathFinder struct {
	lookup       LocationLookup
	instructions string
}

func (pf *PathFinder) getStepCountBetween(startLoc, endLoc string) (steps int) {
	getNextDir := newInstructionIterator(pf.instructions)
	current := startLoc

	for current != endLoc {
		steps++
		edges, hasLoc := pf.lookup[current]
		dir := getNextDir()

		if !hasLoc {
			log.Fatalf("unknown location: %v", current)
		}

		current = edges[Directions[dir]]
	}

	return steps
}

func (pf *PathFinder) getStepCountBetweenFunc(startLoc string, foundEnd func(loc *string) bool) (steps int, endLocation string) {
	getNextDir := newInstructionIterator(pf.instructions)
	current := startLoc

	for !foundEnd(&current) {
		steps++
		edges, hasLoc := pf.lookup[current]
		dir := getNextDir()

		if !hasLoc {
			log.Fatalf("unknown location: %v", current)
		}

		current = edges[Directions[dir]]
	}

	return steps, current
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

func part1(input string) int {
	pf := parseInput(input)
	return pf.getStepCountBetween("AAA", "ZZZ")
}

func part2(input string) int {
	pf := parseInput(input)
	locsEndingInA := []string{}
	distances := []int{}

	for loc := range pf.lookup {
		if endsWithA(&loc) {
			locsEndingInA = append(locsEndingInA, loc)
		}
	}

	for _, locA := range locsEndingInA {
		distance, _ := pf.getStepCountBetweenFunc(locA, endsWithZ)
		distances = append(distances, distance)
	}

	return maths.LCM(distances...)
}

func endsWithA(loc *string) bool {
	return (*loc)[len((*loc))-1:] == "A"
}

func endsWithZ(loc *string) bool {
	return (*loc)[len((*loc))-1:] == "Z"
}

func newInstructionIterator(instructions string) func() string {
	i := 0
	maxLen := len(instructions)

	return func() string {
		if i >= maxLen {
			i = 0
		}
		dir := instructions[i : i+1]
		i++
		return dir
	}
}

func parseInput(input string) *PathFinder {
	lookup := make(LocationLookup)
	instructions, locationLines, _ := strings.Cut(input, "\n\n")

	for _, line := range strings.Split(locationLines, "\n") {
		split := strings.Split(line, " = ")
		dirsStr := strings.Trim(split[1], "()")
		dirs := strings.Split(dirsStr, ", ")
		lookup[split[0]] = [2]string{dirs[0], dirs[1]}
	}

	return &PathFinder{lookup, instructions}
}
