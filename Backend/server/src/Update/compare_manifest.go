package update

import (
	"aetova/server/utils"
	"fmt"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func CompareManifestDir(old utils.ManifestDir, new utils.ManifestDir) {

}

func CompareManifestFile(old utils.ManifestFile, new utils.ManifestFile) {
	dLog := docLogger.NewLoggerWithGOpts("server/cmf")
	if old.Name != new.Name {
		dLog.Error("Name different")
	}

	chunksDif := new.NbChunks - old.NbChunks
	sizeDif := new.Size - old.Size
	if chunksDif != 0 || sizeDif != 0 {
		dLog.Debug(fmt.Sprintf("ChunksDif %d SizeDif %d", chunksDif, sizeDif))
	}

	CompareChunk(old, new, new.Size > old.Size)

}
