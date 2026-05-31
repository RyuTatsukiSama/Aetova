package api_download

import (
	"aetova/client/utils"
	"fmt"
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
			err := resumeFile(file, gManifest, path+manifest.Name+"/")
			if err != nil {
				return err
			}
		} else {
			downloadFile(file, gManifest, path+manifest.Name+"/")
		}
	}

	return nil
}

func downloadUDir(manifest utils.ManifestUDir, gManifest utils.ManifestGame, path string, isResume bool) error {

	// sub dir
	for _, dir := range manifest.SubDir {
		err := downloadUDir(dir, gManifest, path+manifest.Name+"/", isResume)
		if err != nil {
			return err
		}
	}

	// files
	for _, file := range manifest.SubFiles {
		err := handleSubUFile(file, gManifest, path+manifest.Name+"/", isResume)
		if err != nil {
			return err
		}
	}

	return nil
}

func handleSubUFile(file utils.ManifestUFile, gManifest utils.ManifestGame, path string, isResume bool) error {
	switch file.State {
	case utils.Add:
		if isResume {
			err := resumeFile(file.New, gManifest, path)
			if err != nil {
				return err
			}
		} else {
			downloadFile(file.New, gManifest, path)
		}
	case utils.Change:
		if isResume {
			err := resumeUFile(file, gManifest, path)
			if err != nil {
				return err
			}
		} else {
			downloadUFile(file, gManifest, path)
		}
	case utils.Remove:
	default:
		return fmt.Errorf("Error state %d isn't handle", file.State)
	}

	return nil
}
