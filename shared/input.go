package shared

import (
	"os"
	"strconv"
	"strings"
)

func Read(path string) string {
	content, error := os.ReadFile("day-1/input.txt")

	if error != nil {
		panic(error)
	}

	return string(content)
}

func ToIntList(input string) []int {
	var list []int

	for _, line := range strings.Split(input, "\n") {
		number, error := strconv.Atoi(line)

		if error != nil {
			panic(error)
		}

		list = append(list, number)
	}

	return list
}
