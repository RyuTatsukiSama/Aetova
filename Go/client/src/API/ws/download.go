package ws

import (
	"aetova/client/src/API/api_download"
	"aetova/client/utils"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func HandleDownload(mxConn MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")

	start := time.Now()

	chanError := make(chan error)

	dLog.Log(docLogger.Info, "Get manifest.json")

	done, manifest := grGetManifest(mxConn, chanError)
	if !done {
		return
	}

	chanDone := make(chan bool)

	dLog.Log(docLogger.Info, "Reserve space")

	done = grReserveSpace(mxConn, manifest, chanDone, chanError)
	if !done {
		return
	}

	dLog.Log(docLogger.Info, "Download")

	done = grDownload(mxConn, manifest, chanDone, chanError)
	if !done {
		return
	}

	dLog.Log(docLogger.Debug, fmt.Sprintln("Process takes ", time.Since(start)))
}

func grGetManifest(mxConn MutexConnection, ce chan error) (bool, utils.ManifestDir) {
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

func grReserveSpace(mxConn MutexConnection, manifest utils.ManifestDir, cd chan bool, ce chan error) bool {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")
	api_download.Ctx = ctx
	go api_download.ReserveSpaceDir(manifest, "./downloads/", ce, cd)
	select {
	case <-ctx.Done():
		dLog.Log(docLogger.Info, "Request Canceled")
		return false
	case err := <-ce:
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return false
	case <-cd:
	}

	return true
}

func grDownload(mxConn MutexConnection, manifest utils.ManifestDir, cd chan bool, ce chan error) bool {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")

	api_download.Ctx = ctx

	go api_download.DownloadGame(manifest, cd, ce)

	select {
	case <-ctx.Done():
		dLog.Log(docLogger.Info, "Request Canceled")
		return false
	case err := <-ce:
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return false
	case <-cd:
	}

	return true
}
