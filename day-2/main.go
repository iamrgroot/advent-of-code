package main

import (
	"adventOfCode/shared"
	"fmt"
	"strconv"
	"strings"
)

type Submarine struct {
	horizontal int
	depth      int
	aim        int
}

func main() {
	input := shared.Read("day-2/input.txt")
	lines := strings.Split(input, "\n")

	submarine1 := Submarine{horizontal: 0, depth: 0}
	submarine2 := Submarine{horizontal: 0, depth: 0, aim: 0}

	for _, line := range lines {
		split := strings.Split(line, " ")

		direction := split[0]
		amount, error := strconv.Atoi(split[1])

		if error != nil {
			panic(error)
		}

		movePart1(&submarine1, direction, amount)
		movePart2(&submarine2, direction, amount)
	}

	fmt.Printf("Final position part 1: %d, %d\n", submarine1.horizontal, submarine1.depth)
	fmt.Printf("Multiplied part 1: %d\n\n", submarine1.horizontal*submarine1.depth)
	fmt.Printf("Final position part 2: %d, %d\n", submarine2.horizontal, submarine2.depth)
	fmt.Printf("Multiplied part 2: %d\n", submarine2.horizontal*submarine2.depth)
}

func movePart2(submarine *Submarine, direction string, amount int) {
	switch direction {
	case "forward":
		submarine.horizontal += amount
		submarine.depth += submarine.aim * amount
	case "down":
		submarine.aim += amount
	case "up":
		submarine.aim -= amount
	}
}

func movePart1(submarine *Submarine, direction string, amount int) {
	switch direction {
	case "forward":
		submarine.horizontal += amount
	case "down":
		submarine.depth += amount
	case "up":
		submarine.depth -= amount
	}
}
