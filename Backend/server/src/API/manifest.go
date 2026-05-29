package server_api

import (
	butcher "aetova/server/src/Butcher"
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

	data, err := os.ReadFile(fmt.Sprintf("%s%s/AppManifest.json", butcher.ToDir, guid)) // TODO :  Get the guid from the request like this /download?guid=xxxxx
	if err != nil {
		http.Error(w, "Can't read the manifest.json", http.StatusNotFound)
		return
	}
	w.Write(data)
}
