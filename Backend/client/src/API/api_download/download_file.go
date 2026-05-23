package api_download

import (
	"aetova/client/utils"
	"errors"
	"fmt"
	"os"
	"sync/atomic"
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

var dLog docLogger.Logger = *docLogger.NewLogger("download/file", *docLogger.NewOptionsBuilder().Build())

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
	bitmap.StartAutoSave(Ctx, fmt.Sprintf("%s%s%s.mfs", TargetDl, guidFolder, file.Name), time.Second)
	for part := 0; part < file.NbChunks; part++ {
		chanDownload <- DownloaderData{
			path: fmt.Sprintf("%spart%d_%s.bin", path, part, file.Name),
			writerData: WriterData{
				path:       TargetApp + path + file.Name,
				position:   int64(part * utils.SizeChunk),
				parentFile: file,
				bitmap:     bitmap,
			},
		}
	}
}

func resumeDownloadFile(file utils.ManifestFile, path string) error {
	bitmap := NewChunkBitmap(file.NbChunks)
	err := bitmap.LoadFromDisk(fmt.Sprintf("%s%s%s.mfs", TargetDl, guidFolder, file.Name))
	if err != nil {
		return errors.New("Error 20: " + err.Error())
	}

	bitmap.StartAutoSave(Ctx, fmt.Sprintf("%s%s%s.mfs", TargetDl, guidFolder, file.Name), time.Second)

	missingChunks := bitmap.MissingChunks()

	atomic.AddInt64(&downloadDone, int64(file.NbChunks)-int64(len(missingChunks)))
	atomic.AddInt64(&writeDone, int64(file.NbChunks)-int64(len(missingChunks)))

	for _, idChunk := range missingChunks {
		chanDownload <- DownloaderData{
			path: fmt.Sprintf("%spart%d_%s.bin", path, idChunk, file.Name),
			writerData: WriterData{
				path:       TargetApp + path + file.Name,
				position:   int64(idChunk * utils.SizeChunk),
				parentFile: file,
				bitmap:     bitmap,
			},
		}
	}

	return nil
}
