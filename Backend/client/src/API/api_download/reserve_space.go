package api_download

import (
	"aetova/client/utils"
	"context"
	"os"
	"sync"
)

var (
	Ctx     context.Context
	NbChunk int = 0
)

func ReserveSpaceDir(manifest utils.ManifestDir, path string, se []error) {

	err := os.MkdirAll(path+manifest.Name, 0700)
	if err != nil {
		se = append(se, err)
		return
	}

	var wg sync.WaitGroup = sync.WaitGroup{}
	var sErrors []error

	for _, dir := range manifest.SubDir {
		wg.Go(func() {
			ReserveSpaceDir(dir, path+manifest.Name+"/", sErrors)
		})
	}

	for _, file := range manifest.SubFiles {
		wg.Go(func() {
			reserveSpaceFile(file, path+manifest.Name+"/", sErrors)
		})
	}

	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-Ctx.Done():
		return
	case <-done:
		for _, err := range sErrors {
			se = append(se, err)
		}
	}
}

func reserveSpaceFile(file utils.ManifestFile, path string, se []error) {

	NbChunk += file.NbChunks
	reserveBytes := make([]byte, file.Size)
	err := os.WriteFile(path+file.Name, reserveBytes, 0700)
	if err != nil {
		se = append(se, err)
		return
	}

}
