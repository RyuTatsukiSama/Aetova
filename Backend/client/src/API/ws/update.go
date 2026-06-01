package ws

import (
	mc "aetova/client/src/API/MutexConnection"
	"aetova/client/src/API/api_download"
	"aetova/client/utils"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

var currentGame utils.ManifestGame

func handleUpdate(mxConn mc.MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("client/handleUpdate")
	start := time.Now()

	chanError := make(chan error)

	ctx, mc.CancelFunc = context.WithCancel(context.Background())

	dLog.Info("Get the app manifest")

	file, err := os.Open(fmt.Sprintf("AppManifest_%d.json", 0)) // TODO: GUID Hard coded here
	if err != nil {
		dLog.Error(err.Error())
		return
	}

	err = utils.FromJson(&currentGame, file)
	if err != nil {
		dLog.Error(err.Error())
		return
	}

	currentGame.Version += 1

	dLog.Info("Get the umanifest json")

	done, manifest := grGetUManifest(mxConn, chanError)
	if !done {
		return
	}

	dLog.Info("Reserve space")

	done = grReserveUSpace(mxConn, manifest)
	if !done {
		return
	}

	dLog.Info("Download")

	done = grUDownload(mxConn, manifest, false)
	if !done {
		return
	}

	dLog.Info("Update the app manifest")

	file, err = os.OpenFile(fmt.Sprintf("AppManifest_%d.json", 0), os.O_WRONLY|os.O_TRUNC, 0777) // TODO: GUID Hard coded here
	if err != nil {
		dLog.Error(err.Error())
		return
	}

	err = utils.ToJson(&currentGame, file)
	if err != nil {
		dLog.Error(err.Error())
		return
	}
	defer file.Close()

	// Clear the current Game just in case
	currentGame = utils.ManifestGame{}

	dLog.Log(docLogger.Debug, fmt.Sprintln("Update took ", time.Since(start)))
}

func grGetUManifest(mxConn mc.MutexConnection, ce chan error) (bool, utils.ManifestUDir) {
	dLog := docLogger.NewLoggerWithGOpts("client/grGetUManifest")

	chanUManifest := make(chan utils.ManifestUDir)

	go api_download.GetUManifest(chanUManifest, ce, currentGame.Guid, currentGame.Version)

	var manifest utils.ManifestUDir
	select {
	case <-ctx.Done():
		dLog.Info(CancelRequestMsg)
		return false, utils.ManifestUDir{}
	case err := <-ce:
		mxConn.WriteText(err.Error())
		dLog.Error(err.Error())
		return false, utils.ManifestUDir{}
	case manifest = <-chanUManifest:
	}

	return true, manifest
}

func grReserveUSpace(mxConn mc.MutexConnection, manifest utils.ManifestUDir) bool {
	dLog := docLogger.NewLoggerWithGOpts(LoggerCD)

	api_download.Ctx = ctx

	var wg sync.WaitGroup
	var sErrors []error = make([]error, 0)

	wg.Go(func() {
		api_download.ReserveSpaceUDir(manifest, "./app/", sErrors)
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

func grUDownload(mxConn mc.MutexConnection, manifest utils.ManifestUDir, isResume bool) bool {
	dLog := docLogger.NewLoggerWithGOpts(LoggerCD)

	api_download.Ctx = ctx

	var wg sync.WaitGroup
	var errors []error

	wg.Go(func() {
		api_download.UpdateGame(manifest, errors, mxConn, isResume, currentGame)
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
