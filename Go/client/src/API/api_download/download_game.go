package api_download

import (
	"aetova/client/utils"
)

const (
	target string = "downloads/"
)

func DownloadGame(manifest utils.ManifestDir) error {
	return downloadDir(manifest, "")
}
