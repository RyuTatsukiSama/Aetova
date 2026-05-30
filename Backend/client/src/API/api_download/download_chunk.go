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
)

const Server_Url string = "http://localhost:51418"

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

func downloadData(dData DownloaderData, gManifest utils.ManifestGame) error {

	jsonData := fmt.Sprintf(`{"path": "%d/%d/%s"}`, gManifest.Guid, gManifest.Version, dData.path)
	var reader io.Reader = strings.NewReader(jsonData)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/downloader", Server_Url), reader)
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

	// open file
	newFile, err := os.OpenFile(data.path, os.O_RDWR, 0700)
	if err != nil {
		return errors.New(err.Error() + " saveChunkAt")
	}
	defer newFile.Close()

	// write the byte in the new file
	if _, err := newFile.WriteAt(data.data, data.position); err != nil {
		return errors.New(err.Error() + " saveChunkAt")
	}

	data.bitmap.Mark(int(data.position) / utils.SizeChunk)

	return nil
}

func saveChunk(wData WriterData) error {
	// write the byte in the new file
	if err := os.WriteFile(TargetApp+wData.path, wData.data, 0700); err != nil {
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
