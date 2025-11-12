package api_download

import (
	"aetova/client/utils"
	"net/http"
	"os"
)

func ReserveSpaceDir(w http.ResponseWriter, manifest utils.ManifestDir, path string) {
	var err error
	err = os.MkdirAll(path+manifest.Name, 0700)
	if err != nil {
		http.Error(w, "Error creating the folder", http.StatusNotFound) // TODO : find better status
	}
}

func reserveSpaceFile(file utils.ManifestFile, path string) {

}
