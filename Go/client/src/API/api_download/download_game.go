package api_download

import (
	"aetova/client/utils"
	"log"
	"os"
)

func DownloadGame(manifest utils.ManifestDir) {

	if err := os.MkdirAll("BuildOranys", 0700); err != nil {
		log.Fatal(err)
	}

	downloadDir(manifest, "")
}
