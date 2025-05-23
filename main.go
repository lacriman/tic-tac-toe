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
	gameInstance := game.NewGame()

	for gameInstance.PlayAgain {
		gameInstance.NewBoard()
		game.PrintBoard(gameInstance.Board)
		fmt.Println()

		for !gameInstance.Won && gameInstance.Moves < 9 {
			fmt.Println()

			coords := ui.PromptForCoordinate() // position for x - [1 0]
			gameInstance.IncMove()

			for {
				if game.ApplyMove(gameInstance.Board, gameInstance.Coords, gameInstance.CurrentPlayer) {
					gameInstance.Board[gameInstance.Coords[0]][gameInstance.Coords[1]] = gameInstance.CurrentPlayer
					gameInstance.DidWon()
					gameInstance.CurrentPlayer = game.NextPlayer(gameInstance.CurrentPlayer)
					break
				} else {
					gameInstance.Coords = ui.PromptForCoordinate() // position for x - [1 0]
					game.ApplyMove(gameInstance.Board, coords, gameInstance.CurrentPlayer)
					continue
				}
			}
			fmt.Println()
			game.PrintBoard(gameInstance.Board)
			fmt.Printf("Your move is: %v row, %v column\n", coords[0]+1, coords[1]+1)

			if gameInstance.Won {
				fmt.Printf("\n🎂 Congratulations, %s! You won the game! 🥳 🎊\n", game.NextPlayer(gameInstance.CurrentPlayer))
			}
		}
		gameInstance.AskPlayAgain()
	}
}
