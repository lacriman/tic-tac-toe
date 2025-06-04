package handlers

import (
	"net/http"

	"github.com/google/uuid"
)

func SetUsernameHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := uuid.NewString()

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

func GetUsernameHandler(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", false
	}
	return cookie.Value, true
}
