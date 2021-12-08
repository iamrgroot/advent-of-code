package main

import (
	"adventOfCode/shared"
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := shared.Read("day-1/input.txt")
	lines := shared.ToIntList(input, strconv.Atoi)

	var count1 int
	var count2 int
	var previousSum = math.MaxInt

	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			count1++
		}

		// Part 2
		if i < 3 {
			continue
		}

		var sum = lines[i] + lines[i-1] + lines[i-2]

		if sum > previousSum {
			count2++
		}

		previousSum = sum
	}

	fmt.Printf("Part 1: %d \n", count1)
	fmt.Printf("Part 2: %d \n", count2)
}
