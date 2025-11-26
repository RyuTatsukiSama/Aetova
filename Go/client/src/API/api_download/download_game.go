package api_download

import (
	"aetova/client/utils"
)

const (
	target string = "downloads/"
)

func DownloadGame(manifest utils.ManifestDir, cd chan bool, ce chan error) {

	chanDone := make(chan bool)
	chanError := make(chan error)

	go downloadDir(manifest, "", chanDone, chanError)

	select {
	case <-Ctx.Done():
		return
	case err := <-chanError:
		ce <- err
		return
	case <-chanDone:
	}

	cd <- true
}
