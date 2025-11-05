package server_api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handlerDownloader(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePostDownloader(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handlePostDownloader(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var path map[string]string
	if err := json.Unmarshal(body, &path); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile("chop/unzip/" + path["path"])
	if err != nil {
		http.Error(w, "Can't read the file "+path["path"], http.StatusNotFound)
		return
	}
	w.Write(data)
}
