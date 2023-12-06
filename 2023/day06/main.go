package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/util"
)

type raceRecord struct {
	time     int64
	distance int64
}

const IncrementPerMilisecond = 1

func isBetterDistance(holdTime int64, race *raceRecord) bool {
	remainingTime := race.time - holdTime
	speed := holdTime * IncrementPerMilisecond
	return remainingTime*speed > race.distance
}

func findMin(race *raceRecord) (min int64) {
	min = 0
	max := race.time

	for min < max {
		mid := min + (max-min)/2
		if isBetterDistance(mid, race) {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return min
}

func findMax(race *raceRecord) (max int64) {
	var min int64 = 0
	max = race.time

	for min < max {
		mid := min + (max-min)/2
		if isBetterDistance(mid, race) {
			min = mid + 1
		} else {
			max = mid
		}
	}

	return max - 1
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

func part1(input string) (total int64) {
	races := parseInput(input)
	total = 1

	for _, race := range races {
		minHoldTime := findMin(&race)
		maxHoldTime := findMax(&race)
		total *= maxHoldTime - minHoldTime + 1
	}

	return total
}

func part2(input string) int64 {
	race := parsePart2(input)
	log.Println(race)

	minHoldTime := findMin(&race)
	maxHoldTime := findMax(&race)

	return maxHoldTime - minHoldTime + 1
}

func parseInput(input string) (races []raceRecord) {
	lines := strings.Split(input, "\n")
	for lineIdx, line := range lines {
		_, dataStr, ok := strings.Cut(line, ":")
		if !ok {
			panic("invalid input")
		}

		nums := strings.Fields(dataStr)

		if len(races) == 0 {
			races = make([]raceRecord, len(nums))
		}

		for i, numStr := range nums {
			if lineIdx == 0 {
				races[i].time, _ = strconv.ParseInt(numStr, 10, 64)
			} else {
				races[i].distance, _ = strconv.ParseInt(numStr, 10, 64)
			}
		}
	}
	return races
}

func parsePart2(input string) (race raceRecord) {
	lines := strings.Split(input, "\n")

	for lineIdx, line := range lines {
		_, dataStr, ok := strings.Cut(line, ":")
		if !ok {
			panic("invalid input")
		}

		num := strings.Join(strings.Fields(dataStr), "")

		if lineIdx == 0 {
			race.time, _ = strconv.ParseInt(num, 10, 64)
		} else {
			race.distance, _ = strconv.ParseInt(num, 10, 64)
		}
	}
	return race
}
