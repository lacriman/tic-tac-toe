// // +build ignore

// package game

// import "fmt"

// func NewBoard() [3][3]string {
// 	var clearBoard [3][3]string

// 	for i := range clearBoard {
// 		for j := range clearBoard[i] {
// 			clearBoard[i][j] = " "
// 		}
// 	}
// 	return clearBoard
// }

// func PrintBoard(board [3][3]string) {
// 	fmt.Println()
// 	for i := 0; i < 3; i++ {
// 		fmt.Printf(" %v | %v | %v \n", board[i][0], board[i][1], board[i][2])
// 		if i < 2 {
// 			fmt.Println("---+---+---")
// 		}
// 	}
// }
