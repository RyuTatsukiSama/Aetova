package server_api

import (
	"fmt"
	"net/http"
	"os"
)

func LaunchAPI() error {
	fmt.Print("Launch API")
	loadHandle()

	port := ":" + os.Getenv("PORT")
	return http.ListenAndServe(port, nil)
}

func loadHandle() {
	http.HandleFunc("/health", requireAPIKey(healthHandler))
	http.HandleFunc("/manifest", requireAPIKey(manifestHandler))
	http.HandleFunc("/downloader", requireAPIKey(handlerDownloader))
}
