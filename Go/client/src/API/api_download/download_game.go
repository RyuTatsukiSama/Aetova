package api_download

import (
	"aetova/client/utils"
	"errors"
	"os"
	"sync"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

const (
	target string = "downloads/"
)

type DownloaderData struct {
	path       string
	writerData WriterData
}

type WriterData struct {
	path     string
	data     []byte
	position int64
}

var (
	chanDownload chan DownloaderData
	chanWrite    chan WriterData
)

func DownloadGame(manifest utils.ManifestDir, se []error) {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")

	if NbChunk <= 0 {
		se = append(se, errors.New("error 10 : nbchunk is at 0 or less"))
		return
	}

	chanDownload = make(chan DownloaderData, NbChunk)
	chanWrite = make(chan WriterData, NbChunk)
	var wg sync.WaitGroup = sync.WaitGroup{}
	var sErrors []error = make([]error, 0)

	// start dl workers
	for i := 0; i < MaxWorkers; i++ {
		wg.Go(func() {
			for path := range chanDownload {
				downloadData(path)
			}
		})
	}

	// start write workers
	for i := 0; i < MaxWorkers; i++ {
		wg.Go(func() {
			for data := range chanWrite {
				saveChunkAt(data)
			}
		})
	}

	// stock send all the file that need to be download
	err := downloadDir(manifest, "")
	if err != nil {
		se = append(se, err)
	}

	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-Ctx.Done():
		dLog.Info("Request Cancel")
		return
	case <-done:
		for _, err := range sErrors {
			se = append(se, err)
		}
	}

	if len(se) > 0 {
		dLog.Info("Remove Manifest")
		err = os.Remove("Manifest.json")
		if err != nil {
			se = append(se, err)
		}
	}
}
