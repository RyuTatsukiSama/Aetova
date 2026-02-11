package ws

import (
	mc "aetova/client/src/API/MutexConnection"
	"aetova/client/src/API/api_download"
	"aetova/client/utils"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func handleDownload(mxConn mc.MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")
	start := time.Now()

	chanError := make(chan error)

	dLog.Log(docLogger.Info, "Get manifest.json")

	done, manifest := grGetManifest(mxConn, chanError)
	if !done {
		return
	}

	dLog.Log(docLogger.Info, "Reserve space")

	done = grReserveSpace(mxConn, manifest)
	if !done {
		return
	}

	dLog.Log(docLogger.Info, "Download")

	done = grDownload(mxConn, manifest)
	if !done {
		return
	}

	dLog.Log(docLogger.Debug, fmt.Sprintln("Process takes ", time.Since(start)))
}

func grGetManifest(mxConn mc.MutexConnection, ce chan error) (bool, utils.ManifestDir) {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")

	chanManifest := make(chan utils.ManifestDir)

	go api_download.GetManifest(chanManifest, ce)

	var manifest utils.ManifestDir
	select {
	case <-ctx.Done():
		dLog.Log(docLogger.Info, "Request Canceled")
		return false, utils.ManifestDir{}
	case err := <-ce:
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return false, utils.ManifestDir{}
	case manifest = <-chanManifest:
	}

	data, err := json.Marshal(manifest)
	if err != nil {
		ce <- err
		return false, utils.ManifestDir{}
	}

	err = os.WriteFile("Manifest.json", data, 0777)
	if err != nil {
		ce <- err
		return false, utils.ManifestDir{}
	}

	return true, manifest
}

func grReserveSpace(mxConn mc.MutexConnection, manifest utils.ManifestDir) bool {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")

	api_download.Ctx = ctx

	var wg sync.WaitGroup
	var sErrors []error = make([]error, 0)

	wg.Go(func() {
		api_download.ReserveSpaceDir(manifest, "./downloads/", sErrors)
	})

	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-ctx.Done():
		dLog.Log(docLogger.Info, "Request Canceled")
		return false
	case <-done:
		for i, err := range sErrors {
			dLog.Error(err.Error())
			mxConn.WriteText(err.Error())
			if i == len(sErrors)-1 {
				return false
			}
		}
	}

	return true
}

func grDownload(mxConn mc.MutexConnection, manifest utils.ManifestDir) bool {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")

	api_download.Ctx = ctx

	var wg sync.WaitGroup
	var errors []error

	wg.Go(func() {
		api_download.DownloadGame(manifest, errors, mxConn)
	})

	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-ctx.Done():
		dLog.Log(docLogger.Info, "Request Canceled")
		return false
	case <-done:
		for _, err := range errors {
			dLog.Error(err.Error())
			mxConn.WriteText(err.Error())
			return false
		}
	}

	return true
}
