package api_download

import (
	"aetova/client/utils"
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

func downloadFile(file utils.ManifestFile, path string, cd chan bool, ce chan error) {
	var nbWorkers = file.NbChunks / 100
	jobs := make(chan WorkerData, file.NbChunks)
	chanDone := make(chan bool)
	chanError := make(chan error)

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
			return
		case err := <-chanError:
			ce <- err
			return
		case <-chanDone:
		}
	}

	cd <- true
}
