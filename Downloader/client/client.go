package main

import (
	"Aetova/downloader"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error

	if err = os.Chdir("wd"); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/health", handleHealth)

	http.ListenAndServe(":8090", nil)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {

}

// download and assemble
func dlNass() (err error) {
	// download

	// assemble
	if err = downloader.AssembleGame("manifest.json"); err != nil {
		return err
	}

	return err

	// copy into the app folder
}
