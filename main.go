package main

import (
	"fmt"
	"os"

	"github.com/lacriman/tic-tac-toe/game"
	"github.com/lacriman/tic-tac-toe/ui"
)

func main() {
	fmt.Println("Welcome to the Tic Tac Toe!")
	startGame()
}

func startGame() {
	board := game.ClearBoard()
	turn := "X"
	won := false
	continueGame := true

	for continueGame {
		move := ui.TakeMove()
		board, err := game.CheckMove(board, move, turn)
		for err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			move := ui.TakeMove()
			game.CheckMove(board, move, turn)
		}
		turn = game.ChangeTurn(turn)
		fmt.Println()
		game.PrintBoard(board)
		fmt.Printf("Your move is %v", move)

		won = true
		if won {
			break
		}
	}
}
