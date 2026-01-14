package client_api

import (
	"aetova/client/src/API/api_download"
	"aetova/client/src/API/ws"
	"aetova/client/utils"
	"context"
	"net/http"
	"strconv"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

var (
	client *http.Client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: api_download.MaxWorkers,
		},
	}
	server_url string = "http://aetova.duckdns.org:15369"
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

	// TODO : create a file (or other system) for UI to get the port where the client listen, or have an endpoint "test" and the UI do the same thing has "FindFreePort"

	dLog.Log(docLogger.Info, "Client Launch!")

	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
