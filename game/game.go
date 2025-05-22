package game

type Game struct {
	Board [3][3]string
	PlayAgain bool
	CurrentPlayer string
	Winner string
	Won bool
	Moves int
	Coords [2]int
}	

func NewGame() *Game {
	return &Game{
		Board: [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}},
		PlayAgain: true,
		CurrentPlayer: "X",
		Winner: "",
		Won: false,
		Moves: 0,
		Coords: [2]int{0, 0},
	}
}