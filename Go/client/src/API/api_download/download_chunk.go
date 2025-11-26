package api_download

import (
	"aetova/client/utils"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func downloadChunk(path string, fileName string, part int, cd chan bool, ce chan error) {

	currentChunkPath := path + "part" + strconv.Itoa(part) + "_" + fileName + ".bin"

	fmt.Println("Download file " + currentChunkPath + " Start")

	data, err := downloadData(currentChunkPath)
	if err != nil {
		ce <- err
		return
	}

	err = saveChunk(path+fileName, data, int64(part*utils.SizeChunk))
	if err != nil {
		ce <- err
		return
	}

	fmt.Println("Download file " + currentChunkPath + " Done")

	cd <- true
}

func downloadData(path string) ([]byte, error) {
	var client *http.Client = &http.Client{}

	jsonData := `{"path": "` + path + `"}`
	var reader io.Reader = strings.NewReader(jsonData)

	req, err := http.NewRequest("POST", os.Getenv("SERVER_URL")+"/downloader", reader)
	if err != nil {
		return make([]byte, 0), err
	}

	req.Header.Add("api_key", os.Getenv("API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}
	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return make([]byte, 0), err
		}
		return make([]byte, 0), errors.New(resp.Status + " " + string(body[:]))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return make([]byte, 0), err
	}

	err = resp.Body.Close()
	if err != nil {
		return make([]byte, 0), err
	}

	return body, nil
}

func saveChunk(path string, data []byte, position int64) error {
	// open file
	// newFile, err := os.OpenFile(target+path, os.O_RDWR, 0700)
	// if err != nil {
	// 	return err
	// }
	// defer newFile.Close()

	// write the byte in the new file
	if err := os.WriteFile(target+path, data, 0700); err != nil {
		return err
	}

	fmt.Println("Save Chunk done!")

	return nil
}
