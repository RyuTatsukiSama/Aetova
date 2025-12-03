package api_download

import (
	"aetova/client/utils"
	"errors"
	"fmt"
)

const (
	MaxWorkers int = 438
)

type WorkerData struct {
	path     string
	fileName string
	part     int
	cd       chan bool
	ce       chan error
}

func Worker(jobs <-chan WorkerData) {
	for j := range jobs {
		downloadChunk(j.path, j.fileName, j.part, j.cd, j.ce)
	}
}

func downloadFile(file utils.ManifestFile, path string) error {

	fmt.Println(file.Name, "start")

	var nbWorkers int = MaxWorkers
	jobs := make(chan WorkerData, file.NbChunks)
	var chanDone = make(chan bool)
	var chanError = make(chan error)

	for wor := 0; wor < nbWorkers; wor++ {
		go Worker(jobs)
	}

	for part := 0; part < file.NbChunks; part++ {
		jobs <- WorkerData{
			path:     path,
			fileName: file.Name,
			part:     part,
			cd:       chanDone,
			ce:       chanError,
		}
	}

	for range file.NbChunks {
		select {
		case <-Ctx.Done():
			return errors.New("request has been canceled")
		case err := <-chanError:
			return err
		case <-chanDone:
		}
	}

	fmt.Println(file.Name, "done")

	return nil
}
