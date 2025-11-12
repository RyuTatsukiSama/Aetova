package client_api

import (
	"aetova/client/src/API/api_download"
	"net/http"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		HandlePostDownload(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlePostDownload(w http.ResponseWriter, r *http.Request) {
	api_download.GetManifest()
}
