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

func coordsToKey(x, y int) string {
	return fmt.Sprintf("%v, %v", x, y)
}

func keyToCoords(key string) (x, y int) {
	fmt.Sscanf(key, "%v, %v", &x, &y)
	return x, y
}

func (p *Point[T]) String() string {
	return coordsToKey(p.x, p.y)
}

type coordId = string
type numberId = coordId
type numbersCoordMap = map[numberId][]coordId
type symbolMap = map[string]Point[string]
type numbersMap = map[string]Point[int]

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
	grid := buildGrid(input)
	symbolMap, numbersMap, coordsMap := buildMaps(grid)
	nums := []int{}

	for _, numPoint := range numbersMap {
		if hasAdjacentSymbol(&symbolMap, &coordsMap, &numPoint) {
			nums = append(nums, numPoint.value)
		}
	}

	for _, num := range nums {
		total += num
	}
	return total
}

func part2(input string) (total int) {
	grid := buildGrid(input)
	symbolMap, numbersMap, coordsMap := buildMaps(grid)
	nums := []int{}

	for _, point := range symbolMap {
		if point.value != "*" {
			delete(symbolMap, point.String())
		}
	}

	for _, symbol := range symbolMap {
		num1, num2, ok := getTwoAdjacentNumbers(&numbersMap, &coordsMap, &symbol)
		if ok {
			nums = append(nums, num1*num2)
		}
	}

	for _, num := range nums {
		total += num
	}
	return total
}

func buildGrid(input string) (grid *Grid) {
	return &Grid{
		rows: strings.Split(input, "\n"),
	}
}

func buildMaps(grid *Grid) (symbols symbolMap, numbers numbersMap, coordsMap numbersCoordMap) {
	symbols = make(symbolMap)
	numbers = make(numbersMap)
	coordsMap = make(numbersCoordMap)

	for y, row := range grid.rows {
		for x := 0; x < len(row); {
			char := rune(row[x])

			if unicode.IsDigit(char) {
				temp := ""
				coords := []coordId{}

				for idx, char := range row[x:] {
					if !unicode.IsDigit(char) {
						break
					}
					coords = append(coords, coordsToKey(x+idx, y))
					temp += string(char)
				}

				point := Point[int]{
					cast.ToInt(temp),
					x,
					y,
				}
				pointId := point.String()
				numbers[pointId] = point
				coordsMap[pointId] = coords

				x += len(temp)
				continue
			}

			if isSymbol(char) {
				point := Point[string]{
					string(char),
					x,
					y,
				}
				symbols[point.String()] = point
			}

			x++
		}
	}

	return symbols, numbers, coordsMap
}

func hasAdjacentSymbol[Tpoint any](symbolMap *symbolMap, coordsMap *numbersCoordMap, p *Point[Tpoint]) bool {
	found := false

	coords, coordsOk := (*coordsMap)[p.String()]
	if !coordsOk {
		return false
	}

	for _, coord := range coords {
		px, py := keyToCoords(coord)

		for y := py - 1; y <= py+1; y++ {
			for x := px - 1; x <= px+1; x++ {
				if x == px && y == py {
					continue
				}

				pointId := coordsToKey(x, y)

				_, ok := (*symbolMap)[pointId]
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
	seenNumIds := []string{}

	// brute force it for now
	getNumberId := func(x, y int) string {
		for id, coords := range *coordsMap {
			for _, coord := range coords {
				px, py := keyToCoords(coord)
				if px == x && py == y {
					return id
				}
			}
		}
		return ""
	}

	for y := p.y - 1; y <= p.y+1; y++ {
		for x := p.x - 1; x <= p.x+1; x++ {
			if x == p.x && y == p.y {
				continue
			}

			numId := getNumberId(x, y)

			if numId != "" && slices.Contains(seenNumIds, numId) {
				continue
			}

			numPoint, ok := (*numbersMap)[numId]
			if ok {
				seenNumIds = append(seenNumIds, numId)
				adjacentNums = append(adjacentNums, numPoint.value)
			}
		}
	}

	if len(adjacentNums) == 2 {
		return adjacentNums[0], adjacentNums[1], true
	}

	return 0, 0, false
}
