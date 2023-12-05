package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/cast"
	"github.com/Kris-Pelteshki/aoc_2023/util"
	"github.com/Kris-Pelteshki/aoc_2023/util/collections"
	"github.com/Kris-Pelteshki/aoc_2023/util/maths"
)

type mapRangeType struct {
	sourceStart      uint
	sourceEnd        uint
	length           uint
	destinationStart uint
}

type conversionMap struct {
	from   string
	to     string
	ranges []mapRangeType
}

func newMapRange(destinationStart, start, len uint) mapRangeType {
	return mapRangeType{start, start + len, len, destinationStart}
}

type SeedStrategy = func(string) []uint

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

func mapValue(value uint, conversionMap conversionMap) uint {
	for _, mapRange := range conversionMap.ranges {
		if value >= mapRange.sourceStart && value < mapRange.sourceEnd {
			return mapRange.destinationStart + (value - mapRange.sourceStart)
		}
	}

	return value
}

func part1(input string) uint {
	seeds, maps := parseInput(input)
	lowestLocations := []uint{}

	for _, seed := range seeds {
		ans := seed
		for _, convMap := range maps {
			ans = mapValue(ans, convMap)
		}
		lowestLocations = append(lowestLocations, ans)
	}

	return maths.Min(lowestLocations...)
}

/*
* Brute forcing the solution because I'm too tired to reword part 1 to work for part 2.
* A better way would be to perform a reverse search from the lowest location to the seed... but that's for another time
 */
func part2(input string) uint {
	seeds, maps := parseInput(input)
	minLocation := seeds[0]

	seedPairs := collections.Chunks(seeds, 2)

	for _, seedPair := range seedPairs {
		start := seedPair[0]
		end := start + seedPair[1]

		for i := start; i < end; i++ {
			ans := i
			for _, convMap := range maps {
				ans = mapValue(ans, convMap)
			}

			if ans < minLocation {
				minLocation = ans
			}
		}
	}

	return minLocation
}

func buildMapRange(input string) conversionMap {
	ans := conversionMap{}

	name, data, _ := strings.Cut(input, " map:\n")

	from, to, isNameOk := strings.Cut(name, "-to-")
	if !isNameOk {
		panic("invalid name")
	}

	ans.from = from
	ans.to = to

	for _, line := range strings.Split(data, "\n") {
		nums := strings.Fields(line)
		mapRange := newMapRange(
			uint(cast.ToInt(nums[0])),
			uint(cast.ToInt(nums[1])),
			uint(cast.ToInt(nums[2])),
		)
		ans.ranges = append(ans.ranges, mapRange)
	}

	sort.Slice(ans.ranges, func(i, j int) bool {
		return ans.ranges[i].destinationStart < ans.ranges[j].destinationStart
	})

	return ans
}

func parseInput(input string) (seeds []uint, conversionMaps []conversionMap) {
	seperator := "\n\n"

	seedLine, dataLine, _ := strings.Cut(input, seperator)
	seedStrData := strings.Split(seedLine, ": ")

	for _, seedStr := range strings.Fields(seedStrData[1]) {
		seeds = append(seeds, uint(cast.ToInt(seedStr)))
	}

	for _, mapStr := range strings.Split(dataLine, seperator) {
		conversionMaps = append(conversionMaps, buildMapRange(mapStr))
	}

	return seeds, conversionMaps
}
