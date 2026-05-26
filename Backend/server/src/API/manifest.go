package server_api

import (
	"fmt"
	"net/http"
	"os"
)

func manifestHandler(w http.ResponseWriter, r *http.Request) {
	// get the app id

	switch r.Method {
	case "GET":
		handleGetManifest(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetManifest(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")

	data, err := os.ReadFile(fmt.Sprintf("Mfs_%s.json", guid)) // TODO :  Get the guid from the request like this /download?guid=xxxxx
	if err != nil {
		http.Error(w, "Can't read the manifest.json", http.StatusNotFound)
		return
	}
	w.Write(data)
}
