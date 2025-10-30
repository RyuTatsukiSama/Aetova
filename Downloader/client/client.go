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
	if err = dlNass(); err != nil {
		log.Fatal(err)
	}

	// copy into the app folder
}

func dlNass() (err error) {
	// download

	// assemble
	if err = downloader.AssembleGame("manifest.json"); err != nil {
		return err
	}

	return err

	// copy into the app folder
}
