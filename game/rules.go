package game

func CheckWin(board [3][3]string, turn string) bool {
			// vertical lines
	return ((board[0][0] == turn && board[1][0] == turn && board[2][0] == turn) ||
			(board[0][1] == turn && board[1][1] == turn && board[2][1] == turn) ||
			(board[0][2] == turn && board[1][2] == turn && board[2][2] == turn) ||

			// horizontal lines
			(board[0][0] == turn && board[0][1] == turn && board[0][2] == turn) ||
			(board[1][0] == turn && board[1][1] == turn && board[1][2] == turn) ||
			(board[2][0] == turn && board[2][1] == turn && board[2][2] == turn) ||

			// slash lines
			(board[0][0] == turn && board[1][1] == turn && board[2][2] == turn) ||
			(board[0][2] == turn && board[1][1] == turn && board[2][0] == turn))
}
