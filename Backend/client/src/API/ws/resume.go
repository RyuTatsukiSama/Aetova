package ws

// TODO : All the resume need to be refactor because of the new download system

import (
	mc "aetova/client/src/API/MutexConnection"
	"aetova/client/src/API/api_download"
	"aetova/client/src/API/resume_download"
	"aetova/client/utils"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

const (
	LoggerCR string = "Client/resume"
)

func handlePostResume(mxConn mc.MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts(LoggerCR)

	// restart the context after cancel
	ctx, mc.CancelFunc = context.WithCancel(context.Background())

	mfsPath := fmt.Sprintf("downloads/Mfs_%d.json", 0) // TODO: When PostreSQL is here, get it from the manifest

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
	var manifestDir utils.ManifestDir
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

	// browse the manifest and create the new manifest
	done := grSearchOnGoingManifest(mxConn, manifestDir, chanError)
	if !done {
		return
	}

	/*done, newManifest := grCreateNewManifest(mxConn, manifestDir, chanError)
	if !done {
		return
	}

	// download the file that was downloading
	done = grDownloadOnGoingFile(mxConn, file, path, chanError)
	if !done {
		return
	}*/

	// relaunch download game but with the new manifest
	done = grDownload(mxConn, manifestDir, true)
	if !done {
		return
	}
}

func grSearchOnGoingManifest(mxConn mc.MutexConnection, manifestDir utils.ManifestDir, ce chan error) bool {
	dLog := docLogger.NewLoggerWithGOpts(LoggerCR)

	chanDone := make(chan bool)

	go resume_download.CountChunk(manifestDir, chanDone)

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

func grCreateNewManifest(mxConn mc.MutexConnection, manifestDir utils.ManifestDir, ce chan error) (bool, utils.ManifestDir) {
	dLog := docLogger.NewLoggerWithGOpts(LoggerCR)
	chanManifest := make(chan utils.ManifestDir)

	go resume_download.CreateNewManifest(manifestDir, chanManifest, ce)

	var newManifest utils.ManifestDir

	select {
	case <-ctx.Done():
		dLog.Info(CancelRequestMsg)
		return false, utils.ManifestDir{}
	case err := <-ce:
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return false, utils.ManifestDir{}
	case newManifest = <-chanManifest:
	}

	return true, newManifest
}

func grDownloadOnGoingFile(mxConn mc.MutexConnection, manifestFile utils.ManifestFile, path string, ce chan error) bool {
	dLog := docLogger.NewLoggerWithGOpts(LoggerCR)
	api_download.Ctx = ctx

	chanDone := make(chan bool)

	go resume_download.ResumeFile(manifestFile, path, chanDone, ce)

	select {
	case <-ctx.Done():
		dLog.Info(CancelRequestMsg)
		return false
	case err := <-ce:
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return false
	case <-chanDone:
	}

	return true
}

func grResumeDownload(mxConn mc.MutexConnection, manifest utils.ManifestDir, cd chan bool, ce chan error) bool {
	dLog := docLogger.NewLoggerWithGOpts(LoggerCR)
	api_download.Ctx = ctx

	var wg sync.WaitGroup
	var errors []error

	wg.Go(func() {
		api_download.DownloadGame(manifest, errors, mxConn, true)
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
