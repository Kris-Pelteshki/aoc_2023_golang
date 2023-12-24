package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"

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

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(i string) (total int) {
	for _, line := range util.SplitLines(i) {
		first, last := findDigits(line)
		total += cast.ToInt(first + last)
	}
	return total
}

func part2(i string) (total int) {
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range util.SplitLines(i) {
		first, _ := findDigits(frontReplace(line, &digits))
		_, last := findDigits(backReplace(line, &digits))
		total += cast.ToInt(first + last)
	}
	return total
}

func frontReplace(line string, digits *map[string]int) string {
	for i := range line {
		for key := range *digits {
			if len(key)+i > len(line) {
				continue
			}
			if line[i:i+len(key)] == key {
				return line[:i] + strconv.Itoa((*digits)[key]) + line[i+len(key):]
			}
		}
	}
	return line
}

func backReplace(line string, digits *map[string]int) string {
	rline := []rune(line)
	for i := len(rline); i >= 0; i-- {
		for key := range *digits {
			if i-len(key) < 0 {
				continue
			}
			if string(rline[i-len(key):i]) == key {
				return string(rline[:i-len(key)]) + strconv.Itoa((*digits)[key]) + string(rline[i:])
			}
		}
	}
	return line
}

func findDigits(line string) (first, last string) {
	for _, char := range line {
		if unicode.IsDigit(char) {
			last = string(char)
			if len(first) == 0 {
				first = string(char)
			}
		}
	}
	return first, last
}
