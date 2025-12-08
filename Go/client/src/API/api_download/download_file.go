package api_download

import (
	"aetova/client/utils"
	"errors"
	"fmt"
	"os"
)

const (
	MaxWorkers int = 438
)

type WorkerData struct {
	Path     string
	FileName string
	Part     int
	Cd       chan bool
	Ce       chan error
}

func Worker(jobs <-chan WorkerData, manifestFile *os.File) {
	for j := range jobs {
		downloadChunk(j, manifestFile)
	}
}

func downloadFile(file utils.ManifestFile, path string) error {

	fmt.Println(file.Name, "start")

	manifestFile, err := os.Create("Manifest_" + file.Name + ".bin")
	if err != nil {
		return err
	}

	_, err = manifestFile.Write(make([]byte, file.NbChunks))
	if err != nil {
		return err
	}

	var nbWorkers int = MaxWorkers
	jobs := make(chan WorkerData, file.NbChunks)
	var chanDone = make(chan bool)
	var chanError = make(chan error)

	for wor := 0; wor < nbWorkers; wor++ {
		go Worker(jobs, manifestFile)
	}

	for part := 0; part < file.NbChunks; part++ {
		jobs <- WorkerData{
			Path:     path,
			FileName: file.Name,
			Part:     part,
			Cd:       chanDone,
			Ce:       chanError,
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

	manifestFile.Close()

	err = os.Remove(manifestFile.Name())
	if err != nil {
		return err
	}

	fmt.Println(file.Name, "done")

	return nil
}
