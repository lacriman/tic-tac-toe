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
		gameInstance.PrintBoard()
		fmt.Println()

		for !gameInstance.Won && gameInstance.Moves < 9 {
			fmt.Println()

			coords := ui.PromptForCoordinate() // position for x - [1 0]
			gameInstance.IncMove()

			for {
				if gameInstance.ApplyMove() {
					gameInstance.DidWon()
					gameInstance.NextPlayer()
					break
				} else {
					gameInstance.Coords = ui.PromptForCoordinate() // position for x - [1 0]
					gameInstance.ApplyMove()
					continue
				}
			}
			fmt.Println()
			gameInstance.PrintBoard()
			fmt.Printf("Your move is: %v row, %v column\n", coords[0]+1, coords[1]+1)

			if gameInstance.Won {
				fmt.Printf("\n🎂 Congratulations, %s! You won the game! 🥳 🎊\n", gameInstance.Winner)
			}
		}
		gameInstance.AskPlayAgain()
	}
}
