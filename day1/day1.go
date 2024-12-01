package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day1(input []byte) {
	s := string(input)
	lines := strings.Split(s, "\n")
	leftNums := []int{}
	rightNums := []int{}
	sum := 0

	for _, line := range lines {
		lineVals := len((strings.Split(line, " ")))

		leftNum, _ := strconv.Atoi(strings.Split(line, " ")[0])
		leftNums = append(leftNums, leftNum)

		rightNum, _ := strconv.Atoi(strings.Split(line, " ")[lineVals-1])
		rightNums = append(rightNums, rightNum)
		fmt.Printf("Left: %d, right: %d\n", leftNum, rightNum)
	}

	slices.Sort(leftNums)
	slices.Sort(rightNums)

	for i, leftNum := range leftNums {
		rightNum := rightNums[i]
		fmt.Printf("Summing (Left: %d, right: %d)\n", leftNum, rightNum)
		dist := int(math.Abs(float64(leftNum - rightNum)))
		fmt.Printf("Distance: %d\n", dist)
		sum += dist
	}
	fmt.Printf("The result for day 1 is: %d\n", sum)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("An error occured: %v", err)
		return
	}
	day1(input)
}
