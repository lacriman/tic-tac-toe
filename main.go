package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lacriman/tic-tac-toe/handlers"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "Welcome to the Tic Tac Toe!")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	r.Route("/api/game", func(r chi.Router) {
		r.Post("/", handlers.CreateGameHandler)
		r.Get("/{id}", handlers.GetGameHandler)
		r.Post("/{id}/move", handlers.MakeMoveHandler)
	})

	fs := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

// ----- Cli -------
// func startGame() {
// 	gameInstance := game.NewGame()

// 	for gameInstance.PlayAgain {
// 		gameInstance = game.NewGame()
// 		gameInstance.NewBoard()
// 		gameInstance.PrintBoard()
// 		fmt.Println()

// 		for !gameInstance.Won && gameInstance.Moves < 9 {
// 			fmt.Println()

// 			gameInstance.PromptForCoordinate() // position for x - [1 0]
// 			gameInstance.IncMove()

// 			for {
// 				if gameInstance.ApplyMove() {
// 					gameInstance.CheckForWin()
// 					gameInstance.NextPlayer()
// 					break
// 				} else {
// 					gameInstance.PromptForCoordinate() // position for x - [1 0]
// 					continue
// 				}
// 			}
// 			fmt.Println()
// 			gameInstance.PrintBoard()
// 			fmt.Printf("Your move is: %v row, %v column\n", gameInstance.Coords[0]+1, gameInstance.Coords[1]+1)

// 			if gameInstance.Won {
// 				fmt.Printf("\n🎂 Congratulations, %s! You won the game! 🥳 🎊\n", gameInstance.Winner)
// 			}
// 		}
// 		gameInstance.AskPlayAgain()
// 	}
// }
