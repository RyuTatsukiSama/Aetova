package client_api

import (
	"net/http"
	"os"
)

var (
	client     *http.Client
	server_url string = os.Getenv("SERVER_URL")
)

func LaunchAPI() (err error) {
	client = &http.Client{}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/download", downloadHandler)

	return http.ListenAndServe(":8090", nil)
}
