package client_api

import (
	"aetova/client/src/API/api_download"
	"aetova/client/utils"
	"context"
	"net/http"
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

func HandlePostDownload(w http.ResponseWriter, r *http.Request) {
	ctx, cancel = context.WithCancel(r.Context())

	chanManifest := make(chan utils.ManifestDir)
	chanError := make(chan error)

	done, manifest := grGetManifest(w, chanManifest, chanError)
	if !done {
		return
	}

	chanDone := make(chan bool)

	done = grReserveSpace(w, manifest, chanDone, chanError)
	if !done {
		return
	}

	/*err := api_download.DownloadGame(manifest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/
}

func grGetManifest(w http.ResponseWriter, cm chan utils.ManifestDir, ce chan error) (bool, utils.ManifestDir) {

	go api_download.GetManifest(cm, ce)

	var manifest utils.ManifestDir
	select {
	case <-ctx.Done():
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
		return false, utils.ManifestDir{}
	case err := <-ce:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false, utils.ManifestDir{}
	case manifest = <-cm:
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
