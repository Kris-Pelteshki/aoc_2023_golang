package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/cast"
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

func nextSequence(seq []int) ([]int, bool) {
	nextSequenceLen := len(seq) - 1
	nextSequence := make([]int, 0, nextSequenceLen)
	hasNonZero := false

	for i := 0; i < nextSequenceLen; i++ {
		nextValue := seq[i+1] - seq[i]
		if !hasNonZero && nextValue != 0 {
			hasNonZero = true
		}
		nextSequence = append(nextSequence, nextValue)
	}
	return nextSequence, !hasNonZero
}

func nextNumInSequence(seq []int) (nextNum int) {
	allZeros := false
	for !allZeros {
		nextNum += seq[len(seq)-1]
		seq, allZeros = nextSequence(seq)
	}
	return nextNum
}

func prevNumInSequence(seq []int) (prevNum int) {
	allZeros := false
	firstNumsInSequences := []int{}

	for !allZeros {
		firstNumsInSequences = append(firstNumsInSequences, seq[0])
		seq, allZeros = nextSequence(seq)
	}

	for i := len(firstNumsInSequences) - 1; i >= 0; i-- {
		prevNum = firstNumsInSequences[i] - prevNum
	}

	return prevNum
}

func part1(input string) (total int) {
	sequences := parseInput(input)

	for _, seq := range sequences {
		total += nextNumInSequence(seq)
	}

	return total
}

func part2(input string) (total int) {
	sequences := parseInput(input)

	for _, seq := range sequences {
		total += prevNumInSequence(seq)
	}

	return total
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, num := range strings.Fields(line) {
			row = append(row, cast.ToInt(num))
		}
		ans = append(ans, row)
	}
	return ans
}
