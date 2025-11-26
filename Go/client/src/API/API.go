package client_api

import (
	"aetova/client/utils"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var (
	client     *http.Client
	server_url string = os.Getenv("SERVER_URL")
)

func LaunchAPI() (err error) {
	client = &http.Client{}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/pause", pauseHandler)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}

	port, err = utils.FindFreePort(port)
	if err != nil {
		return err
	}

	// TODO : create a file (or other system) for UI to get the port where the client listen

	fmt.Print("Server start")

	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
