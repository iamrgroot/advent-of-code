package field

type Row = []int

type Board struct {
	Board             map[int]map[int]int
	IgnoreDiagonals   bool
	AtLeastTwoCounter int
}

type Line struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func (board *Board) AddLine(line Line) {
	if line.X1 == line.X2 {
		board.addVerticalLine(line)
	} else if line.Y1 == line.Y2 {
		board.addHorizontalLine(line)
	} else if !board.IgnoreDiagonals {
		board.addDiagonalLine(line)
	}
}

func (board *Board) addVerticalLine(line Line) {
	smallest, biggest := sorted(line.Y1, line.Y2)

	for i := smallest; i <= biggest; i++ {
		board.increment(line.X1, i)
	}
}

func (board *Board) addHorizontalLine(line Line) {
	smallest, biggest := sorted(line.X1, line.X2)

	for i := smallest; i <= biggest; i++ {
		board.increment(i, line.Y1)
	}
}

func (board *Board) addDiagonalLine(line Line) {
	smallestX, biggestX := sorted(line.X1, line.X2)
	diffence := biggestX - smallestX

	for i := 0; i <= diffence; i++ {
		var x, y int

		if line.X1 > line.X2 {
			x = line.X1 - i
		} else {
			x = line.X1 + i
		}
		if line.Y1 > line.Y2 {
			y = line.Y1 - i
		} else {
			y = line.Y1 + i
		}

		board.increment(x, y)
	}
}

func (board *Board) increment(x, y int) {
	_, xExists := board.Board[x]
	if !xExists {
		board.Board[x] = make(map[int]int)
	}

	_, yExists := board.Board[x][y]
	if !yExists {
		board.Board[x][y] = 0
	}

	board.Board[x][y]++

	if board.Board[x][y] == 2 {
		board.AtLeastTwoCounter++
	}
}

func sorted(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
