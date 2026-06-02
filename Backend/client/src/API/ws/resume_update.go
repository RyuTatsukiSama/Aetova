package ws

import (
	mc "aetova/client/src/API/MutexConnection"
	"aetova/client/src/API/resume_download"
	"aetova/client/utils"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func handleResumeUpdate(mxConn mc.MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("client/update_resume")

	// restart the context after cancel
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

	mfsPath := fmt.Sprintf("downloads/UMfs_%d.json", 0) // TODO: When PostreSQL is here, get it from the manifest

	// check if the manifest exist
	ok, err := utils.Exists(mfsPath)
	if err != nil {
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return
	}
	if !ok {
		mxConn.WriteText("Resume Manifest.json not found, no download to resume")
		dLog.Log(docLogger.Error, "Resume Manifest.json not found, no download to resume")
		return
	}

	// get the data of the manifest
	var manifestDir utils.ManifestUDir
	data, err := os.ReadFile(mfsPath)
	if err != nil {
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return
	}

	err = json.Unmarshal(data, &manifestDir)
	if err != nil {
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return
	}

	chanError := make(chan error)

	done := grSearchOnGoingUManifest(mxConn, manifestDir, chanError)
	if !done {
		return
	}

	done = grUDownload(mxConn, manifestDir, true)
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
}

func grSearchOnGoingUManifest(mxConn mc.MutexConnection, manifestDir utils.ManifestUDir, ce chan error) bool {
	dLog := docLogger.NewLoggerWithGOpts(LoggerCR)

	chanDone := make(chan bool)

	go resume_download.CountUChunk(manifestDir, chanDone)

	select {
	case <-ctx.Done():
		dLog.Info(CancelRequestMsg)
		return false
	case err := <-ce:
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return false
	case done := <-chanDone:
		return done
	}
}
