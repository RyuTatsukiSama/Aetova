package update

import (
	butcher "aetova/server/src/Butcher"
	"aetova/server/utils"
	"fmt"
	"os"
)

func CompareChunks(old utils.ManifestFile, new utils.ManifestFile, isAdding bool, path string) ([]int, error) {
	// dLog := docLogger.NewLoggerWithGOpts("server/compare_chunk")

	limit := int(old.NbChunks)
	if isAdding {
		limit = int(new.NbChunks)
	}

	// start := time.Now()

	var chk_change []int
	for chk_id := 0; chk_id < limit-1; chk_id++ {
		oldChunkName := fmt.Sprintf("Part%d_%s.chk", chk_id, old.Name)
		oldChunkData, err := os.ReadFile(fmt.Sprintf("%s%d/%d%s/%s", butcher.ToDir, currentGame.Guid, currentGame.Version-1, path, oldChunkName))
		if err != nil {
			return make([]int, 0), err
		}

		newChunkName := fmt.Sprintf("Part%d_%s.chk", chk_id, new.Name)
		newChunkData, err := os.ReadFile(fmt.Sprintf("%s%d/%d%s/%s", butcher.ToDir, currentGame.Guid, currentGame.Version, path, newChunkName))
		if err != nil {
			return make([]int, 0), err
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
	} else if old.NbChunks == new.NbChunks && old.Size != new.Size {
		chk_change = append(chk_change, int(new.NbChunks-1))
	}

	// dLog.Debug(fmt.Sprintf("File %s took %s with %d as the limit", new.Name, time.Since(start)/time.Duration(limit), limit))

	return chk_change, nil
}
