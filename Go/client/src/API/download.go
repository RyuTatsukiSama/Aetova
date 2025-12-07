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
)

var (
	cancel context.CancelFunc
	ctx    context.Context
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		HandlePostDownload(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlePostDownload(w http.ResponseWriter, r *http.Request) { // TODO : check if refactor can be useful, the channels really need to be created here ?
	start := time.Now()

	api_download.Client = client
	ctx, cancel = context.WithCancel(r.Context())

	chanError := make(chan error)

	fmt.Println("Get Manifest")

	done, manifest := grGetManifest(w, chanError)
	if !done {
		return
	}

	chanDone := make(chan bool)

	fmt.Println("Reserve Space")

	done = grReserveSpace(w, manifest, chanDone, chanError)
	if !done {
		return
	}

	fmt.Println("Reserve Download")

	done = grDownload(w, manifest, chanDone, chanError)
	if !done {
		return
	}

	fmt.Println("Process takes ", time.Since(start))
}

func grGetManifest(w http.ResponseWriter, ce chan error) (bool, utils.ManifestDir) {

	chanManifest := make(chan utils.ManifestDir)

	go api_download.GetManifest(chanManifest, ce)

	var manifest utils.ManifestDir
	select {
	case <-ctx.Done():
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
		return false, utils.ManifestDir{}
	case err := <-ce:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false, utils.ManifestDir{}
	case manifest = <-chanManifest:
	}

	data, err := json.Marshal(manifest)
	if err != nil {
		ce <- err
		return false, utils.ManifestDir{}
	}

	err = os.WriteFile("Manifest.json", data, os.ModeAppend)
	if err != nil {
		ce <- err
		return false, utils.ManifestDir{}
	}

	return true, manifest
}

func grReserveSpace(w http.ResponseWriter, manifest utils.ManifestDir, cd chan bool, ce chan error) bool {
	api_download.Ctx = ctx
	go api_download.ReserveSpaceDir(manifest, "./downloads/", ce, cd)
	select {
	case <-ctx.Done():
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
		return false
	case err := <-ce:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	case <-cd:
	}

	return true
}

func grDownload(w http.ResponseWriter, manifest utils.ManifestDir, cd chan bool, ce chan error) bool {
	api_download.Ctx = ctx

	go api_download.DownloadGame(manifest, cd, ce)
	select {
	case <-ctx.Done():
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
		return false
	case err := <-ce:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	case <-cd:
	}

	return true
}
