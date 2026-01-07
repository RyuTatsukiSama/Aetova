package client_api

import (
	"aetova/client/src/API/api_download"
	"aetova/client/utils"
	"fmt"
	"net/http"
	"strconv"
)

var (
	client *http.Client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: api_download.MaxWorkers,
		},
	}
	server_url string = "http://aetova.duckdns.org:15369"
)

func setupRoutes() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/pause", pauseHandler)
	http.HandleFunc("/cancel", cancelHandler)
	http.HandleFunc("/resume", resumeHandler)
	http.HandleFunc("/kill", killHandler)
	http.HandleFunc("/ws", webSocketHandler)
}

func LaunchAPI() (err error) {
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
	fmt.Println("Server start")

	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
