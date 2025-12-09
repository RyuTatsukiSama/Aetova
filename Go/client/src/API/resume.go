package client_api

import (
	"aetova/client/src/API/api_download"
	"aetova/client/src/API/resume_download"
	"aetova/client/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		HandlePostResume(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlePostResume(w http.ResponseWriter, r *http.Request) {

	// check if the manifest exist
	ok, err := utils.Exists("Manifest.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Resume Manifest.json not found, no download to resume", http.StatusNotFound)
		return
	}

	// get the data of the manifest
	var manifestDir utils.ManifestDir
	data, err := os.ReadFile("Manifest.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(data, &manifestDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// setup context and client
	api_download.Client = client
	ctx, cancel = context.WithCancel(r.Context())

	chanError := make(chan error)

	// browse the manifest and create the new manifest
	done, file, path := grSearchOnGoingManifest(w, manifestDir, chanError)
	if !done {
		return
	}

	done, newManifest := grCreateNewManifest(w, manifestDir, chanError)
	if !done {
		return
	}

	// download the file that was downloading
	done = grDownloadOnGoingFile(w, file, path, chanError)
	if !done {
		return
	}

	// relaunch download game but with the new manifest
	fmt.Println(newManifest)
}

func grSearchOnGoingManifest(w http.ResponseWriter, manifestDir utils.ManifestDir, ce chan error) (bool, utils.ManifestFile, string) {

	chanManifest := make(chan utils.ManifestFile)
	chanPath := make(chan string)

	go resume_download.SearchOnGoingManifest(manifestDir, chanManifest, chanPath, ce)

	var manifest utils.ManifestFile
	var path string

	select {
	case <-ctx.Done():
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
		return false, utils.ManifestFile{}, ""
	case err := <-ce:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false, utils.ManifestFile{}, ""
	case manifest = <-chanManifest:
		path = <-chanPath
	}

	return true, manifest, path
}

func grCreateNewManifest(w http.ResponseWriter, manifestDir utils.ManifestDir, ce chan error) (bool, utils.ManifestDir) {
	chanManifest := make(chan utils.ManifestDir)

	go resume_download.CreateNewManifest(manifestDir, chanManifest, ce)

	var newManifest utils.ManifestDir

	select {
	case <-ctx.Done():
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
		return false, utils.ManifestDir{}
	case err := <-ce:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false, utils.ManifestDir{}
	case newManifest = <-chanManifest:
	}

	return true, newManifest
}

func grDownloadOnGoingFile(w http.ResponseWriter, manifestFile utils.ManifestFile, path string, ce chan error) bool {
	api_download.Ctx = ctx

	chanDone := make(chan bool)

	go resume_download.ResumeFile(manifestFile, path, chanDone, ce)

	select {
	case <-ctx.Done():
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
		return false
	case err := <-ce:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	case <-chanDone:
	}

	return true
}
