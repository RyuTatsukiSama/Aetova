package resume_download

import (
	"aetova/client/utils"
	"encoding/json"
	"errors"
	"os"
	"slices"
)

func CreateNewManifest(manifestDir utils.ManifestDir, cm chan utils.ManifestDir, ce chan error) {
	found, _, err := browseManifest(&manifestDir)
	if !found && err == nil {
		ce <- errors.New("didn't find the manifest")
		return
	}
	if err != nil {
		ce <- err
		return
	}

	data, err := json.Marshal(manifestDir)
	if err != nil {
		ce <- err
	}

	err = os.WriteFile("Manifest.json", data, 0777)
	if err != nil {
		ce <- err
	}

	cm <- manifestDir
}

func browseManifest(manifestDir *utils.ManifestDir) (found bool, needDelete bool, err error) {
	found = false

	var endRange int = 0 // the range we need to delete
	for _, subDir := range manifestDir.SubDir {
		foundRe, needDelete, err := browseManifest(&subDir)
		if err != nil {
			return false, false, err
		}
		if needDelete {
			endRange++
		}
		if foundRe {
			found = true
			break
		}
	}
	manifestDir.SubDir = slices.Delete(manifestDir.SubDir, 0, endRange)

	if !found {
		endRange = 0
		for _, subFile := range manifestDir.SubFiles {
			isExist, err := utils.Exists("Manifest_" + subFile.Name + ".bin")
			if err != nil {
				return false, false, err
			}
			endRange++
			if isExist {
				found = true
				break
			}
		}
		manifestDir.SubFiles = slices.Delete(manifestDir.SubFiles, 0, endRange)
	}

	if len(manifestDir.SubDir) == 0 && len(manifestDir.SubFiles) == 0 {
		return found, true, nil
	}

	return found, false, nil
}
