package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lacriman/tic-tac-toe/game"
)

func main() {
	fmt.Println("Welcome to the Tic Tac Toe!")
	startGame()
}

func startGame() {
	continueGame := true
	for continueGame {
		move := makeMove()
		game.PrintBoard()
		fmt.Printf("Your move is %v", move)
		break
	}
}

func validateInput(message string) int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)

	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		num, err := strconv.Atoi(input)

		if err != nil || num < 1 || num > 3 {
			fmt.Printf("You have to write a number from 1 to 3, try again: \n")
			continue
		} else {
			return num - 1
		}
	}
}

func makeMove() [2]int {
	column := validateInput("Write a number of a column (1-3): ")
	row := validateInput("Write a number of a row (1-3): ")
	var move [2]int
	move[0] = column
	move[1] = row
	return move
}
