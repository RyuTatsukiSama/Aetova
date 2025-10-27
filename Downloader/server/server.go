package main

import (
	"Aetova/downloader"
	"log"
	"os"
)

func main() {
	var err error = nil

	if err = os.Chdir("wd"); err != nil {
		log.Fatal(err)
	}

	if err = downloader.ChopGame("docLogger_v1-1-1.zip"); err != nil {
		log.Fatal(err)
	}
}
