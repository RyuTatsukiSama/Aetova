package api_download

import (
	"aetova/client/utils"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

const (
	target string = "downloads/"
)

func DownloadGame(manifest utils.ManifestDir, cd chan bool, ce chan error) {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")

	err := downloadDir(manifest, "")
	if err != nil {
		ce <- err
	}

	if err == nil {
		dLog.Info("Remove Manifest")
		err = os.Remove("Manifest.json")
		if err != nil {
			ce <- err
		}
	}

	cd <- true
}
