package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
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

func part1(input string) (total int) {
	parsed := strings.Split(input, ",")

	for _, s := range parsed {
		total += hash(s)
	}
	return total
}

func part2(input string) (total int) {
	Boxes := make([]Box, 256)
	operations := parseOperations(input)

	for _, op := range operations {
		boxNum := hash(op.Label)
		box := &Boxes[boxNum]
		if op.Operation == "-" {
			box.removeLens(op.Label)
		} else {
			box.addOrReplaceLens(Lens{op.Label, op.FocalLength})
		}
	}

	for i, box := range Boxes {
		for j, lens := range box.Lenses {
			total += (i + 1) * (j + 1) * lens.FocalLength
		}
	}

	return total
}

type Lens struct {
	Label       string
	FocalLength int
}

type Box struct {
	Lenses []Lens
}

type Operation struct {
	Label       string
	Operation   string
	FocalLength int
}

func (box *Box) removeLens(label string) {
	box.Lenses = slices.DeleteFunc(box.Lenses, func(lens Lens) bool {
		return lens.Label == label
	})
}

func (box *Box) addOrReplaceLens(lens Lens) {
	for i, existingLens := range box.Lenses {
		if existingLens.Label == lens.Label {
			box.Lenses[i] = lens
			return
		}
	}
	box.Lenses = append(box.Lenses, lens)
}

func hash(s string) int {
	currentValue := 0
	for _, c := range s {
		currentValue = ((currentValue + int(c)) * 17) % 256
	}
	return currentValue
}

func parseOperations(input string) (operations []Operation) {
	for _, s := range strings.Split(input, ",") {
		operation := Operation{}

		if strings.HasSuffix(s, "-") {
			operation.Operation = "-"
			operation.Label = s[:len(s)-1]
		} else {
			operation.Operation = "="
			operation.Label = s[:len(s)-2]
			operation.FocalLength = cast.ToInt(s[len(s)-1:])
		}
		operations = append(operations, operation)
	}
	return operations
}
