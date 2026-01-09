package client_api

import (
	"aetova/client/src/API/api_download"
	"aetova/client/src/API/resume_download"
	"aetova/client/utils"
	"encoding/json"
	"net/http"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// HandlePostResume(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlePostResume(mxConn MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("Client/resume")

	// check if the manifest exist
	ok, err := utils.Exists("Manifest.json")
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
	data, err := os.ReadFile("Manifest.json")
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

	// setup context and client
	api_download.Client = client

	chanError := make(chan error)

	// browse the manifest and create the new manifest
	done, file, path := grSearchOnGoingManifest(mxConn, manifestDir, chanError)
	if !done {
		return
	}

	done, newManifest := grCreateNewManifest(mxConn, manifestDir, chanError)
	if !done {
		return
	}

	// download the file that was downloading
	done = grDownloadOnGoingFile(mxConn, file, path, chanError)
	if !done {
		return
	}

	// relaunch download game but with the new manifest
	done = grResumeDownload(mxConn, newManifest, make(chan bool), chanError)
	if !done {
		return
	}
}

func grSearchOnGoingManifest(mxConn MutexConnection, manifestDir utils.ManifestDir, ce chan error) (bool, utils.ManifestFile, string) {
	dLog := docLogger.NewLoggerWithGOpts("Client/resume")

	chanManifest := make(chan utils.ManifestFile)
	chanPath := make(chan string)

	go resume_download.SearchOnGoingManifest(manifestDir, chanManifest, chanPath, ce)

	var manifest utils.ManifestFile
	var path string

	select {
	case <-ctx.Done():
		dLog.Log(docLogger.Info, "Request Canceled")
		return false, utils.ManifestFile{}, ""
	case err := <-ce:
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return false, utils.ManifestFile{}, ""
	case manifest = <-chanManifest:
		path = <-chanPath
	}

	return true, manifest, path
}

func grCreateNewManifest(mxConn MutexConnection, manifestDir utils.ManifestDir, ce chan error) (bool, utils.ManifestDir) {
	dLog := docLogger.NewLoggerWithGOpts("Client/resume")
	chanManifest := make(chan utils.ManifestDir)

	go resume_download.CreateNewManifest(manifestDir, chanManifest, ce)

	var newManifest utils.ManifestDir

	select {
	case <-ctx.Done():
		dLog.Log(docLogger.Info, "Request Canceled")
		return false, utils.ManifestDir{}
	case err := <-ce:
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return false, utils.ManifestDir{}
	case newManifest = <-chanManifest:
	}

	return true, newManifest
}

func grDownloadOnGoingFile(mxConn MutexConnection, manifestFile utils.ManifestFile, path string, ce chan error) bool {
	dLog := docLogger.NewLoggerWithGOpts("Client/resume")
	api_download.Ctx = ctx

	chanDone := make(chan bool)

	go resume_download.ResumeFile(manifestFile, path, chanDone, ce)

	select {
	case <-ctx.Done():
		dLog.Log(docLogger.Info, "Request Canceled")
		return false
	case err := <-ce:
		mxConn.WriteText(err.Error())
		dLog.Log(docLogger.Error, err.Error())
		return false
	case <-chanDone:
	}

	return true
}

func grResumeDownload(mxConn MutexConnection, manifest utils.ManifestDir, cd chan bool, ce chan error) bool {
	dLog := docLogger.NewLoggerWithGOpts("Client/resume")
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
