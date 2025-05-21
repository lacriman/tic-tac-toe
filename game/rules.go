package game

func CheckMove(board [3][3]string, move [2]int, player string) ([3][3]string, error) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] != " " {
				board[i][j] = player
				break
			}
		}
	}
	return board, nil
}
