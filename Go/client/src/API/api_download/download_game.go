package api_download

import (
	"aetova/client/utils"
	"net/http"
	"os"
)

const (
	target string = "downloads/"
)

var (
	Client *http.Client
)

func DownloadGame(manifest utils.ManifestDir, cd chan bool, ce chan error) {

	err := downloadDir(manifest, "")
	if err != nil {
		ce <- err
	}

	err = os.Remove("Manifest.json")
	if err != nil {
		ce <- err
	}

	cd <- true
}
