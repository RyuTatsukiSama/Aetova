package api_download

import (
	"aetova/client/utils"
)

func downloadDir(manifest utils.ManifestDir, path string, isResume bool) error {

	// sub dir
	for _, dir := range manifest.SubDir {
		err := downloadDir(dir, path+manifest.Name+"/", isResume)
		if err != nil {
			return err
		}

	}

	// files
	for _, file := range manifest.SubFiles {
		if isResume {
			resumeDownloadFile(file, path+manifest.Name+"/")
		} else {
			newDownloadFile(file, path+manifest.Name+"/")
		}
	}

	return nil
}
