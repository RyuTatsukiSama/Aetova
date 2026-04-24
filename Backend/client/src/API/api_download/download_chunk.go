package api_download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
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

func downloadData(dData DownloaderData) error {

	jsonData := fmt.Sprintf(`{"path": "%s"}`, dData.path)
	var reader io.Reader = strings.NewReader(jsonData)

	req, err := http.NewRequest("POST", "http://aetova.duckdns.org:15369"+"/downloader", reader)
	if err != nil {
		return errors.New(err.Error() + " downloadData Request")
	}

	req.Header.Add("api_key", "c7e642cc-9928-4248-bd3f-c9588490bb60")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return errors.New(err.Error() + " downloadData Do")
	}
	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.New(err.Error() + " downloadData Response")
		}
		return errors.New(resp.Status + " " + string(body[:]))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(err.Error() + " downloadData Read body")
	}

	err = resp.Body.Close()
	if err != nil {
		return errors.New(err.Error() + " downloadData close resp")
	}

	// give the data to the writer
	dData.writerData.data = body
	chanWrite <- dData.writerData

	return nil
}

func saveChunkAt(data WriterData) error {
	dLog := docLogger.NewLoggerWithGOpts("Client/SaveChunk")

	dLog.Debug("Open file")
	// open file
	newFile, err := os.OpenFile(targetApp+data.path, os.O_RDWR, 0700)
	if err != nil {
		return errors.New(err.Error() + " saveChunkAt")
	}
	defer newFile.Close()

	dLog.Debug("Write file")
	// write the byte in the new file
	if _, err := newFile.WriteAt(data.data, data.position); err != nil {
		return errors.New(err.Error() + " saveChunkAt")
	}

	dLog.Debug("Mark bitmap")
	data.bitmap.Mark(int(data.position))

	return nil
}

func saveChunk(wData WriterData) error {
	// write the byte in the new file
	if err := os.WriteFile(targetApp+wData.path, wData.data, 0700); err != nil {
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
