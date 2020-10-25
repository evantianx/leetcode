// The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
//
//
//
// Given an integer n, return all distinct solutions to the n-queens puzzle.
//
// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.
//
// Example:
//
//
// Input: 4
// Output: [
//  [".Q..",  // Solution 1
//   "...Q",
//   "Q...",
//   "..Q."],
//
//  ["..Q.",  // Solution 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]
// Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.
//
//


// helper function that returns the string representation of the board.
func printBoard(places []int, size int) []string {
	board := make([]string, size)
	for i := range board {
		buf := make([]byte, size)
		for j := 0; j < size; j++ {
			if j == places[i] {
				buf[j] = byte('Q')
			} else {
				buf[j] = byte('.')
			}
		}

		board[i] = string(buf)
	}

	return board
}

func solveNQueens(n int) [][]string {
	boards := [][]string{}

	col := make([]int, n)
	diag := make([]int, n*2)
	diag2 := make([]int, n*2)
	places := make([]int, n)

	// recursive backtracking function
	var search func(y int)
	search = func(y int) {
		if y == n {
			boards = append(boards, printBoard(places, n))
			return
		}

		for x := 0; x < n; x++ {
			if col[x] != 0 || diag[x+y] != 0 || diag2[x-y+n-1] != 0 {
				continue
			}
			col[x], diag[x+y], diag2[x-y+n-1], places[x] = 1, 1, 1, y
			search(y + 1)
			col[x], diag[x+y], diag2[x-y+n-1], places[x] = 0, 0, 0, 0
		}
	}

	search(0)
	return boards
}

