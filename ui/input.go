package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func TakeMove() [2]int {
	column := validateInput("Write a number of a row (1-3): ")
	row := validateInput("Write a number of a column (1-3): ")
	var move [2]int
	move[0] = column
	move[1] = row
	return move
}

func PlayAgain() string {
	fmt.Println("Thanks for your game, do you wanna play again? (Y/N)")
	answer := ""
	fmt.Scan(&answer)
	answer = strings.TrimSpace(strings.ToLower(answer))
	fmt.Println()
	return answer
}
