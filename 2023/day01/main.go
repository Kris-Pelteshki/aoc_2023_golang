package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	// "github.com/Kris-Pelteshki/aoc_2023/cast"
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

// func main() {
// 	var part int
// 	flag.IntVar(&part, "part", 1, "part 1 or 2")
// 	flag.Parse()
// 	fmt.Println("Running part", part)

// 	if part == 10 {
// 		ans := part1(input)
// 		util.CopyToClipboard(fmt.Sprintf("%v", ans))
// 		fmt.Println("Output:", ans)
// 	} else {
// 		ans := part2(input)
// 		util.CopyToClipboard(fmt.Sprintf("%v", ans))
// 		fmt.Println("Output:", ans)
// 	}
// }

// func part1(input string) int {
// 	parsed := parseInput(input)
// 	return util.Sum(parsed)
// }

// func part2(input string) int {
// 	parsed := parseInput(input)
// 	return util.Sum(parsed)
// }

// var searchTerms = map[string]string{
// 	"one":   "1",
// 	"two":   "2",
// 	"three": "3",
// 	"four":  "4",
// 	"five":  "5",
// 	"six":   "6",
// 	"seven": "7",
// 	"eight": "8",
// 	"nine":  "9",
// }

// func parseInput(input string) (ans []int) {
// 	for _, line := range strings.Split(input, "\n") {
// 		for k, v := range searchTerms {
// 			line = strings.ReplaceAll(line, k, v)
// 		}

// 		digits := strings.Map(func(r rune) rune {
// 			if r >= '0' && r <= '9' {
// 				return r
// 			}
// 			return -1
// 		}, line)

// 		first := digits[:1]
// 		last := digits[len(digits)-1:]

// 		num := cast.ToInt(first + last)
// 		ans = append(ans, num)
// 	}
// 	return ans
// }

func main() {
	// lines := strings.Split(input, "\n")

	fmt.Println("part1:", part1(input))
	fmt.Println("part2:", part2(input))
}

func part1(i string) int {
	lines := strings.Split(i, "\n")

	var total int = 0
	for _, line := range lines {
		first, last := findDigits(line)
		ni, _ := strconv.Atoi(first + last)
		total += ni
	}
	return total
}

func part2(i string) int {
	lines := strings.Split(i, "\n")

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

	total := 0
	for _, line := range lines {
		first, _ := findDigits(frontReplace(line, &digits))
		_, last := findDigits(backReplace(line, &digits))
		num, _ := strconv.Atoi(first + last)
		total += num
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

func findDigits(line string) (string, string) {
	var first, last string
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
