package api_download

import (
	mc "aetova/client/src/API/MutexConnection"
	"aetova/client/utils"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

const (
	targetApp string = "app/"
	targetDl  string = "downloads/"
)

type DownloaderData struct {
	path       string
	writerData WriterData
}

type WriterData struct {
	path       string
	data       []byte
	position   int64
	parentFile utils.ManifestFile
	bitmap     *ChunkBitmap
}

type MonitoringData struct {
	DlPrc   float64
	DlSpeed float64
	WrPrc   float64
	WrSpeed float64
}

var (
	chanDownload chan DownloaderData
	downloadDone int64
	chanWrite    chan WriterData
	writeDone    int64
)

func DownloadGame(manifest utils.ManifestDir, se []error, mxConn mc.MutexConnection) {
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
	for i := 0; i < MaxDLWorkers; i++ {
		wg.Go(func() {
			for data := range chanDownload {
				select {
				case <-Ctx.Done():
					dLog.Info("Downloader" + strconv.Itoa(i) + " Request Cancel")
					return
				default:
					err := downloadData(data)
					if err != nil {
						dLog.Error("Error 14: " + err.Error())
					} else {
						atomic.AddInt64(&downloadDone, 1)
					}
				}
			}
		})
	}

	// start write workers
	for i := 0; i < MaxWRWorkers; i++ {
		wg.Go(func() {
			for data := range chanWrite {
				select {
				case <-Ctx.Done():
					dLog.Info("Writer" + strconv.Itoa(i) + " Request Cancel")
					return
				default:
					err := saveChunkAt(data)
					if err != nil {
						dLog.Error("Error 15: " + err.Error())
					} else {
						atomic.AddInt64(&writeDone, 1)
					}
				}
			}
		})
	}

	// stock send all the file that need to be download
	err := downloadDir(manifest, "")
	if err != nil {
		se = append(se, err)
	}

	done := make(chan bool, 2)
	go grWaitGroup(&wg, done)

	go grMonitoring(done, mxConn)

	select {
	case <-Ctx.Done():
		dLog.Info("Request Cancel")
		return
	case <-done:
		for _, err := range sErrors {
			se = append(se, err)
		}
	}

	if len(se) == 0 {
		dLog.Info("Remove Manifest")
		err = os.Remove("Manifest.json")
		if err != nil {
			se = append(se, err)
		}
	}
}

func grWaitGroup(wg *sync.WaitGroup, done chan bool) {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")
	wg.Wait()
	dLog.Info("Game downloaded!")
	done <- true
	done <- true
}

func grMonitoring(done chan bool, mxConn mc.MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("Client/download")
	var (
		previousNbDlDone int64 = 0
		previousNbWrDone int64 = 0
		chanDlClosed     bool  = false
		chanWrClosed     bool  = false
	)

	for {
		select {
		case <-Ctx.Done():
			dLog.Info("Monitoring Request Cancel")
			return
		case <-done:
			mxConn.WriteJSON("Dl Done", mc.DownloadDone)
			return
		default:
			nbDlDone := atomic.LoadInt64(&downloadDone)
			dlSpeed := float64(nbDlDone-previousNbDlDone) * float64(utils.SizeChunk)
			dlSpeed /= 1000
			previousNbDlDone = nbDlDone

			nbWrDone := atomic.LoadInt64(&writeDone)
			wrSpeed := float64(nbWrDone-previousNbWrDone) * float64(utils.SizeChunk)
			wrSpeed /= 1000
			previousNbWrDone = nbWrDone

			mxConn.WriteJSON(MonitoringData{
				DlPrc:   float64(nbDlDone) / float64(NbChunk) * 100,
				DlSpeed: dlSpeed,
				WrPrc:   float64(nbWrDone) / float64(NbChunk) * 100,
				WrSpeed: wrSpeed,
			}, mc.Monitoring)

			dLog.Info(fmt.Sprintf("Download %f%% at %f kB/s | Write %f%% at %f kB/s", float64(nbDlDone)/float64(NbChunk)*100, dlSpeed, float64(nbWrDone)/float64(NbChunk)*100, wrSpeed))
			time.Sleep(time.Second)

			if nbDlDone >= int64(NbChunk) && !chanDlClosed {
				close(chanDownload)
				chanDlClosed = true
			}
			if nbWrDone >= int64(NbChunk) && !chanWrClosed {
				close(chanWrite)
				chanWrClosed = true
			}
		}
	}
}
