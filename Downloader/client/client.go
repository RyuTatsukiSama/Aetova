package main

import (
	"Aetova/downloader"
	"log"
	"os"
)

func main() {
	var err error

	if err = os.Chdir("wd"); err != nil {
		log.Fatal(err)
	}

	// download

	// assemble
	if err = downloader.AssembleGame("manifest.json"); err != nil {
		log.Fatal(err)
	}

	// copy into the app folder
}
