package resume_download

import (
	"aetova/client/src/API/api_download"
	"aetova/client/utils"
	"fmt"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

var (
	downloadDone int64
	writeDone    int64
)

func CountChunk(manifestDir utils.ManifestDir, cd chan bool) {
	api_download.NbChunk = 0

	dLog := docLogger.NewLoggerWithGOpts("Client/BrowseManifest")

	browseDir(manifestDir)

	dLog.Info(fmt.Sprintf("%d dd %d wd", downloadDone, writeDone))

	cd <- true
}

func CountUChunk(manifestDir utils.ManifestUDir, cd chan bool) {
	api_download.NbChunk = 0

	dLog := docLogger.NewLoggerWithGOpts("Client/BrowseManifest")

	browseUDir(manifestDir)

	dLog.Info(fmt.Sprintf("%d dd %d wd", downloadDone, writeDone))

	cd <- true
}

func browseDir(manifestDir utils.ManifestDir) {
	for _, subDir := range manifestDir.SubDir {
		browseDir(subDir)
	}

	for _, subFile := range manifestDir.SubFiles {
		api_download.NbChunk += subFile.NbChunks

		bitmap := api_download.NewChunkBitmap(subFile.NbChunks)
		bitmap.LoadFromDisk(fmt.Sprintf("%s%s%s.mfs", "downloads/", "0/", subFile.Name))

		missingChunks := bitmap.MissingChunks()

		downloadDone += int64(subFile.NbChunks) - int64(len(missingChunks))
		writeDone += int64(subFile.NbChunks) - int64(len(missingChunks))
	}
}

func browseUDir(manifestDir utils.ManifestUDir) {
	for _, subDir := range manifestDir.SubDir {
		browseUDir(subDir)
	}

	for _, subFile := range manifestDir.SubFiles {
		api_download.NbChunk += subFile.New.NbChunks

		bitmap := api_download.NewUChunkBitmap(subFile.New.NbChunks, subFile.Chk_changes)
		bitmap.LoadFromDisk(fmt.Sprintf("%s%s%s.mfs", "downloads/", "0/", subFile.Name))

		missingChunks := bitmap.MissingChunks()

		downloadDone += int64(subFile.New.NbChunks) - int64(len(missingChunks))
		writeDone += int64(subFile.New.NbChunks) - int64(len(missingChunks))
	}
}
