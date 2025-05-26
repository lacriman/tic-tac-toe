package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/lacriman/tic-tac-toe/game"
)

type CreateGameResponse struct {
	ID            string       `json:"id"`
	Board         [3][3]string `json:"board"`
	CurrentPlayer string       `json:"currentPlayer"`
}

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()
	newGame := game.NewGame()

	gameMux.Lock()
	games[id] = newGame
	gameMux.Unlock()

	response := CreateGameResponse{
		ID:            id,
		Board:         newGame.Board,
		CurrentPlayer: newGame.CurrentPlayer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
