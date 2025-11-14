package api_download

import (
	"aetova/client/utils"
	"net/http"
	"os"
)

func ReserveSpaceDir(w http.ResponseWriter, manifest utils.ManifestDir, path string) {
	err := os.MkdirAll(path+manifest.Name, 0700)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, dir := range manifest.SubDir {
		ReserveSpaceDir(w, dir, path+manifest.Name+"/")
	}

	for _, file := range manifest.SubFiles {
		reserveSpaceFile(w, file, path+manifest.Name+"/")
	}
}

func reserveSpaceFile(w http.ResponseWriter, file utils.ManifestFile, path string) {
	reserveBytes := make([]byte, file.Size)
	err := os.WriteFile(path+file.Name, reserveBytes, 0700)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
