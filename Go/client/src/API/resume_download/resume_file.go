package resume_download

import (
	"aetova/client/src/API/api_download"
	"aetova/client/utils"
	"errors"
	"io"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func ResumeFile(file utils.ManifestFile, path string, cd chan bool, ce chan error) {

	// get the data from the manifest
	manifestFile, err := os.OpenFile("Manifest_"+file.Name+".bin", os.O_RDWR, os.ModeAppend)
	if err != nil {
		ce <- err
		return
	}

	data, err := io.ReadAll(manifestFile)
	if err != nil {
		ce <- err
		return
	}

	// put all the chun that need to be download
	var remainingChunk []int
	for indexChunk, b := range data {
		// if the chunk hasn't been download
		if b == 0 {
			remainingChunk = append(remainingChunk, indexChunk)
		}
	}

	err = downloadRemaining(remainingChunk, manifestFile, file, path)
	if err != nil {
		ce <- err
		return
	}

	cd <- true
}

func downloadRemaining(remainingChunk []int, manifestFile *os.File, file utils.ManifestFile, path string) error {
	dLog := docLogger.NewLoggerWithGOpts("Client/resume")

	dLog.Info(file.Name + " resume")

	var nbWorkers int = api_download.MaxWorkers
	jobs := make(chan api_download.WorkerData, len(remainingChunk))
	var chanDone = make(chan bool)
	var chanError = make(chan error)

	for wor := 0; wor < nbWorkers; wor++ {
		go api_download.Worker(jobs, manifestFile)
	}

	for _, chunk := range remainingChunk {
		jobs <- api_download.WorkerData{
			Path:     path,
			FileName: file.Name,
			Part:     chunk,
			Cd:       chanDone,
			Ce:       chanError,
		}
	}

	for range len(remainingChunk) {
		select {
		case <-api_download.Ctx.Done():
			return errors.New("request has been canceled")
		case err := <-chanError:
			return err
		case <-chanDone:
		}
	}

	manifestFile.Close()

	err := os.Remove(manifestFile.Name())
	if err != nil {
		return err
	}

	dLog.Info(file.Name + " done")

	return nil
}
