package api_download

import (
	"aetova/client/utils"
	"fmt"
	"log"
	"os"
)

func downloadDir(manifest utils.ManifestDir, path string) {

	fmt.Println("Download dir " + manifest.Name + " Start")

	if err := os.MkdirAll(path+manifest.Name, 0700); err != nil {
		log.Fatal(err)
	}

	// sub dir
	for _, dir := range manifest.SubDir {
		downloadDir(dir, path+manifest.Name+"/")
	}

	// files
	for _, file := range manifest.SubFiles {
		downloadFile(file, path+manifest.Name+"/")
	}

	fmt.Println("Download dir " + manifest.Name + " Done")
}
