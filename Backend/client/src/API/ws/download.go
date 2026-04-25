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

const (
	LoggerCD         string = "Client/download"
	CancelRequestMsg string = "Request Canceled"
)

func handleDownload(mxConn mc.MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts(LoggerCD)
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
	dLog := docLogger.NewLoggerWithGOpts(LoggerCD)

	chanManifest := make(chan utils.ManifestDir)

	go api_download.GetManifest(chanManifest, ce)

	var manifest utils.ManifestDir
	select {
	case <-ctx.Done():
		dLog.Info(CancelRequestMsg)
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

	err = os.WriteFile(fmt.Sprintf(api_download.TargetDl+"Mfs_%d.json", 0), data, 0777) // TODO: When PostgreSQL is here, change 0 by the guid in the manifest
	if err != nil {
		ce <- err
		return false, utils.ManifestDir{}
	}

	return true, manifest
}

func grReserveSpace(mxConn mc.MutexConnection, manifest utils.ManifestDir) bool {
	dLog := docLogger.NewLoggerWithGOpts(LoggerCD)

	api_download.Ctx = ctx

	var wg sync.WaitGroup
	var sErrors []error = make([]error, 0)

	wg.Go(func() {
		api_download.ReserveSpaceDir(manifest, "./app/", sErrors)
	})

	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-ctx.Done():
		dLog.Info(CancelRequestMsg)
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
	dLog := docLogger.NewLoggerWithGOpts(LoggerCD)

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
		dLog.Info(CancelRequestMsg)
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
