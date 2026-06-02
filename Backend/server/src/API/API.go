package server_api

import (
	"aetova/server/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func LaunchAPI() error {
	dLog := docLogger.NewLoggerWithGOpts("Server/LaunchAPI")

	loadHandle()

	port, err := strconv.Atoi("51418") // TODO : Change to avoid hard coded
	if err != nil {
		return err
	}

	port, err = utils.FindFreePort(port)
	if err != nil {
		return err
	}

	dLog.Info(fmt.Sprintf("Server Launch on port %d", port))

	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func loadHandle() {
	http.HandleFunc("/health", requireAPIKey(healthHandler))
	http.HandleFunc("/manifest", requireAPIKey(manifestHandler))
	http.HandleFunc("/downloader", requireAPIKey(handlerDownloader))
	http.HandleFunc("/version", requireAPIKey(versionHandler))
}
