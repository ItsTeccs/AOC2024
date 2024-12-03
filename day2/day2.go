package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkLineSafe(fields []string) bool {
	max_dist := 3
	safe := true
	inRange := true
	monotonicIncr := true
	monotonicDecr := true
	for i := 0; i < len(fields)-1 && safe; i++ {
		thisNum, _ := strconv.Atoi(fields[i])
		nextNum, _ := strconv.Atoi(fields[i+1])

		if monotonicIncr {
			monotonicIncr = thisNum < nextNum
		}

		if monotonicDecr {
			monotonicDecr = thisNum > nextNum
		}

		if math.Abs(float64(thisNum-nextNum)) > float64(max_dist) {
			inRange = false
		}

		safe = (monotonicDecr || monotonicIncr) && inRange
	}
	return safe
}

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		if checkLineSafe(strings.Fields(line)) {
			sum++
		}
	}
	fmt.Printf("The number of safe reports in part 1 is: %d\n", sum)
}

func part2(lines []string) {
	sum := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		safe := checkLineSafe(fields)
		for i := 0; i < len(fields) && !safe; i++ {
			// need to assign each time because delete will alter in place
			fields := strings.Fields(line)
			slice := slices.Delete(fields, i, i+1)
			safe = checkLineSafe(slice)
		}
		if safe {
			sum++
		}
	}
	fmt.Printf("The number of safe reports in part 2 is: %d\n", sum)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	lines := strings.Split(string(input), "\n")
	part1(lines)
	part2(lines)
}
