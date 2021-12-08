package shared

import (
	"os"
	"strconv"
	"strings"
)

func Read(path string) string {
	content, error := os.ReadFile(path)

	if error != nil {
		panic(error)
	}

	return string(content)
}

func ToIntList(input string, convertClosure func(string) (int, error), split string) []int {
	var list []int

	for _, line := range strings.Split(input, split) {
		number, error := convertClosure(line)

		if error != nil {
			continue
		}

		list = append(list, number)
	}

	return list
}

func BinaryStrToInt(binary string) (int, error) {
	integer, error := strconv.ParseInt(binary, 2, 32)

	return int(integer), error
}

func Copy(source []int) []int {
	copySlice := make([]int, len(source))

	copy(copySlice, source)

	return copySlice
}
