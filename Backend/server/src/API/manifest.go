package server_api

import (
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
	data, err := os.ReadFile("manifest.json")
	if err != nil {
		http.Error(w, "Can't read the manifest.json", http.StatusNotFound)
		return
	}
	w.Write(data)
}
