package api_download

import (
	"aetova/client/utils"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func downloadChunk(job WorkerData, manifestFile *os.File) {

	// currentChunkPath := fmt.Sprintf("%spart%d_%s.bin", job.Path, job.Part, job.FileName)

	// _, err := downloadData(currentChunkPath)
	// if err != nil {
	// 	job.Ce <- err
	// 	return
	// }

	// err = saveChunkAt(job.Path+job.FileName, data, int64(job.Part*utils.SizeChunk))
	// err = saveChunk(path+fileName, data)
	// if err != nil {
	// 	job.Ce <- err
	// 	return
	// }

	// done := make([]byte, 1)
	// done[0] = 1
	// _, err = manifestFile.WriteAt(done, int64(job.Part))
	// if err != nil {
	// 	job.Ce <- err
	// 	return
	// }

	job.Cd <- true
}

func downloadData(data DownloaderData) ([]byte, error) {

	jsonData := fmt.Sprintf(`{"path": "%s"}`, data.path)
	var reader io.Reader = strings.NewReader(jsonData)

	req, err := http.NewRequest("POST", "http://aetova.duckdns.org:15369"+"/downloader", reader)
	if err != nil {
		return make([]byte, 0), errors.New(err.Error() + " downloadData Request")
	}

	req.Header.Add("api_key", "c7e642cc-9928-4248-bd3f-c9588490bb60")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
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

	// give the data to the writer
	data.writerData.data = body
	chanWrite <- data.writerData

	return body, nil
}

var manifestMu sync.Mutex

func saveChunkAt(data WriterData) error {
	manifestMu.Lock()
	err := createManifestDownload(data)
	manifestMu.Unlock()
	if err != nil {
		return err
	}

	// open file
	newFile, err := os.OpenFile(target+data.path, os.O_RDWR, 0700)
	if err != nil {
		return errors.New(err.Error() + " saveChunkAt")
	}
	defer newFile.Close()

	// write the byte in the new file
	if _, err := newFile.WriteAt(data.data, data.position); err != nil {
		return errors.New(err.Error() + " saveChunkAt")
	}

	name := strings.TrimSuffix(data.parentFile.Name, filepath.Ext(data.parentFile.Name))
	manifest_download, err := os.OpenFile("manifest_"+name+".bin", os.O_RDWR, 0700)
	if err != nil {
		return err
	}

	_, err = manifest_download.WriteAt([]byte{1}, data.position/int64(utils.SizeChunk))
	if err != nil {
		return err
	}
	manifest_download.Close()

	manifestMu.Lock()
	err = removeManifestDownload(data)
	manifestMu.Unlock()
	if err != nil {
		return err
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

func createManifestDownload(data WriterData) error {
	name := strings.TrimSuffix(data.parentFile.Name, filepath.Ext(data.parentFile.Name))
	isExist := fileExists("manifest_" + name + ".bin")
	if !isExist {
		f, err := os.Create("manifest_" + name + ".bin")
		if err != nil {
			return err
		}
		defer f.Close()

		bytes := make([]byte, data.parentFile.NbChunks)
		_, err = f.Write(bytes)
		if err != nil {
			return err
		}
	}

	return nil
}

func removeManifestDownload(data WriterData) error {
	name := data.parentFile.Name
	name = "manifest_" + name + ".bin"

	// Check if he has already been removed
	isExist := fileExists(name)
	if !isExist {
		return nil
	}

	bytes, err := os.ReadFile(name)
	if err != nil {
		return err
	}

	for _, b := range bytes {
		if b != 1 {
			return nil
		}
	}

	err = os.Remove(name)
	if err != nil {
		return err
	}

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
