package api_download

import (
	"aetova/client/utils"
	"fmt"
)

func downloadDir(manifest utils.ManifestDir, path string, cd chan bool, ce chan error) {

	fmt.Println("Download dir " + manifest.Name + " Start")

	chanDone := make(chan bool)
	chanError := make(chan error)

	// sub dir
	for _, dir := range manifest.SubDir {
		go downloadDir(dir, path+manifest.Name+"/", chanDone, chanError)

	}

	// files
	for _, file := range manifest.SubFiles {
		go downloadFile(file, path+manifest.Name+"/", chanDone, chanError)
	}

	for range len(manifest.SubDir) + len(manifest.SubFiles) {
		select {
		case <-Ctx.Done():
			return
		case err := <-chanError:
			ce <- err
			return
		case <-chanDone:
		}
	}

	fmt.Println("Download dir " + manifest.Name + " Done")

	cd <- true
}
