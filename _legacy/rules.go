// // +build ignore

// package game

// func CheckForWin(board [3][3]string, turn string) bool {
// 	return checkVertical(board, turn) || checkHorizontal(board, turn) || checkDiagonal(board, turn)
// }

// func checkHorizontal(board [3][3]string, turn string) bool {
// 	matchingCells := 0
// 	for row := range board {
// 		for column := range board[row] {
// 			if board[row][column] == turn {
// 				matchingCells++
// 				if matchingCells == 3 {
// 					return true
// 				}
// 			}
// 		}
// 		matchingCells = 0
// 	}
// 	return false
// }

// func checkVertical(board [3][3]string, turn string) bool {
// 	matchingCells := 0
// 	for column := range board {
// 		for row := range board[column] {
// 			if board[row][column] == turn {
// 				matchingCells++
// 				if matchingCells == 3 {
// 					return true
// 				}
// 			}
// 		}
// 		matchingCells = 0
// 	}
// 	return false
// }

// func checkDiagonal(board [3][3]string, turn string) bool {
// 	diag1, diag2 := 0, 0
// 	for i := 0; i < 3; i++ {
// 		if board[i][i] == turn {
// 			diag1++
// 		}
// 		if board[i][2-i] == turn {
// 			diag2++
// 		}
// 	}
// 	return diag1 == 3 || diag2 == 3
// }

// // func checkDiagonal(board [3][3]string, turn string) bool {
// // 	winPoint := 0
// // 	for column := 0; column < 3; {
// // 		for row := 0; row < 3; row++ {
// // 			if board[column] == board[row] && board[row][column] == turn {
// // 				winPoint++
// // 				if winPoint == 3 {
// // 					return true
// // 				}
// // 			}

// // 			if column+row == 2 && board[row][column] == turn {
// // 				winPoint++
// // 				if winPoint == 3 {
// // 					return true
// // 				}
// // 			}
// // 		}
// // 		column++
// // 		winPoint = 0
// // 	}
// // 	return false
// // }

// // -------------------------------------------------------------------------------//

// // // horizontal lines
// // for i := 0; i < len(board); {
// // 	if board[i][0] == turn && board[i][1] == turn && board[i][2] == turn {
// // 		return true
// // 	} else {
// // 		i++
// // 	}
// // }
// // // vertical lines
// // for j := 0; j < len(board); {
// // 	if board[0][j] == turn && board[1][j] == turn && board[2][j] == turn {
// // 		return true
// // 	} else {
// // 		j++
// // 	}
// // }

// // -------------------------------------------------------------------------------//

// // 			// vertical lines
// // 	return ((board[0][0] == turn && board[1][0] == turn && board[2][0] == turn) ||
// // 			(board[0][1] == turn && board[1][1] == turn && board[2][1] == turn) ||
// // 			(board[0][2] == turn && board[1][2] == turn && board[2][2] == turn) ||

// // 			// horizontal lines
// // 			(board[0][0] == turn && board[0][1] == turn && board[0][2] == turn) ||
// // 			(board[1][0] == turn && board[1][1] == turn && board[1][2] == turn) ||
// // 			(board[2][0] == turn && board[2][1] == turn && board[2][2] == turn) ||

// // 			// diagonal lines
// // 			(board[0][0] == turn && board[1][1] == turn && board[2][2] == turn) ||
// // 			(board[0][2] == turn && board[1][1] == turn && board[2][0] == turn))
