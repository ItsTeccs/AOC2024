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
	}

	slices.Sort(leftNums)
	slices.Sort(rightNums)

	for i, leftNum := range leftNums {
		rightNum := rightNums[i]
		dist := int(math.Abs(float64(leftNum - rightNum)))
		sum += dist
	}
	fmt.Printf("The result for day 1 is: %d", sum)
}

func day2(input []byte) {
	s := string(input)
	lines := strings.Split(s, "\n")
	simScore := 0
	rightOccurs := make(map[int]int)

	leftNums := []int{}

	for _, line := range lines {
		lineVals := len((strings.Split(line, " ")))

		leftNum, _ := strconv.Atoi(strings.Split(line, " ")[0])
		leftNums = append(leftNums, leftNum)

		rightNum, _ := strconv.Atoi(strings.Split(line, " ")[lineVals-1])
		rightOccurs[rightNum] += 1
	}

	for _, num := range leftNums {
		simScore += (num * rightOccurs[num])
	}

	fmt.Printf("The result for day 2 is: %d", simScore)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}
	day1(input)
	day2(input)
}
