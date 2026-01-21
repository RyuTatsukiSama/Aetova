package api_download

import (
	"aetova/client/utils"
	"context"
	"os"
	"sync"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

var (
	Ctx context.Context
)

func ReserveSpaceDir(manifest utils.ManifestDir, path string, ce chan error) {
	dLog := docLogger.NewLoggerWithGOpts("Client/reserve_space")

	dLog.Info(path + manifest.Name + " start")

	err := os.MkdirAll(path+manifest.Name, 0700)
	if err != nil {
		ce <- err
		return
	}

	var wg sync.WaitGroup
	var chanError chan error

	for _, dir := range manifest.SubDir {
		wg.Go(func() {
			ReserveSpaceDir(dir, path+manifest.Name+"/", chanError)
			dLog.Debug(path + manifest.Name + "/" + dir.Name + " sub folder done")
		})
	}

	for _, file := range manifest.SubFiles {
		wg.Go(func() {
			reserveSpaceFile(file, path+manifest.Name+"/", chanError)
			dLog.Debug(path + manifest.Name + "/" + file.Name + " sub file done")
		})
	}

	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		dLog.Debug(path + " all sub done")
		close(chanError)
		done <- true
	}()

	select {
	case <-Ctx.Done():
		return
	case <-done:
		for err := range chanError {
			ce <- err
		}
	}

	dLog.Info(path + manifest.Name + " done")
}

func reserveSpaceFile(file utils.ManifestFile, path string, ce chan error) {
	dLog := docLogger.NewLoggerWithGOpts("Client/reserve_space")

	dLog.Info(path + file.Name + " start")
	reserveBytes := make([]byte, file.Size)
	err := os.WriteFile(path+file.Name, reserveBytes, 0700)
	if err != nil {
		ce <- err
		return
	}

	dLog.Info(path + file.Name + " done")
}
