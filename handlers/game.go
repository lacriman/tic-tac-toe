package handlers

import (
	"encoding/json"
	"net/http"

	// Universally Unique Identifier
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/lacriman/tic-tac-toe/game"
)

type CreateGameResponse struct {
	ID            string       `json:"id"`
	Board         [3][3]string `json:"board"`
	CurrentPlayer string       `json:"currentPlayer"`
	// YourPlayer    string       `json:"yourPlayer"`
	Status string `json:"status"`
	Winner string `json:"winner"`
}

type MoveRequest struct {
	GameID string `json:"gameID"`
	Row    int    `json:"row"`
	Col    int    `json:"col"`
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

func GetGameHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	gameMux.RLock()
	gameInstance, exists := games[id]
	gameMux.RUnlock()

	if !exists {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	response := CreateGameResponse{
		ID:            id,
		Board:         gameInstance.Board,
		CurrentPlayer: gameInstance.CurrentPlayer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func MakeMoveHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var move MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&move); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	gameMux.Lock()
	gameInstance, exists := games[id]
	if !exists {
		gameMux.Unlock()
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	err := gameInstance.MakeMove(move.Row, move.Col)
	gameMux.Unlock()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, winner := gameInstance.GetStatusAndWinner()
	response := CreateGameResponse{
		ID:            id,
		Board:         gameInstance.Board,
		CurrentPlayer: gameInstance.CurrentPlayer,
		Status:        status,
		Winner:        winner,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
