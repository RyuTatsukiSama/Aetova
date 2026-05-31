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

func downloadFile(file utils.ManifestFile, gManifest utils.ManifestGame, path string) {
	bitmap := NewChunkBitmap(file.NbChunks)
	bitmap.StartAutoSave(Ctx, fmt.Sprintf("%s%d/%s.mfs", TargetDl, gManifest.Guid, file.Name), time.Second)
	for part := 0; part < file.NbChunks; part++ {
		chanDownload <- DownloaderData{
			path: fmt.Sprintf("%spart%d_%s.chk", path, part, file.Name),
			writerData: WriterData{
				path:       TargetApp + path + file.Name,
				position:   int64(part * utils.SizeChunk),
				parentFile: file,
				bitmap:     bitmap,
			},
		}
	}
}

func downloadUFile(file utils.ManifestUFile, gManifest utils.ManifestGame, path string) {
	bitmap := NewUChunkBitmap(file.New.NbChunks, file.Chk_changes)
	bitmap.StartAutoSave(Ctx, fmt.Sprintf("%s%d/%s.mfs", TargetDl, gManifest.Guid, file.Name), time.Second)

	for _, part := range file.Chk_changes {
		chanDownload <- DownloaderData{
			path: fmt.Sprintf("%spart%d_%s.chk", path, part, file.Name),
			writerData: WriterData{
				path:       TargetApp + path + file.Name,
				position:   int64(part * utils.SizeChunk),
				parentFile: file.New,
				bitmap:     bitmap,
			},
		}
	}
}

func resumeFile(file utils.ManifestFile, gManifest utils.ManifestGame, path string) error {
	bitmap := NewChunkBitmap(file.NbChunks)
	err := bitmap.LoadFromDisk(fmt.Sprintf("%s%d/%s.mfs", TargetDl, gManifest.Guid, file.Name))
	if err != nil {
		return errors.New("Error 20: " + err.Error())
	}

	bitmap.StartAutoSave(Ctx, fmt.Sprintf("%s%d/%s.mfs", TargetDl, gManifest.Guid, file.Name), time.Second)

	missingChunks := bitmap.MissingChunks()

	atomic.AddInt64(&downloadDone, int64(file.NbChunks)-int64(len(missingChunks)))
	atomic.AddInt64(&writeDone, int64(file.NbChunks)-int64(len(missingChunks)))

	for _, idChunk := range missingChunks {
		chanDownload <- DownloaderData{
			path: fmt.Sprintf("%spart%d_%s.chk", path, idChunk, file.Name),
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

func resumeUFile(file utils.ManifestUFile, gManifest utils.ManifestGame, path string) error {
	bitmap := NewChunkBitmap(file.New.NbChunks)
	err := bitmap.LoadFromDisk(fmt.Sprintf("%s%d/%s.mfs", TargetDl, gManifest.Guid, file.Name))
	if err != nil {
		return errors.New("Error 20: " + err.Error())
	}

	bitmap.StartAutoSave(Ctx, fmt.Sprintf("%s%d/%s.mfs", TargetDl, gManifest.Guid, file.Name), time.Second)

	missingChunks := bitmap.MissingChunks()

	atomic.AddInt64(&downloadDone, int64(file.New.NbChunks)-int64(len(missingChunks)))
	atomic.AddInt64(&writeDone, int64(file.New.NbChunks)-int64(len(missingChunks)))

	for _, idChunk := range missingChunks {
		chanDownload <- DownloaderData{
			path: fmt.Sprintf("%spart%d_%s.chk", path, idChunk, file.Name),
			writerData: WriterData{
				path:       TargetApp + path + file.Name,
				position:   int64(idChunk * utils.SizeChunk),
				parentFile: file.New,
				bitmap:     bitmap,
			},
		}
	}

	return nil
}
