package api_download

import (
	"aetova/client/utils"
	"net/http"
)

const (
	target string = "downloads/"
)

var (
	Client     *http.Client
	nbDownload int = 0
)

func DownloadGame(manifest utils.ManifestDir, cd chan bool, ce chan error) {

	err := downloadDir(manifest, "")
	if err != nil {
		ce <- err
	}

	cd <- true
}
