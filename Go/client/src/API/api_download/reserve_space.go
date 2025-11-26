package api_download

import (
	"aetova/client/utils"
	"context"
	"os"
)

var (
	Ctx context.Context
)

func ReserveSpaceDir(manifest utils.ManifestDir, path string, ce chan error, cd chan bool) {
	err := os.MkdirAll(path+manifest.Name, 0700)
	if err != nil {
		ce <- err
		return
	}

	chanDone := make(chan bool)
	chanError := make(chan error)

	for _, dir := range manifest.SubDir {
		go ReserveSpaceDir(dir, path+manifest.Name+"/", chanError, chanDone)
	}

	for _, file := range manifest.SubFiles {
		go reserveSpaceFile(file, path+manifest.Name+"/", chanError, chanDone)
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

	cd <- true
}

func reserveSpaceFile(file utils.ManifestFile, path string, ce chan error, cd chan bool) {
	/*reserveBytes := make([]byte, file.Size)
	err := os.WriteFile(path+file.Name, reserveBytes, 0700)
	if err != nil {
		ce <- err
		return
	}*/

	cd <- true
}
