package game

import (
	"fmt"
	"strings"

	"github.com/lacriman/tic-tac-toe/ui"
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

func (g *Game) PrintBoard() {
	fmt.Println()
	for i := range 3 {
		fmt.Printf(" %v | %v | %v \n", g.Board[i][0], g.Board[i][1], g.Board[i][2])
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
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
	} else {
		g.CurrentPlayer = "X"
	}
}

// ------- Win Logic ------------------------------------------------------

func (g *Game) CheckForWin() {
	if g.checkVertical() || g.checkHorizontal() || g.checkDiagonal() {
		g.Winner = g.CurrentPlayer
		g.Won = true
		return
	}
	g.Won = false
}

func (g *Game) checkVertical() bool {
	matchingCells := 0
	for column := range g.Board {
		for row := range g.Board[column] {
			if g.Board[row][column] == g.CurrentPlayer {
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

func (g *Game) checkHorizontal() bool {
	matchingCells := 0
	for row := range g.Board {
		for column := range g.Board[row] {
			if g.Board[row][column] == g.CurrentPlayer {
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

func (g *Game) checkDiagonal() bool {
	diag1, diag2 := 0, 0
	for i := range 3 {
		if g.Board[i][i] == g.CurrentPlayer {
			diag1++
		}
		if g.Board[i][2-i] == g.CurrentPlayer {
			diag2++
		}
	}
	return diag1 == 3 || diag2 == 3
}

// ------- Move Logic ------------------------------------------------------

func (g *Game) PromptForCoordinate() {
	g.Coords[0] = ui.ValidateInput("Write a number of a row (1-3): ")
	g.Coords[1] = ui.ValidateInput("Write a number of a column (1-3): ")
}

func (g *Game) ApplyMove() bool {
	if g.Board[g.Coords[0]][g.Coords[1]] == " " {
		g.Board[g.Coords[0]][g.Coords[1]] = g.CurrentPlayer
		return true
	} else {
		fmt.Println("\nThis cell is already occupied, try another one")
		return false
	}
}

// ------- Other ------------------------------------------------------

func (g *Game) IncMove() {
	g.Moves++
}
