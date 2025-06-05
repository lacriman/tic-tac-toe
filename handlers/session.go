package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

var sessionStore = map[string]string{}

func SetUsernameHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Name == "" {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	sessionId := uuid.NewString()
	sessionStore[sessionId] = body.Name

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func GetSessionHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_id")
    if err != nil {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("null"))
        return
    }

    username, exists := sessionStore[cookie.Value]
    if !exists {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("null"))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "name": username,
    })
}