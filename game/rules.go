package game

import (
	"errors"
)

func CheckMove(board [3][3]string, move [2]int, player string) ([3][3]string, error) {
	if board[move[0]][move[1]] == " " {
		board[move[0]][move[1]] = player
	} else {
		return board, errors.New("cell is already occupied, try again")
	}
	return board, nil
}
