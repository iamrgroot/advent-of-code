package main

import (
	"adventOfCode/day-4/bingo"
	"adventOfCode/shared"
	"fmt"
	"strconv"
	"strings"
)

type Board = bingo.Board
type Field = bingo.Field

func main() {
	input := shared.Read("day-4/input.txt")
	lines := strings.Split(input, "\n")

	boards := createBoards(lines)

	numbersDrawn := shared.ToIntList(lines[0], strconv.Atoi, ",")

	firstWinner, firstNumber, lastWinner, lastNumber := getWinners(boards, numbersDrawn)

	fmt.Printf("Part 1\n")
	fmt.Printf("    Winner: %+v\n", firstWinner)
	fmt.Printf("    Winning number: %d\n", firstNumber)
	fmt.Printf("    Multiplied: %d\n\n", firstWinner.SumUnmarked()*firstNumber)

	fmt.Printf("Part 2\n")
	fmt.Printf("    Winner: %+v\n", lastWinner)
	fmt.Printf("    Winning number: %d\n", lastNumber)
	fmt.Printf("    Multiplied: %d\n\n", lastWinner.SumUnmarked()*lastNumber)
}

func createBoards(lines []string) []*Board {
	var boards []*Board

	currentLine := 2
	for currentLine < len(lines) {

		currentBoard := Board{}
		for x := 0; x < bingo.BoardSize; x++ {
			line := lines[currentLine]

			for y, number := range shared.ToIntList(line, strconv.Atoi, " ") {
				currentBoard.Board[x][y] = Field{Value: number, Drawn: false}
			}

			currentLine++
		}

		boards = append(boards, &currentBoard)

		currentLine++
	}

	return boards
}

func getWinners(boards []*Board, numbersDrawn []int) (firstWinner *Board, firstNumber int, lastWinner *Board, lastNumber int) {
	boardsWithBingo := make(map[*Board]bool)

	for _, number := range numbersDrawn {
		for _, board := range boards {
			if boardsWithBingo[board] {
				continue
			}

			bingo := board.Mark(number)

			if bingo {
				boardsWithBingo[board] = true

				if firstWinner == nil {
					firstWinner = board
					firstNumber = number
				}

				lastWinner = board
				lastNumber = number
			}
		}
	}

	return firstWinner, firstNumber, lastWinner, lastNumber
}
