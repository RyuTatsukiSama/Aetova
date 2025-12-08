package api_download

import (
	"aetova/client/utils"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func downloadChunk(job WorkerData, manifestFile *os.File) {

	currentChunkPath := job.Path + "part" + strconv.Itoa(job.Part) + "_" + job.FileName + ".bin"

	data, err := downloadData(currentChunkPath)
	if err != nil {
		job.Ce <- err
		return
	}

	err = saveChunkAt(job.Path+job.FileName, data, int64(job.Part*utils.SizeChunk))
	//err = saveChunk(path+fileName, data)
	if err != nil {
		job.Ce <- err
		return
	}

	done := make([]byte, 1)
	done[0] = 1
	_, err = manifestFile.WriteAt(done, int64(job.Part))
	if err != nil {
		job.Ce <- err
		return
	}

	job.Cd <- true
}

func downloadData(path string) ([]byte, error) {

	jsonData := `{"path": "` + path + `"}`
	var reader io.Reader = strings.NewReader(jsonData)

	req, err := http.NewRequest("POST", os.Getenv("SERVER_URL")+"/downloader", reader)
	if err != nil {
		return make([]byte, 0), errors.New(err.Error() + " downloadData Request")
	}

	req.Header.Add("api_key", os.Getenv("API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := Client.Do(req)
	if err != nil {
		return make([]byte, 0), errors.New(err.Error() + " downloadData Do")
	}
	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return make([]byte, 0), errors.New(err.Error() + " downloadData Response")
		}
		return make([]byte, 0), errors.New(resp.Status + " " + string(body[:]))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return make([]byte, 0), errors.New(err.Error() + " downloadData Read body")
	}

	err = resp.Body.Close()
	if err != nil {
		return make([]byte, 0), errors.New(err.Error() + " downloadData close resp")
	}

	return body, nil
}

func saveChunkAt(path string, data []byte, position int64) error {
	// open file
	newFile, err := os.OpenFile(target+path, os.O_RDWR, 0700)
	if err != nil {
		return errors.New(err.Error() + " saveChunkAt")
	}
	defer newFile.Close()

	// write the byte in the new file
	if _, err := newFile.WriteAt(data, position); err != nil {
		return errors.New(err.Error() + " saveChunkAt")
	}

	return nil
}

func saveChunk(path string, data []byte) error {
	// write the byte in the new file
	if err := os.WriteFile(target+path, data, 0700); err != nil {
		return errors.New(err.Error() + " saveChunk")
	}

	return nil
}
