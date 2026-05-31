package api_download

import (
	"aetova/client/utils"
	"context"
	"fmt"
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

func ReserveSpaceUDir(manifest utils.ManifestUDir, path string, se []error) {

	if manifest.State == utils.Remove {
		err := os.RemoveAll(path + manifest.Name + "/")
		if err != nil {
			se = append(se, err)
		}
		return
	} else {
		err := os.MkdirAll(path+manifest.Name, 0700)
		if err != nil {
			se = append(se, err)
			return
		}
	}

	var wg sync.WaitGroup = sync.WaitGroup{}
	var sErrors []error

	for _, dir := range manifest.SubDir {
		wg.Go(func() {
			ReserveSpaceUDir(dir, path+manifest.Name+"/", sErrors)
		})
	}

	for _, file := range manifest.SubFiles {
		wg.Go(func() {
			switch file.State {
			case utils.Add:
				reserveSpaceFile(file.New, path+manifest.Name+"/", sErrors)
			case utils.Remove:
				err := os.RemoveAll(fmt.Sprintf("%s%s/%s", path, manifest.Name, file.Name))
				if err != nil {
					sErrors = append(sErrors, err)
				}
			case utils.Change:
				reserveSpaceUFile(file, path+manifest.Name+"/", sErrors)
			default:
				sErrors = append(sErrors, fmt.Errorf("Error state %d isn't handle", file.State))
			}
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

func reserveSpaceUFile(mfsFile utils.ManifestUFile, path string, se []error) {

	NbChunk += len(mfsFile.Chk_changes)

	if mfsFile.New.Size == mfsFile.Old.Size {
		return
	}

	data, err := os.ReadFile(path + mfsFile.Name)
	if err != nil {
		se = append(se, err)
		return
	}

	if mfsFile.New.Size < int64(len(data)) {
		// narrow
		data = data[:mfsFile.New.Size]
	} else {
		// Grow
		data = append(data, make([]byte, int(mfsFile.New.Size)-len(data))...)
	}

	err = os.WriteFile(path+mfsFile.Name, data, 0700)
	if err != nil {
		se = append(se, err)
		return
	}

}
