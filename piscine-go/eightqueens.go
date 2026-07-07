package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(0, &board)
}

func solve(col int, board *[8]int) {
	if col == 8 {
		printBoard(board)
		return
	}
	for row := 1; row <= 8; row++ {
		if isSafe(row, col, board) {
			board[col] = row
			solve(col+1, board)
		}
	}
}

func isSafe(row, col int, board *[8]int) bool {
	for c := 0; c < col; c++ {
		if board[c] == row ||
			board[c]-c == row-col ||
			board[c]+c == row+col {
			return false
		}
	}
	return true
}

func printBoard(board *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '0'))
	}
	z01.PrintRune('\n')
}
