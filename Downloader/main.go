package main

import (
	"Aetova/downloader"
	"log"
	"os"
)

func main() {
	var err error = nil

	err = os.Chdir("WorkingDirectory")

	downloader.Unzip("docLogger_v1-1-1.zip")

	downloader.ChopGame("uncompressed/docLogger_v1-1-1")

	if err != nil {
		log.Fatal(err)
	}

	// err = downloader.Chop("docLogger_v1-1-1.zip")

	if err != nil {
		log.Fatal(err)
	}

	// downloader.Assemble()

	if err != nil {
		log.Fatal(err)
	}
}
