package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidateInput(message string) int {
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


// func PlayAgain() string {
// 	fmt.Println("\nType 'y' to play again:")
// 	answer := ""
// 	fmt.Scan(&answer)
// 	answer = strings.TrimSpace(strings.ToLower(answer))
// 	fmt.Println()
// 	return answer
// }
