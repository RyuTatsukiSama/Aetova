package server_api

import (
	"aetova/server/utils"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func LaunchAPI() error {
	fmt.Print("Launch API")
	loadHandle()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}

	port, err = utils.FindFreePort(port)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func loadHandle() {
	http.HandleFunc("/health", requireAPIKey(healthHandler))
	http.HandleFunc("/manifest", requireAPIKey(manifestHandler))
	http.HandleFunc("/downloader", requireAPIKey(handlerDownloader))
}
