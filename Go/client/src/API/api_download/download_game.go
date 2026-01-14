package api_download

import (
	"aetova/client/utils"
	"fmt"
	"os"
)

const (
	target string = "downloads/"
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
