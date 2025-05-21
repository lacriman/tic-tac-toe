package game

import "errors"

func ChangeTurn(turn string) string {
	if turn == "X" {
		return "O"
	}
	return "X"
}

func CheckMove(board [3][3]string, move [2]int, turn string) ([3][3]string, error) {
	if board[move[0]][move[1]] == " " {
		board[move[0]][move[1]] = turn
	} else {
		return board, errors.New("cell is already occupied, try again")
	}
	return board, nil
}
