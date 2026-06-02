package server_api

import (
	"aetova/server/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetVersion(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetVersion(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")

	file, err := os.Open(fmt.Sprintf("data/%s/AppManifest.json", guid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var gManifest utils.ManifestGame
	err = utils.FromJson(&gManifest, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(gManifest.Version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
