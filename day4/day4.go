package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

func getNeighborOffsets() [][]int {
	return [][]int{
		{-1, 1},
		{0, 1},
		{1, 1},
		{-1, 0},
		{1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
}

func checkNeighbors(grid []string, pos position, stack string) bool {
	if len(stack) == 0 {
		return true
	}

	offsets := getNeighborOffsets()
	for _, offset := range offsets {
		x := pos.x + offset[0]
		y := pos.y + offset[1]

		if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
			continue
		}

		if grid[y][x] != stack[0] {
			continue
		}

		return checkNeighbors(grid, position{x, y}, stack[1:])
	}
	return false
}

func part1(grid []string) int {
	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'X' {
				if checkNeighbors(grid, position{x: j, y: i}, "MAS") {
					sum++
				}
			}
		}
	}

	return sum
}

func main() {
	input, err := os.ReadFile("dummyinput.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}
	grid := strings.Split(string(input), "\n")
	fmt.Printf("The result for part 1 of day 4 is: %d!\n", part1(grid))

}
