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
	"github.com/emirpasic/gods/stacks/arraystack"
)

type Interval struct {
	start uint
	end   uint
}

type MapRange struct {
	source      Interval
	destination Interval
}

func (mr *MapRange) contains(value uint) bool {
	return value >= mr.source.start && value <= mr.source.end
}

func (mr *MapRange) getDestinationValue(value uint) uint {
	return mr.destination.start + (value - mr.source.start)
}

func newMapRange(destinationStart, sourceStart, len uint) MapRange {
	return MapRange{
		source: Interval{
			start: sourceStart,
			end:   sourceStart + len - 1,
		},
		destination: Interval{
			start: destinationStart,
			end:   destinationStart + len - 1,
		},
	}
}

type ConversionMap struct {
	ranges []MapRange
}

func (cm *ConversionMap) mapValue(value uint) uint {
	for _, mapRange := range cm.ranges {
		if mapRange.contains(value) {
			return mapRange.getDestinationValue(value)
		}
	}
	return value
}

func (cm *ConversionMap) splitIntervals(intervals []Interval) []Interval {
	var result []Interval
	stack := arraystack.New()

	for _, interval := range intervals {
		stack.Push(interval)
	}

OuterLoop:
	for !stack.Empty() {
		val, _ := stack.Pop()
		interval := val.(Interval)

		for _, mapRange := range cm.ranges {
			min := maths.Max(interval.start, mapRange.source.start)
			max := maths.Min(interval.end, mapRange.source.end)

			if min <= max {
				result = append(result, Interval{min, max})

				if min > interval.start {
					stack.Push(Interval{interval.start, min - 1})
				}
				if max < interval.end {
					stack.Push(Interval{max + 1, interval.end})
				}

				continue OuterLoop
			}
		}
		result = append(result, interval)
	}

	return result
}

func (cm *ConversionMap) mapIntervals(intervals []Interval) []Interval {
	var result []Interval

	for _, interval := range intervals {
		inRange := false
		for _, mapRange := range cm.ranges {
			if mapRange.contains(interval.start) {
				inRange = true
				result = append(result, Interval{
					mapRange.getDestinationValue(interval.start),
					mapRange.getDestinationValue(interval.end),
				})
			}
		}

		if !inRange {
			result = append(result, interval)
		}
	}
	return result
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

func part1(input string) uint {
	seeds, maps := parseInput(input)
	lowestLocations := []uint{}

	for _, seed := range seeds {
		ans := seed
		for _, convMap := range maps {
			ans = convMap.mapValue(ans)
		}
		lowestLocations = append(lowestLocations, ans)
	}

	return maths.Min(lowestLocations...)
}

func part2(input string) uint {
	seeds, maps := parseInput(input)
	seedPairs := collections.Chunks(seeds, 2)
	minLocation := seeds[0]

	for _, seedPair := range seedPairs {
		start := seedPair[0]
		end := start + seedPair[1] - 1
		intervals := []Interval{{start, end}}

		for _, convMap := range maps {
			intervals = convMap.splitIntervals(intervals)
			intervals = convMap.mapIntervals(intervals)
		}

		for _, interval := range intervals {
			if interval.start < minLocation {
				minLocation = interval.start
			}
		}
	}

	return minLocation
}

func buildMapRange(input string) *ConversionMap {
	conversionMap := ConversionMap{}

	_, data, _ := strings.Cut(input, " map:\n")

	for _, line := range strings.Split(data, "\n") {
		nums := strings.Fields(line)
		mapRange := newMapRange(
			uint(cast.ToInt(nums[0])),
			uint(cast.ToInt(nums[1])),
			uint(cast.ToInt(nums[2])),
		)
		conversionMap.ranges = append(conversionMap.ranges, mapRange)
	}

	sort.Slice(conversionMap.ranges, func(i, j int) bool {
		return conversionMap.ranges[i].destination.start < conversionMap.ranges[j].destination.start
	})

	return &conversionMap
}

func parseInput(input string) (seeds []uint, conversionMaps []*ConversionMap) {
	seperator := "\n\n"

	seedLine, dataLine, _ := strings.Cut(input, seperator)
	seedStrings := strings.Fields(seedLine)[1:]

	for _, seedStr := range seedStrings {
		seeds = append(seeds, uint(cast.ToInt(seedStr)))
	}
	for _, mapStr := range strings.Split(dataLine, seperator) {
		conversionMaps = append(conversionMaps, buildMapRange(mapStr))
	}

	return seeds, conversionMaps
}
