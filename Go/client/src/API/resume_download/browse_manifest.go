package resume_download

import (
	"aetova/client/utils"
	"errors"
)

func BrowseManifest(manifestDir utils.ManifestDir, cm chan utils.ManifestFile, cs chan string, ce chan error) {

	manifestFile, path, err := browseDir(manifestDir)
	if path == "" && err == nil {
		err = errors.New("didn't find the manifest")
		ce <- err
		return
	}
	if err != nil {
		ce <- err
		return
	}

	cm <- manifestFile
	cs <- path
}

func browseDir(manifestDir utils.ManifestDir) (utils.ManifestFile, string, error) {
	for _, subFiles := range manifestDir.SubFiles {
		ok, err := utils.Exists("Manifest_" + subFiles.Name + ".bin")
		if err != nil {
			return utils.ManifestFile{}, "", err
		}

		if ok {
			return subFiles, manifestDir.Name + "/", nil
		}
	}

	for _, subDir := range manifestDir.SubDir {
		subManifest, path, err := browseDir(subDir)
		if err != nil {
			return utils.ManifestFile{}, "", err
		}
		if path != "" {
			return subManifest, manifestDir.Name + "/" + path, nil
		}
	}

	return utils.ManifestFile{}, "", nil
}
