package api_download

import (
	"aetova/client/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func CheckManifest() {
	// if manifest.json exist, just get the data from the file
	// else get it from the server
}

func GetManifest() (manifest utils.ManifestDir, err error) {
	fmt.Print("Download Manifest Start")
	var client *http.Client = &http.Client{}

	req, err := http.NewRequest("GET", os.Getenv("SERVER_URL")+"/manifest", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("api_key", os.Getenv("API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &manifest); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Download Manifest Done")

	return manifest, err
}
