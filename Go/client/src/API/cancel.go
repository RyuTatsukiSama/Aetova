package client_api

import (
	"net/http"
)

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		HandlePostCancel(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlePostCancel(w http.ResponseWriter, r *http.Request) {
	cancel()
	// delete all downloaded file
}
