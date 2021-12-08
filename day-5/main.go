package main

import (
	"adventOfCode/day-5/field"
	"adventOfCode/shared"
	"fmt"
	"strconv"
	"strings"
)

type Line = field.Line
type Board = field.Board

func main() {
	input := shared.Read("day-5/input.txt")
	inputLines := strings.Split(input, "\n")

	lines := readLines(inputLines)

	var board1 = Board{
		Board:           make(map[int]map[int]int),
		IgnoreDiagonals: true,
	}

	var board2 = Board{
		Board:           make(map[int]map[int]int),
		IgnoreDiagonals: false,
	}

	for _, line := range lines {
		board1.AddLine(line)
		board2.AddLine(line)
	}

	fmt.Printf("Part 1\n")
	fmt.Printf("    At least 2 intersect: %d\n", board1.AtLeastTwoCounter)

	fmt.Printf("Part 2\n")
	fmt.Printf("    At least 2 intersect: %d\n", board2.AtLeastTwoCounter)
}

func readLines(input []string) []Line {
	var lines []Line

	for _, inputLine := range input {
		split := strings.Split(inputLine, " -> ")
		left := strings.Split(split[0], ",")
		right := strings.Split(split[1], ",")

		x1, _ := strconv.Atoi(left[0])
		y1, _ := strconv.Atoi(left[1])
		x2, _ := strconv.Atoi(right[0])
		y2, _ := strconv.Atoi(right[1])

		line := Line{X1: x1, Y1: y1, X2: x2, Y2: y2}

		lines = append(lines, line)
	}

	return lines
}
