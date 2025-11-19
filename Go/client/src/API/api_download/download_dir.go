package api_download

import (
	"aetova/client/utils"
	"fmt"
)

func downloadDir(manifest utils.ManifestDir, path string) error {

	fmt.Println("Download dir " + manifest.Name + " Start")

	// sub dir
	for _, dir := range manifest.SubDir {
		err := downloadDir(dir, path+manifest.Name+"/")
		if err != nil {
			return err
		}
	}

	// files
	for _, file := range manifest.SubFiles {
		err := downloadFile(file, path+manifest.Name+"/")
		if err != nil {
			return err
		}
	}

	fmt.Println("Download dir " + manifest.Name + " Done")

	return nil
}
