package api_download

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func downloadChunk(path string) (data []byte) {
	var client *http.Client = &http.Client{}

	jsonData := `{"path": "` + path + `"}`
	var reader io.Reader = strings.NewReader(jsonData)

	req, err := http.NewRequest("POST", os.Getenv("SERVER_URL")+"/downloader", reader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("api_key", os.Getenv("API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(resp.Status + " " + string(body[:]))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func saveChunk(path string, data []byte) {
	// create new file
	newFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// write the byte in the new file
	if _, err := newFile.Write(data); err != nil {
		log.Fatal(err)
	}
}
