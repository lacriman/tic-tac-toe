package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lacriman/tic-tac-toe/handlers"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	r.Route("/api/game", func(r chi.Router) {
		r.Post("/", handlers.CreateGameHandler)
		r.Post("/set-username", handlers.SetUserNameHandler)
		r.Get("/{id}", handlers.GetGameHandler)
		r.Post("/{id}/join", handlers.JoinGameHandler)
		r.Post("/{id}/move", handlers.MakeMoveHandler)
	})

	fs := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
