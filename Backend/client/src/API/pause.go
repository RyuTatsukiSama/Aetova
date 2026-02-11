package client_api

import (
	"net/http"
)

func pauseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		HandlePostPause(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlePostPause(w http.ResponseWriter, r *http.Request) {
	cancel()
}
