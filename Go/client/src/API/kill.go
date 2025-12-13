package client_api

import (
	"fmt"
	"net/http"
	"os"
)

func killHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		HandleGetKill(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleGetKill(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Client close")
	os.Exit(0)
}
