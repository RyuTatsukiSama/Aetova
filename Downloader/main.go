package main

import (
	downloader "Aetova/downloader"
	"log"
	"os"
)

func main() {
	err := os.Chdir("WorkingDirectory")

	if err != nil {
		log.Fatal(err)
	}

	err = downloader.Chop("docLogger_v1-1-1.zip")

	if err != nil {
		log.Fatal(err)
	}

	downloader.Assemble()

	if err != nil {
		log.Fatal(err)
	}
}
