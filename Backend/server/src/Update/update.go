package update

import (
	butcher "aetova/server/src/Butcher"
	"aetova/server/utils"
	"fmt"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func UpdateGame(updatePath string, guid uint) error {
	dLog := docLogger.NewLoggerWithGOpts("server/UpdateGame")

	dLog.Info("Update start")

	var manifestGame utils.ManifestGame
	file, err := os.OpenFile(fmt.Sprintf("%s%d/AppManifest.json", butcher.ToDir, guid), os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	err = utils.FromJson(&manifestGame, file)
	if err != nil {
		return err
	}

	// Update version
	manifestGame.Version += 1
	err = utils.ToJson(manifestGame, file)
	if err != nil {
		return err
	}

	// Chop the update
	err = butcher.ChopGame(updatePath, manifestGame.Name, guid, manifestGame.Version)
	if err != nil {
		return err
	}

	// Compare it

	// Create the download update manifest

	dLog.Info("Update done")

	return nil
}
