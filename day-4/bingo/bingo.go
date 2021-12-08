package bingo

const (
	BoardSize = 5
)

type Field struct {
	Value int
	Drawn bool
}

type Row = [BoardSize]Field

type Board struct {
	Board [BoardSize]Row
}

func (board *Board) SumUnmarked() int {
	sum := 0
	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			if !board.Board[x][y].Drawn {
				sum += board.Board[x][y].Value
			}
		}
	}
	return sum
}

func (board *Board) Mark(value int) (hasBingo bool) {
	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			if board.Board[x][y].Value == value {
				board.Board[x][y].Drawn = true

				return board.isBingo(x, y)
			}
		}
	}

	return false
}

func (board *Board) isBingo(x int, y int) bool {
	for _, row := range board.rowsToCheck(x, y) {
		if isBingo(&row) {
			return true
		}
	}

	return false
}

func (board *Board) rowsToCheck(x int, y int) [2]Row {
	return [2]Row{
		board.Board[x],
		board.getVerticalRow(y),
	}
}

func (board *Board) getVerticalRow(column int) Row {
	var row Row

	for i := 0; i < BoardSize; i++ {
		row[i] = board.Board[i][column]
	}

	return row
}

func isBingo(row *Row) bool {
	for _, field := range row {
		if !field.Drawn {
			return false
		}
	}

	return true
}
