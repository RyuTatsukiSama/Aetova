package client_api

import (
	"aetova/client/src/API/api_download"
	"aetova/client/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

var (
	cancel context.CancelFunc
	ctx    context.Context
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// HandlePostDownload(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlePostDownload(mxConn MutexConnection) {
	dLog, dCtx, _ := docLogger.NewLogger("Client/download", *docLogger.NewOptionsBuilder().Build(), context.Background())

	start := time.Now()

	api_download.Client = client

	chanError := make(chan error)

	dLog.Log(docLogger.Info, "Get manifest.json")

	done, manifest := grGetManifest(mxConn, chanError, dCtx)
	if !done {
		return
	}

	chanDone := make(chan bool)

	dLog.Log(docLogger.Info, "Reserve space")

	done = grReserveSpace(mxConn, manifest, chanDone, chanError, dCtx)
	if !done {
		return
	}

	dLog.Log(docLogger.Info, "Download")

	done = grDownload(mxConn, manifest, chanDone, chanError, dCtx)
	if !done {
		return
	}

	fmt.Println("Process takes ", time.Since(start))

	dLog.Log(docLogger.Debug, fmt.Sprintln("Process takes ", time.Since(start)))
}

func grGetManifest(mxConn MutexConnection, ce chan error, dCtx context.Context) (bool, utils.ManifestDir) {
	dLog, _, _ := docLogger.NewLogger("", *docLogger.NewOptionsBuilder().Build(), dCtx)

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

func grReserveSpace(mxConn MutexConnection, manifest utils.ManifestDir, cd chan bool, ce chan error, dCtx context.Context) bool {
	dLog, _, _ := docLogger.NewLogger("", *docLogger.NewOptionsBuilder().Build(), dCtx)

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

func grDownload(mxConn MutexConnection, manifest utils.ManifestDir, cd chan bool, ce chan error, dCtx context.Context) bool {
	dLog, _, _ := docLogger.NewLogger("", *docLogger.NewOptionsBuilder().Build(), dCtx)

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
