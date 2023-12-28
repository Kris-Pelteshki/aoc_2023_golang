package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
	"time"
	"unicode"

	"github.com/Kris-Pelteshki/aoc_2023/cast"
	"github.com/Kris-Pelteshki/aoc_2023/util"
)

type Point[T any] struct {
	value T
	x     int
	y     int
}

func (p *Point[T]) getCoords() coordKey {
	return coordKey{p.x, p.y}
}

type coordKey = [2]int
type numbersCoordMap = map[coordKey][]coordKey
type symbolMap = map[coordKey]*Point[rune]
type numbersMap = map[coordKey]*Point[int]

type Grid struct {
	rows []string
}

func isSymbol(r rune) bool {
	return r != '.' && !unicode.IsDigit(r)
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

	startTime := time.Now()

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
	fmt.Printf("Time taken: %v\n", time.Since(startTime))
}

func part1(input string) (total int) {
	symbolMap, numbersMap, coordsMap := buildMaps(&input)

	for _, point := range numbersMap {
		if hasAdjacentSymbol(&symbolMap, &coordsMap, point) {
			total += point.value
		}
	}
	return total
}

func part2(input string) (total int) {
	symbolMap, numbersMap, coordsMap := buildMaps(&input)

	for _, point := range symbolMap {
		if point.value != '*' {
			delete(symbolMap, point.getCoords())
		}
	}

	for _, symbol := range symbolMap {
		num1, num2, ok := getTwoAdjacentNumbers(&numbersMap, &coordsMap, symbol)
		if ok {
			total += num1 * num2
		}
	}
	return total
}

func buildMaps(input *string) (symbols symbolMap, numbers numbersMap, coordsMap numbersCoordMap) {
	grid := &Grid{
		rows: strings.Split(*input, "\n"),
	}

	symbols = make(symbolMap)
	numbers = make(numbersMap)
	coordsMap = make(numbersCoordMap)

	for y, row := range grid.rows {
		for x := 0; x < len(row); {
			char := rune(row[x])

			if unicode.IsDigit(char) {
				temp := ""
				coords := []coordKey{}

				for idx, char := range row[x:] {
					if !unicode.IsDigit(char) {
						break
					}
					coords = append(coords, coordKey{x + idx, y})
					temp += string(char)
				}

				point := Point[int]{cast.ToInt(temp), x, y}
				numbers[point.getCoords()] = &point
				coordsMap[point.getCoords()] = coords

				x += len(temp)
				continue
			}

			if isSymbol(char) {
				symbols[coordKey{x, y}] = &Point[rune]{char, x, y}
			}

			x++
		}
	}

	return symbols, numbers, coordsMap
}

func hasAdjacentSymbol[Tpoint any](symbolMap *symbolMap, coordsMap *numbersCoordMap, p *Point[Tpoint]) bool {
	found := false

	coords, coordsOk := (*coordsMap)[coordKey{p.x, p.y}]
	if !coordsOk {
		return false
	}

	for _, coord := range coords {
		px, py := coord[0], coord[1]

		for y := py - 1; y <= py+1; y++ {
			for x := px - 1; x <= px+1; x++ {
				if x == px && y == py {
					continue
				}

				_, ok := (*symbolMap)[coordKey{x, y}]
				if ok {
					found = true
					break
				}
			}
		}
	}

	return found
}

func getTwoAdjacentNumbers[Tpoint any](numbersMap *numbersMap, coordsMap *numbersCoordMap, p *Point[Tpoint]) (int, int, bool) {
	adjacentNums := []int{}
	seenNumIds := []coordKey{}

	getNumberCoords := func(x, y int) (coordKey, bool) {
		for key, coords := range *coordsMap {
			for _, coord := range coords {
				px, py := coord[0], coord[1]
				if px == x && py == y {
					return key, true
				}
			}
		}
		return coordKey{x, y}, false
	}

	for y := p.y - 1; y <= p.y+1; y++ {
		for x := p.x - 1; x <= p.x+1; x++ {
			if x == p.x && y == p.y {
				continue
			}

			numKey, found := getNumberCoords(x, y)

			if found && slices.Contains(seenNumIds, numKey) {
				continue
			}

			numPoint, ok := (*numbersMap)[numKey]
			if ok {
				seenNumIds = append(seenNumIds, numKey)
				adjacentNums = append(adjacentNums, numPoint.value)
			}
		}
	}

	if len(adjacentNums) == 2 {
		return adjacentNums[0], adjacentNums[1], true
	}

	return 0, 0, false
}
