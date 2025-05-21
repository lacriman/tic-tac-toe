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
		board := game.ClearBoard() //empty board
		turn := "X"
		moves := 0
		won := false
		for !won && moves < 9 {
			fmt.Println()
			move := ui.TakeMove() // position for x - [1 0]
			moves++
			// fmt.Printf("Moves: %d", moves)
			for {
				if game.CheckMove(board, move, turn) {
					board[move[0]][move[1]] = turn
					won = game.CheckWin(board, turn)
					turn = game.ChangeTurn(turn)
					break
				} else {
					move = ui.TakeMove() // position for x - [1 0]
					game.CheckMove(board, move, turn)
					continue
				}
			}
			fmt.Println()
			game.PrintBoard(board)
			fmt.Printf("Your move is: %v row, %v column\n", move[0]+1, move[1]+1)

			if won {
				fmt.Printf("\n🎂 Congratulations, %s! You won the game! 🥳 🎊\n", game.ChangeTurn(turn))
			}
		}
		if ui.PlayAgain() != "y" {
			playAgain = false
		}
	}
}
