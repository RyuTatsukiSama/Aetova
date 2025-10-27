package downloader

import (
	"Aetova/util"
	"fmt"
	"os"
)

func ChopGame(zip string) (err error) {
	// unzip the game send by the dev
	var unzipPath string
	if unzipPath, err = Unzip(zip); err != nil {
		return err
	}

	var manifestDir util.ManifestDir
	if manifestDir, err = newChopDir(unzipPath); err != nil {
		return err
	}
	fmt.Println(manifestDir)

	return err
}

func newChopDir(dirPath string) (manifestDir util.ManifestDir, err error) {
	// get the data of the directory
	dir, _err := os.ReadDir(dirPath)
	if _err != nil {
		return util.ManifestDir{}, _err
	}

	fmt.Println(dir)
	manifestDir.Name = dir[0].Name()

	return manifestDir, err
}
