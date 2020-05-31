package main

type Board struct {
	cells [9][9]int
}

func (b *Board) unassignedCell() (int, int, bool) {
	for rowIndex, row := range b.cells {
		for colIndex, cell := range row {
			if cell == 0 {
				return rowIndex, colIndex, true
			}
		}
	}

	return -1, -1, false
}

func (b *Board) isSafe(x int, y int, cell int) (bool, string, int, int) {

	var i, j int
	for i = 0; i < 9; i++ {
		if b.cells[x][i] == cell && i != y {
			return false, "row", x, -1
		}
	}

	for i = 0; i < 9; i++ {
		if b.cells[i][y] == cell && i != x {
			return false, "column", y, -1
		}
	}

	rowStart := (x / 3) * 3
	colStart := (y / 3) * 3

	for i = rowStart; i < rowStart+3; i++ {
		for j = colStart; j < colStart+3; j++ {
			if b.cells[i][j] == cell && x != i && y != j {
				return false, "matrix", rowStart, colStart
			}
		}
	}

	return true, "None", -1, -1
}

func (b *Board) checkStartingBoard() (bool, string, int, int) {
	for rowIndex, row := range b.cells {
		for colIndex, cell := range row {
			if cell != 0 {
				safe, whereFound, location1, location := b.isSafe(rowIndex, colIndex, cell)
				if safe {
					continue
				} else {
					return false, whereFound, location1, location
				}
			}
		}
	}

	return true, "none", -1, -1
}

func (b *Board) solveSudoku() bool {
	if boardGood, _, _, _ := b.checkStartingBoard(); !boardGood {
		return false
	}

	var i int
	row, col, unassignedCell := b.unassignedCell()
	if !unassignedCell {
		return true
	}

	for i = 1; i < 10; i++ {
		safe, _, _, _ := b.isSafe(row, col, i)
		if safe {
			b.cells[row][col] = i

			if b.solveSudoku() {
				return true
			}

			b.cells[row][col] = 0
		} else {
			continue
		}
	}

	return false
}
