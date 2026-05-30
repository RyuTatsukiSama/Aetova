package update

import (
	"aetova/server/utils"
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

	file, err := os.Create("test.json") // TODO : Get a better way and dir organisation
	if err != nil {
		return err
	}

	err = utils.ToJson(manifestUDir, file)
	if err != nil {
		return err
	}

	return nil
}
