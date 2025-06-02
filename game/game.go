package game

import (
	"fmt"
)

type Game struct {
	Board         [3][3]string
	CurrentPlayer string
	Winner        string
	Won           bool
	Moves         int
	Coords        [2]int
	Players       []PlayerInfo
	Status        string
}

type PlayerInfo struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name,omitempty"`
}

func NewGame() *Game {
	return &Game{
		Board:         [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}},
		CurrentPlayer: "X",
		Winner:        "",
		Won:           false,
		Moves:         0,
		Coords:        [2]int{0, 0},
		Players:       []PlayerInfo{},
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
	for i := 0; i < 3; i++ {
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

func (g *Game) MakeMove(row, col int) error {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return fmt.Errorf("invalid coordinates")
	}
	if g.Board[row][col] != " " {
		return fmt.Errorf("cell is already occupied")
	}
	if g.Won || g.Moves == 9 {
		return fmt.Errorf("game is already finished")
	}

	g.Board[row][col] = g.CurrentPlayer
	g.IncMoves()
	g.CheckForWin()

	if g.Won {
		g.Winner = g.CurrentPlayer
		return nil
	}

	if g.Moves == 9 {
		g.Won = true
		// g.Winner stays empty for a draw
		return nil
	}

	g.NextPlayer()
	return nil
}

// ------- Get status and winner ------------------------------------------------------

func (g *Game) GetStatusAndWinner() (string, string) {
	if g.Won {
		return "won", g.Winner
	} else if g.Moves == 9 {
		return "draw", ""
	}
	return "in_progress", ""
}

// ------- Other ------------------------------------------------------

func (g *Game) IncMoves() {
	g.Moves++
}
