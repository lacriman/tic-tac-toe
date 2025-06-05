package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/lacriman/tic-tac-toe/game"
)

type CreateGameResponse struct {
	ID            string          `json:"id"`
	Board         [3][3]string    `json:"board"`
	CurrentPlayer string          `json:"currentPlayer"`
	Status        string          `json:"status"`
	Winner        string          `json:"winner,omitempty"`
	LastUpdated   time.Time       `json:"lastUpdated"`
	Message       string          `json:"message,omitempty"`
	Players       game.PlayerInfo `json:"players"`
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
		Status:        newGame.Status,
		LastUpdated:   newGame.LastUpdated,
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
		Status:        gameInstance.Status,
		LastUpdated:   gameInstance.LastUpdated,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func MakeMoveHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	username, exist := sessionStore[cookie.Value]
	if !exist {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var move MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&move); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	gameMux.Lock()
	defer gameMux.Unlock()

	gameInstance, exists := games[id]
	if !exists {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	var playerSymbol string
	var isPlayer bool
	for _, p := range gameInstance.Players {
		if p.Name == username {
			playerSymbol = p.Symbol
			isPlayer = true
			break
		}
	}

	if !isPlayer {
		http.Error(w, "Game not found", http.StatusBadRequest)
		return
	}

	if playerSymbol != gameInstance.CurrentPlayer {
		http.Error(w, "Game not found", http.StatusBadRequest)
		return
	}

	err = gameInstance.MakeMove(move.Row, move.Col)
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
		LastUpdated:   gameInstance.LastUpdated,
		Players:       gameInstance.Players,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	gameID := chi.URLParam(r, "id")
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "name required", http.StatusBadRequest)
		return
	}

	gameMux.Lock()
	defer gameMux.Unlock()

	currentGame, ok := games[gameID]
	if !ok {
		http.Error(w, "game not found", http.StatusBadRequest)
		return
	}

	if len(currentGame.Players) >= 2 {
		http.Error(w, "there are already 2 players", http.StatusBadRequest)
		return
	}

	symbol := "X"
	if len(currentGame.Players) == 1 {
		symbol = "O"
		currentGame.Status = "ready"
	}

	player := game.PlayerInfo{Name: name, Symbol: symbol}
	currentGame.Players = append(currentGame.Players, player)
	currentGame.LastUpdated = time.Now()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Player joined successfully",
		"symbol":   symbol,
		"status":   currentGame.Status,
		"yourName": name,
	})
}
