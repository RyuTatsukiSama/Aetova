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
