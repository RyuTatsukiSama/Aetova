package api_download

import (
	"aetova/client/utils"
)

func downloadDir(manifest utils.ManifestDir, gManifest utils.ManifestGame, path string, isResume bool) error {

	// sub dir
	for _, dir := range manifest.SubDir {
		err := downloadDir(dir, gManifest, path+manifest.Name+"/", isResume)
		if err != nil {
			return err
		}

	}

	// files
	for _, file := range manifest.SubFiles {
		if isResume {
			resumeDownloadFile(file, gManifest, path+manifest.Name+"/")
		} else {
			newDownloadFile(file, gManifest, path+manifest.Name+"/")
		}
	}

	return nil
}
