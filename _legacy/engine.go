// +build ignore

package game

import "fmt"

func NextPlayer(turn string) string {
	if turn == "X" {
		return "O"
	}
	return "X"
}

func ApplyMove(board [3][3]string, move [2]int, turn string) bool {
	if board[move[0]][move[1]] == " " {
		board[move[0]][move[1]] = turn
		return true
	} else {
		fmt.Println("\nThis cell is already occupied, try another one")
	}
	return false
}
