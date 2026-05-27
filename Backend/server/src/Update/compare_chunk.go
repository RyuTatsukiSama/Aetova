package update

import (
	"aetova/server/utils"
	"fmt"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func CompareChunk(old utils.ManifestFile, new utils.ManifestFile, isAdding bool) {
	dLog := docLogger.NewLoggerWithGOpts("server/compare_chunk")

	limit := int(old.NbChunks)
	if isAdding {
		limit = int(new.NbChunks)
	}

	var chk_change []int
	for chk_id := 0; chk_id < limit-1; chk_id++ {
		oldChunkName := fmt.Sprintf("Part%d_%s.chk", chk_id, old.Name)
		oldChunkData, err := os.ReadFile("chop/" + oldChunkName)
		if err != nil {
			dLog.Error(err.Error())
		}

		newChunkName := fmt.Sprintf("Part%d_%s.chk", chk_id, new.Name)
		newChunkData, err := os.ReadFile("chop/" + newChunkName)
		if err != nil {
			dLog.Error(err.Error())
		}

		for byte_id := 0; byte_id < int(utils.SizeChunk); byte_id++ {
			if oldChunkData[byte_id] != newChunkData[byte_id] {
				chk_change = append(chk_change, chk_id)
				break
			}
		}
	}

	if isAdding {
		for i := old.NbChunks - 1; i < new.NbChunks; i++ {
			chk_change = append(chk_change, int(i))
		}
	}

	dLog.Debug(fmt.Sprintln(len(chk_change), new.NbChunks))
}
