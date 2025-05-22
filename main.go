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
		turn := "X"
		moves := 0
		won := false
		game.PrintBoard(board)
		fmt.Println()
		for !won && moves < 9 {
			fmt.Println()

			move := ui.PromptForCoordinate() // position for x - [1 0]
			moves++
			// fmt.Printf("Moves: %d", moves)
			for {
				if game.ApplyMove(board, move, turn) {
					board[move[0]][move[1]] = turn
					won = game.CheckForWin(board, turn)
					turn = game.NextPlayer(turn)
					break
				} else {
					move = ui.PromptForCoordinate() // position for x - [1 0]
					game.ApplyMove(board, move, turn)
					continue
				}
			}
			fmt.Println()
			game.PrintBoard(board)
			fmt.Printf("Your move is: %v row, %v column\n", move[0]+1, move[1]+1)

			if won {
				fmt.Printf("\n🎂 Congratulations, %s! You won the game! 🥳 🎊\n", game.NextPlayer(turn))
			}
		}
		if ui.PlayAgain() != "y" {
			playAgain = false
		}
	}
}
