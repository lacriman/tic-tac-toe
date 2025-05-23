package main

import (
	"fmt"

	"github.com/lacriman/tic-tac-toe/game"
)

func main() {
	fmt.Println("Welcome to the Tic Tac Toe!")
	startGame()
}

func startGame() {
	gameInstance := game.NewGame()

	for gameInstance.PlayAgain {
		gameInstance = game.NewGame()
		gameInstance.NewBoard()
		gameInstance.PrintBoard()
		fmt.Println()

		for !gameInstance.Won && gameInstance.Moves < 9 {
			fmt.Println()

			gameInstance.PromptForCoordinate() // position for x - [1 0]
			gameInstance.IncMove()

			for {
				if gameInstance.ApplyMove() {
					gameInstance.DidWon()
					gameInstance.NextPlayer()
					break
				} else {
					gameInstance.PromptForCoordinate() // position for x - [1 0]
					gameInstance.ApplyMove()
					continue
				}
			}
			fmt.Println()
			gameInstance.PrintBoard()
			fmt.Printf("Your move is: %v row, %v column\n", gameInstance.Coords[0]+1, gameInstance.Coords[1]+1)

			if gameInstance.Won {
				fmt.Printf("\n🎂 Congratulations, %s! You won the game! 🥳 🎊\n", gameInstance.Winner)
			}
		}
		gameInstance.AskPlayAgain()
	}
}
