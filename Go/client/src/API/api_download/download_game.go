package api_download

import (
	"aetova/client/utils"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

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
	downloadDone int64
	chanWrite    chan WriterData
	writeDone    int64
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
			for data := range chanDownload {
				downloadData(data)
				atomic.AddInt64(&downloadDone, 1)
			}
		})
	}

	// start write workers
	for i := 0; i < MaxWorkers; i++ {
		wg.Go(func() {
			for data := range chanWrite {
				saveChunkAt(data)
				atomic.AddInt64(&writeDone, 1)
			}
		})
	}

	// stock send all the file that need to be download
	err := downloadDir(manifest, "")
	if err != nil {
		se = append(se, err)
	}

	done := make(chan bool, 2)
	go func() {

		wg.Wait()
		dLog.Info("Download done!")
		done <- true
		done <- true
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				dLog.Info(fmt.Sprintf("Download %d/%d | Write %d/%d", atomic.LoadInt64(&downloadDone), NbChunk, atomic.LoadInt64(&writeDone), NbChunk))
				time.Sleep(time.Second)
			}
		}
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
