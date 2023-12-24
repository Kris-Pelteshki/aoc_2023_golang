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

type GameValues struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id      int
	results []GameValues
}

var Limits = GameValues{
	red:   12,
	green: 13,
	blue:  14,
}

func (g *Game) getMaxOfEachColor() (max GameValues) {
	for _, result := range g.results {
		if result.red > max.red {
			max.red = result.red
		}
		if result.green > max.green {
			max.green = result.green
		}
		if result.blue > max.blue {
			max.blue = result.blue
		}
	}

	return max
}

func (g *Game) isPossible() bool {
	for _, result := range g.results {
		if result.red > Limits.red || result.green > Limits.green || result.blue > Limits.blue {
			return false
		}
	}
	return true
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
	fmt.Println("Time", time.Since(startTime))
}

func part1(input string) (total int) {
	games := parseInput(input)

	for _, game := range games {
		if game.isPossible() {
			total += game.id
		}
	}

	return total
}

func part2(input string) (total int) {
	games := parseInput(input)

	for _, game := range games {
		maxValues := game.getMaxOfEachColor()
		total += maxValues.red * maxValues.green * maxValues.blue
	}

	return total
}

func parseInput(input string) (games []*Game) {
	for _, line := range strings.Split(input, "\n") {
		game := new(Game)

		gameStr, dataStr, found := strings.Cut(line, ":")
		if !found {
			panic("no colon in line")
		}

		_, err := fmt.Sscanf(gameStr, "Game %d", &game.id)
		if err != nil {
			panic(err)
		}

		for _, resultString := range strings.Split(dataStr, ";") {
			result := GameValues{}
			colors := map[string]*int{
				"red":   &result.red,
				"green": &result.green,
				"blue":  &result.blue,
			}
			values := strings.Split(resultString, ",")

			for _, value := range values {
				num, color, found := strings.Cut(strings.TrimSpace(value), " ")
				if !found {
					panic("no space in value")
				}

				*colors[color] = cast.ToInt(num)
			}

			game.results = append(game.results, result)
		}

		games = append(games, game)
	}
	return games
}
