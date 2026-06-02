package update

import (
	butcher "aetova/server/src/Butcher"
	"aetova/server/utils"
	"fmt"
	"os"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

var currentGame utils.ManifestGame

func UpdateGame(updatePath string, guid uint) error {
	dLog := docLogger.NewLoggerWithGOpts("server/UpdateGame")

	dLog.Info("Update start")

	start := time.Now()

	err := os.MkdirAll(fmt.Sprintf("%s%d/UManifest", butcher.ToDir, guid), 0777)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fmt.Sprintf("%s%d/AppManifest.json", butcher.ToDir, guid), os.O_RDONLY, 0777)
	if err != nil {
		return err
	}

	err = utils.FromJson(&currentGame, file)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	// Update version
	currentGame.Version += 1

	// Chop the update
	err = butcher.ChopGame(updatePath, currentGame.Name, guid, currentGame.Version)
	if err != nil {
		return err
	}

	// Get current manifest and previous one
	currentFile, err := os.OpenFile(fmt.Sprintf("%s%d/Manifest/Mfs_%d.json", butcher.ToDir, guid, currentGame.Version), os.O_RDONLY, 0777)
	if err != nil {
		return err
	}

	var currentManifest utils.ManifestDir
	err = utils.FromJson(&currentManifest, currentFile)
	if err != nil {
		return err
	}

	previousFile, err := os.OpenFile(fmt.Sprintf("%s%d/Manifest/Mfs_%d.json", butcher.ToDir, guid, currentGame.Version-1), os.O_RDONLY, 0777)
	if err != nil {
		return err
	}

	var previousManifest utils.ManifestDir
	err = utils.FromJson(&previousManifest, previousFile)
	if err != nil {
		return err
	}

	// Compare it and create the update manifest
	err = CompareGame(previousManifest, currentManifest)
	if err != nil {
		return err
	}

	// Save the updated version
	file, err = os.OpenFile(fmt.Sprintf("%s%d/AppManifest.json", butcher.ToDir, guid), os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}

	err = utils.ToJson(currentGame, file)
	if err != nil {
		return err
	}

	// Clear current game, just in case
	currentGame = utils.ManifestGame{}

	dLog.Info(fmt.Sprintln("Update done, it took", time.Since(start)))

	return nil
}
