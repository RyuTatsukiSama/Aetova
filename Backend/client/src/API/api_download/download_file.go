package api_download

import (
	"aetova/client/utils"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

const (
	MaxDLWorkers int = 438
	MaxWRWorkers int = 5000
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
	dLog := docLogger.NewLoggerWithGOpts("Client/download")
	dLog.Info(file.Name + " start")

	// file for the resume
	manifestFile, err := os.Create("Manifest_" + file.Name + ".bin")
	if err != nil {
		return err
	}
	defer manifestFile.Close()

	_, err = manifestFile.Write(make([]byte, file.NbChunks))
	if err != nil {
		return err
	}

	// init worker
	var nbWorkers int = MaxDLWorkers
	jobs := make(chan WorkerData, file.NbChunks)
	var chanDone = make(chan bool)
	var chanError = make(chan error)

	// launch worekr
	for wor := 0; wor < nbWorkers; wor++ {
		go Worker(jobs, manifestFile)
	}

	// send jobs to worker
	for part := 0; part < file.NbChunks; part++ {
		jobs <- WorkerData{
			Path:     path,
			FileName: file.Name,
			Part:     part,
			Cd:       chanDone,
			Ce:       chanError,
		}
	}

	// wait for all jobs to be done
	for range file.NbChunks {
		select {
		case <-Ctx.Done():
			return errors.New("request has been canceled")
		case err := <-chanError:
			return err
		case <-chanDone:
		}
	}

	// close and remove manifest of the download
	err = manifestFile.Close()
	if err != nil {
		return err
	}

	err = os.Remove(manifestFile.Name())
	if err != nil {
		return err
	}

	dLog.Info(file.Name + " done")

	return nil
}

func newDownloadFile(file utils.ManifestFile, path string) {
	bitmap := NewChunkBitmap(file.NbChunks)
	bitmap.StartAutoSave(Ctx, targetDl+file.Name+".mfs", 500*time.Millisecond)
	for part := 0; part < file.NbChunks; part++ {
		chanDownload <- DownloaderData{
			path: fmt.Sprintf("%spart%d_%s.bin", path, part, file.Name),
			writerData: WriterData{
				path:       path + file.Name,
				position:   int64(part * utils.SizeChunk),
				parentFile: file,
				bitmap:     bitmap,
			},
		}
	}
}
