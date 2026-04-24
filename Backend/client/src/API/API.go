package client_api

import (
	"aetova/client/src/API/api_download"
	"aetova/client/src/API/ws"
	"aetova/client/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

var (
	client *http.Client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: api_download.MaxDLWorkers,
		},
	}
	server_url string = "http://localhost:51418"
	ctx        context.Context
	cancel     context.CancelFunc
)

func setupRoutes() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ws", ws.WebsocketHandler)
}

func LaunchAPI() (err error) {
	dLog := docLogger.NewLoggerWithGOpts("Client/main")

	setupRoutes()
	port, err := strconv.Atoi("51419") // TODO : Change that to avoid hard code port
	if err != nil {
		return err
	}

	port, err = utils.FindFreePort(port)
	if err != nil {
		return err
	}

	dLog.Log(docLogger.Info, fmt.Sprintf("Client Launch on port %d!", port))

	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
