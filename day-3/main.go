package main

import (
	"adventOfCode/shared"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := shared.Read("day-3/input.txt")
	lines := shared.ToIntList(input, shared.BinaryStrToInt, "\n")

	firstLine := strings.Split(input, "\n")[0]
	digits := len(firstLine)

	gamma := getMostFrequentBits(lines, 0, digits-1)
	espilon := invertedBinaryValue(gamma, digits)

	fmt.Printf("Part 1\n")
	fmt.Printf("    Gamma: %d\n", espilon)
	fmt.Printf("    Espilon: %d\n", gamma)
	fmt.Printf("    Multiplied: %d\n\n", espilon*gamma)

	oxygenOptions := shared.Copy(lines)
	carbonOptions := shared.Copy(lines)

	for i := digits - 1; i >= 0; i-- {
		if len(oxygenOptions) > 1 {
			oxygenMostFrequent := getMostFrequentBits(oxygenOptions, i, i)
			oxygenOptions = filterOptions(oxygenOptions, oxygenMostFrequent, i)
		}
		if len(carbonOptions) > 1 {
			carbonMostFrequent := getMostFrequentBits(carbonOptions, i, i)
			carbonOptions = filterOptions(carbonOptions, carbonMostFrequent^1, i)
		}
		if len(oxygenOptions) <= 1 && len(carbonOptions) <= 1 {
			break
		}
	}

	oxygen := oxygenOptions[0]
	carbon := carbonOptions[0]

	fmt.Printf("Part 2\n")
	fmt.Printf("    Oxygen: %d\n", oxygen)
	fmt.Printf("    Carbon: %d\n", carbon)
	fmt.Printf("    Multiplied: %d\n\n", oxygen*carbon)
}

func getMostFrequentBits(list []int, fromDigit int, toDigit int) int {
	countForEachDigit := countBitsForDigits(list, fromDigit, toDigit)

	return countsToBinaryMostFrequent(countForEachDigit, float32(len(list))/2)
}

func filterOptions(options []int, expected int, digit int) []int {
	filteredOptions := make([]int, 0)

	for _, option := range options {
		if getNthBit(option, digit) == expected {
			filteredOptions = append(filteredOptions, option)
		}
	}

	return filteredOptions
}

func countBitsForDigits(lines []int, fromDigit int, toDigit int) []int {
	counts := make([]int, toDigit-fromDigit+1)

	for _, line := range lines {
		for i := toDigit; i >= fromDigit; i-- {
			counts[i-fromDigit] += getNthBit(line, i)
		}
	}

	return counts
}

func countsToBinaryMostFrequent(counts []int, half float32) int {
	mostFrequent := 0

	for index, count := range counts {
		if float32(count) >= half {
			mostFrequent |= (1 << index)
		}
	}

	return mostFrequent
}

func getNthBit(number int, n int) int {
	return (number >> n) & 1
}

func invertedBinaryValue(binary int, amountOfDigits int) int {
	return binary ^ int(math.Pow(2, float64(amountOfDigits))-1)
}
