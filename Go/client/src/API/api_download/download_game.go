package api_download

import (
	"aetova/client/utils"
	"fmt"
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

	if err == nil {
		fmt.Println("Remove Manifest")
		err = os.Remove("Manifest.json")
		if err != nil {
			ce <- err
		}
	}

	cd <- true
}
