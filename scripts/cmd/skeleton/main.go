package main

import (
	"flag"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/scripts/skeleton"
)

func main() {
	today := time.Now()
	year := flag.Int("year", today.Year(), "AOC year")
	flag.Parse()

	for day := 5; day <= 25; day++ {
		skeleton.Run(day, *year)
	}
}
