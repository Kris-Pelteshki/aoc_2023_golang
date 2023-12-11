package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Kris-Pelteshki/aoc_2023/util"
	"github.com/emirpasic/gods/queues/arrayqueue"
)

// TODO
// refactor this a bit

type Grid [][]string

const StartSymbol = "S"

var dirs = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
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
	grid, startPos := inputToGrid(input)
	visited := make(map[[2]int]bool)

	queue := arrayqueue.New()
	queue.Enqueue(startPos)
	visited[startPos] = true
	distance := make(map[[2]int]int)
	distance[startPos] = 0
	maxDistance := 0

	for queue.Size() > 0 {
		current, _ := queue.Dequeue()
		currentCoords := current.([2]int)

		for _, neighbor := range getConnectedNeighbors(&grid, currentCoords[0], currentCoords[1]) {
			if !visited[neighbor] {
				queue.Enqueue(neighbor)
				visited[neighbor] = true
				distance[neighbor] = distance[currentCoords] + 1
				if distance[neighbor] > maxDistance {
					maxDistance = distance[neighbor]
				}
			}
		}
	}

	return maxDistance
}

func part2(input string) int {
	return 0
}

func getNeighbors(grid *Grid, x int, y int) (neighbors [][2]int) {
	gridLen := len(*grid)

	for _, dir := range dirs {
		col := x + dir[0]
		row := y + dir[1]

		if col < 0 || col >= gridLen || row < 0 || row >= gridLen || (*grid)[row][col] == "." {
			continue
		}

		neighbors = append(neighbors, [2]int{col, row})
	}
	return neighbors
}

func getConnectedNeighbors(grid *Grid, x int, y int) (connectedNeighbors [][2]int) {
	for _, coords := range getNeighbors(grid, x, y) {
		adjacentPipes := getAdjacentPipes(grid, coords[0], coords[1])

		for _, pipe := range adjacentPipes {
			if pipe[0] == x && pipe[1] == y {
				connectedNeighbors = append(connectedNeighbors, coords)
			}
		}
	}
	return connectedNeighbors
}

func getAdjacentPipes(grid *Grid, x int, y int) (res [2][2]int) {
	symbol := (*grid)[y][x]

	switch symbol {
	case "|":
		res = [2][2]int{
			{x, y - 1},
			{x, y + 1},
		}
	case "-":
		res = [2][2]int{
			{x - 1, y},
			{x + 1, y},
		}
	case "L":
		res = [2][2]int{
			{x + 1, y},
			{x, y - 1},
		}
	case "J":
		res = [2][2]int{
			{x - 1, y},
			{x, y - 1},
		}
	case "7":
		res = [2][2]int{
			{x - 1, y},
			{x, y + 1},
		}
	case "F":
		res = [2][2]int{
			{x + 1, y},
			{x, y + 1},
		}
	}

	return res
}

func inputToGrid(input string) (grid Grid, startPos [2]int) {
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	for i := range grid {
		for j := range grid {
			if grid[i][j] == StartSymbol {
				startPos = [2]int{j, i}
				break
			}
		}
	}
	return grid, startPos
}
