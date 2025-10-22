package main

import (
	"log"
	"net/http"
	"strconv"
)

const (
	// Sample download file provided by Vodafone
	// to test internet
	url = "https://github.com/RyuTatsukiSama/docLogger/releases/download/v1.1.1/docLogger_v1-1-1.zip"

	// Number of Goroutines to spawn for download.
	workers = 5
)

type Part struct {
	Data  []byte
	Index int
}

func main() {
	var size int
	// results := make(chan Part, workers)
	// parts := [workers][]byte{}

	// create the client
	client := &http.Client{}

	// create a new request
	req, err := http.NewRequest("HEAD", url, nil)

	// check if the creation has return an error
	if err != nil {
		log.Fatal(err)
	}

	// send the request and stock the respond
	resp, err := client.Do(req)

	// check if the request return an error
	if err != nil {
		log.Fatal(err)
	}

	// check if the respond contain the content length
	if header, ok := resp.Header["Content-Length"]; ok {
		// get the size of the file we want to download
		fileSize, err := strconv.Atoi(header[0])

		// check if the getting the filesize return an error
		if err != nil {
			log.Fatal("File size could not be determined : ", err)
		}

		size = fileSize / workers

	} else {
		log.Fatal("File size was not provided!")
	}

	println(size)
}
