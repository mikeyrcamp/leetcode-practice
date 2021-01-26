package rooksaresafe

// RooksAreSafe checks if rooks represented by 1 on the chess board can
// one another
// Edge Cases: board cannot be empty. board will be a square
func RooksAreSafe(board [][]int) bool {
	// Iterate through the rows and sum the total of a particular row.
	// If it is greater than one than we have two rooks.
	columnCount := make([]int, len(board[0]))

	for _, row := range board {
		var rowSum int
		for col, v := range row {
			columnCount[col] = columnCount[col] + v
			rowSum += v
			if rowSum > 1 || columnCount[col] > 1 {
				return false
			}
		}
	}

	return true
}
