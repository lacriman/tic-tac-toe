package game

import "fmt"

func defaultBoard() [3][3]string {
	var board [3][3]string

	for i := range board {
		for j := range board[i] {
			board[i][j] = "  "
		}
	}
	return board
}

func printBoard() {
	board := defaultBoard()

	for i := 0; i < 3; i++ {
		fmt.Printf("%v | %v | %v\n", board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println("---+----+---")
		}
	}
}
