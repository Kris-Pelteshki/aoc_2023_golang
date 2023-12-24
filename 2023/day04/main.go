package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/cast"
	"github.com/Kris-Pelteshki/aoc_2023/util"
)

type Card struct {
	id               int
	potentialWinners []int
	nums             []int
}

func (c *Card) getWinningNumbers() (winners []int) {
	for _, num := range c.nums {
		if slices.Contains(c.potentialWinners, num) {
			winners = append(winners, num)
		}
	}
	return winners
}

func cardScore(amountOfWinningNumbers int) (score int) {
	if amountOfWinningNumbers == 0 {
		return 0
	}
	if amountOfWinningNumbers == 1 {
		return 1
	}
	exponent := amountOfWinningNumbers - 1
	return int(math.Pow(2, float64(exponent)))
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

func part1(input string) (total int) {
	cards := parseInput(input)

	for _, c := range cards {
		winners := c.getWinningNumbers()
		total += cardScore(len(winners))
	}

	return total
}

func part2(input string) (total int) {
	cards := parseInput(input)
	cardCountMap := make(map[int]int)

	for _, c := range cards {
		count, hasCount := cardCountMap[c.id]

		if hasCount {
			count++
		} else {
			count = 1
		}

		winnersLen := len(c.getWinningNumbers())
		for i := 1; i <= winnersLen; i++ {
			next := c.id + i
			cardCountMap[next] += count
		}

		total += count
	}

	return total
}

func parseInput(input string) (cards []*Card) {
	for _, line := range strings.Split(input, "\n") {
		c := Card{}
		first, second, _ := strings.Cut(line, ":")
		fmt.Sscanf(first, "Card %d", &c.id)

		winStr, numStr, _ := strings.Cut(second, "|")
		c.potentialWinners = cast.ToInts(strings.Fields(winStr))
		c.nums = cast.ToInts(strings.Fields(numStr))

		cards = append(cards, &c)
	}
	return cards
}
