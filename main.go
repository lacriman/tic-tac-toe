package main

import (
	"fmt"

	"github.com/lacriman/tic-tac-toe/game"
	"github.com/lacriman/tic-tac-toe/ui"
)

func main() {
	fmt.Println("Welcome to the Tic Tac Toe!")
	startGame()
}

func startGame() {
	board := game.ClearBoard()

	continueGame := true
	for continueGame {
		move := ui.TakeMove()
		fmt.Println()
		game.PrintBoard(board)
		fmt.Printf("Your move is %v", move)
		break
	}
}


