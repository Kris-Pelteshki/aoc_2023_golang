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

type card = rune
type handType = int
type labelsPriorityMap = map[card]int

type Hand struct {
	cards    []card
	bid      int
	handType handType
}

var HandTypes = map[string]handType{
	"High Card":       0,
	"One Pair":        1,
	"Two Pairs":       2,
	"Three of a Kind": 3,
	"Full house":      4,
	"Four of a kind":  5,
	"Five of a kind":  6,
}

var LabelsPriority = labelsPriorityMap{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

var LabelsPriorityPart2 = labelsPriorityMap{
	'J': -1,
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func (hand *Hand) setHandType() {
	labelCounts := make(map[card]int)
	countsSlice := []int{}

	for _, label := range hand.cards {
		labelCounts[label]++
	}
	for _, count := range labelCounts {
		countsSlice = append(countsSlice, count)
	}

	distinctLabels := len(countsSlice)
	maxGroupLen := slices.Max(countsSlice)

	switch true {
	case distinctLabels == 5:
		hand.handType = HandTypes["High Card"]
	case distinctLabels == 4:
		hand.handType = HandTypes["One Pair"]
	case distinctLabels == 3 && slices.Contains(countsSlice, 2):
		hand.handType = HandTypes["Two Pairs"]
	case distinctLabels == 3:
		hand.handType = HandTypes["Three of a Kind"]
	case distinctLabels == 2 && maxGroupLen == 3:
		hand.handType = HandTypes["Full house"]
	case distinctLabels == 2:
		hand.handType = HandTypes["Four of a kind"]
	case distinctLabels == 1:
		hand.handType = HandTypes["Five of a kind"]
	}
}

func (hand *Hand) setHandTypeWithWildCard(wildcard rune) {
	labelCounts := make(map[card]int)
	countsSlice := []int{}
	maxGroupLen := 0
	var mostCommonLabel rune

	for _, label := range hand.cards {
		labelCounts[label]++
		if labelCounts[label] > maxGroupLen && label != wildcard {
			maxGroupLen = labelCounts[label]
			mostCommonLabel = label
		}
	}

	wildCount, hasWildCard := labelCounts[wildcard]

	if hasWildCard {
		maxGroupLen += wildCount
		labelCounts[mostCommonLabel] += wildCount
		delete(labelCounts, wildcard)
	}

	for _, count := range labelCounts {
		countsSlice = append(countsSlice, count)
	}

	distinctLabels := len(countsSlice)

	switch true {
	case distinctLabels == 5:
		hand.handType = HandTypes["High Card"]
	case distinctLabels == 4:
		hand.handType = HandTypes["One Pair"]
	case distinctLabels == 3 && slices.Contains(countsSlice, 2):
		hand.handType = HandTypes["Two Pairs"]
	case distinctLabels == 3:
		hand.handType = HandTypes["Three of a Kind"]
	case distinctLabels == 2 && maxGroupLen == 3:
		hand.handType = HandTypes["Full house"]
	case distinctLabels == 2:
		hand.handType = HandTypes["Four of a kind"]
	case distinctLabels == 1:
		hand.handType = HandTypes["Five of a kind"]
	default:
		hand.handType = HandTypes["High Card"]
	}
}

func compareHands(hand1, hand2 *Hand, labelMap *labelsPriorityMap) int {
	if hand1.handType > hand2.handType {
		return 1
	} else if hand1.handType < hand2.handType {
		return -1
	} else {
		// compare cards
		for i := 0; i < len(hand1.cards); i++ {
			if (*labelMap)[hand1.cards[i]] > (*labelMap)[hand2.cards[i]] {
				return 1
			}
			if (*labelMap)[hand1.cards[i]] < (*labelMap)[hand2.cards[i]] {
				return -1
			}
		}
		return 0
	}
}

func getTotal(hands []*Hand, priorityMap *labelsPriorityMap) (total int) {
	slices.SortStableFunc(hands, func(hand1, hand2 *Hand) int {
		return compareHands(hand1, hand2, priorityMap)
	})

	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return total
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

func part1(input string) int {
	hands := parseInput(input)

	for i := range hands {
		hands[i].setHandType()
	}

	return getTotal(hands, &LabelsPriority)
}

func part2(input string) (total int) {
	hands := parseInput(input)

	for i := range hands {
		hands[i].setHandTypeWithWildCard('J')
	}

	return getTotal(hands, &LabelsPriorityPart2)
}

func parseInput(input string) (hands []*Hand) {
	for _, line := range util.SplitLines(input) {
		fields := strings.Fields(line)
		hands = append(hands, &Hand{
			cards: []rune(fields[0]),
			bid:   cast.ToInt(fields[1]),
		})
	}
	return hands
}
