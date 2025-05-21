package game

func ChangeTurn(turn string) string {
	if turn == "X" {
		return "O"
	}
	return "X"
}
