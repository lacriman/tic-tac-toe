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
	playAgain := true
	for playAgain {
		board := game.NewBoard() //empty board
		currentPlayer := "X"
		moves := 0
		won := false
		game.PrintBoard(board)
		fmt.Println()
		for !won && moves < 9 {
			fmt.Println()

			coords := ui.PromptForCoordinate() // position for x - [1 0]
			moves++
			// fmt.Printf("Moves: %d", moves)
			for {
				if game.ApplyMove(board, coords, currentPlayer) {
					board[coords[0]][coords[1]] = currentPlayer
					won = game.CheckForWin(board, currentPlayer)
					currentPlayer = game.NextPlayer(currentPlayer)
					break
				} else {
					coords = ui.PromptForCoordinate() // position for x - [1 0]
					game.ApplyMove(board, coords, currentPlayer)
					continue
				}
			}
			fmt.Println()
			game.PrintBoard(board)
			fmt.Printf("Your move is: %v row, %v column\n", coords[0]+1, coords[1]+1)

			if won {
				fmt.Printf("\n🎂 Congratulations, %s! You won the game! 🥳 🎊\n", game.NextPlayer(currentPlayer))
			}
		}
		if ui.PlayAgain() != "y" {
			playAgain = false
		}
	}
}
