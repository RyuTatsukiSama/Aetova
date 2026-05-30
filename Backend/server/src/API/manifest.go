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
	mType := r.URL.Query().Get("mType")
	guid := r.URL.Query().Get("guid")

	var data []byte
	var err error
	switch mType {
	case "app":
		data, err = os.ReadFile(fmt.Sprintf("%s%s/AppManifest.json", butcher.ToDir, guid))
		if err != nil {
			http.Error(w, "Can't read the manifest.json", http.StatusNotFound)
			return
		}
	case "dl":
		version := r.URL.Query().Get("version")
		data, err = os.ReadFile(fmt.Sprintf("%s%s/Manifest/Mfs_%s.json", butcher.ToDir, guid, version))
		if err != nil {
			http.Error(w, "Can't read the manifest.json", http.StatusNotFound)
			return
		}
	default:
		http.Error(w, "This manifest type doesn't exist", http.StatusNotFound)
		return
	}

	w.Write(data)
}
