package game

import (
	"fmt"
	"strings"
)

type Game struct {
	Board         [3][3]string
	PlayAgain     bool
	CurrentPlayer string
	Winner        string
	Won           bool
	Moves         int
	Coords        [2]int
}

func NewGame() *Game {
	return &Game{
		Board:         [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}},
		PlayAgain:     true,
		CurrentPlayer: "X",
		Winner:        "",
		Won:           false,
		Moves:         0,
		Coords:        [2]int{0, 0},
	}
}

func (g *Game) NewBoard() {
	var clearBoard [3][3]string

	for i := range clearBoard {
		for j := range clearBoard[i] {
			clearBoard[i][j] = " "
		}
	}
	g.Board = clearBoard
}

func (g *Game) AskPlayAgain() {
	fmt.Println("\nType 'y' to play again:")
	answer := ""
	fmt.Scan(&answer)
	answer = strings.TrimSpace(strings.ToLower(answer))
	g.PlayAgain = (answer == "y")

}

func (g *Game) NextPlayer() {
	if g.CurrentPlayer == "X" {
		g.CurrentPlayer = "O"
	}
	g.CurrentPlayer = "X"
}

// ------- Win Logic ------------------------------------------------------

func (g *Game) CheckForWin() bool {
	if g.checkVertical(g.Board, g.CurrentPlayer) || g.checkHorizontal(g.Board, g.CurrentPlayer) || g.checkDiagonal(g.Board, g.CurrentPlayer) {
		g.Winner = g.CurrentPlayer
		return true
	}
	return false
}

func (g *Game) checkHorizontal(board [3][3]string, CurrentPlayer string) bool {
	matchingCells := 0
	for row := range board {
		for column := range board[row] {
			if board[row][column] == CurrentPlayer {
				matchingCells++
				if matchingCells == 3 {
					return true
				}
			}
		}
		matchingCells = 0
	}
	return false
}

func (g *Game) checkVertical(board [3][3]string, CurrentPlayer string) bool {
	matchingCells := 0
	for column := range board {
		for row := range board[column] {
			if board[row][column] == CurrentPlayer {
				matchingCells++
				if matchingCells == 3 {
					return true
				}
			}
		}
		matchingCells = 0
	}
	return false
}

func (g *Game) checkDiagonal(board [3][3]string, CurrentPlayer string) bool {
	diag1, diag2 := 0, 0
	for i := 0; i < 3; i++ {
		if board[i][i] == CurrentPlayer {
			diag1++
		}
		if board[i][2-i] == CurrentPlayer {
			diag2++
		}
	}
	return diag1 == 3 || diag2 == 3
}

// ------------------------------------------------------------------------

func (g *Game) ApplyMove() bool {
	if g.Board[g.Coords[0]][g.Coords[1]] == " " {
		g.Board[g.Coords[0]][g.Coords[1]] = g.CurrentPlayer
		return true
	} else {
		fmt.Println("\nThis cell is already occupied, try another one")
	}
	return false
}

func (g *Game) DidWon() {
	g.Winner = g.CurrentPlayer
	g.Won = true
}

func (g *Game) IncMove() {
	g.Moves++
}
