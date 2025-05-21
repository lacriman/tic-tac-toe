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
	playAgain := true
	for playAgain {
		
		board := game.ClearBoard()
		turn := "X"
		moves := 0
		won := false
		for !won && moves < 9 {
			fmt.Println()
			move := ui.TakeMove()
			moves++
			fmt.Printf("Moves: %d", moves)
			
			board, err := game.CheckMove(board, move, turn)
			for err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				move := ui.TakeMove()
				game.CheckMove(board, move, turn)
			}

			turn = game.ChangeTurn(turn)
			fmt.Println()
			game.PrintBoard(board)
			fmt.Printf("Your move is %v\n", move)
		}

		if ui.PlayAgain() != "y" {
			playAgain = false
		}
	}
}
