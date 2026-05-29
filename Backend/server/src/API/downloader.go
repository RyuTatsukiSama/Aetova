package server_api

import (
	butcher "aetova/server/src/Butcher"
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

// Send the path of a file, get the content of it
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

	data, err := os.ReadFile(butcher.ToDir + path["path"])
	if err != nil {
		http.Error(w, "Can't read the file "+path["path"], http.StatusNotFound)
		return
	}
	w.Write(data)
}
