package update

import (
	butcher "aetova/server/src/Butcher"
	"aetova/server/utils"
	"fmt"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func CompareGame(old utils.ManifestDir, new utils.ManifestDir) error {
	dLog := docLogger.NewLoggerWithGOpts("server/compareGame")
	dLog.Info("Compare Game start")

	manifestUDir := CompareManifestDir(old, new)

	err := CheckManifestDir(&manifestUDir, "")
	if err != nil {
		return err
	}

	// Save the update manifest
	file, err := os.Create(fmt.Sprintf("%s%d/UManifest/UMfs_%d.json", butcher.ToDir, currentGame.Guid, currentGame.Version)) // TODO : Get a better way and dir organisation
	if err != nil {
		return err
	}

	err = utils.ToJson(manifestUDir, file)
	if err != nil {
		return err
	}

	return nil
}
