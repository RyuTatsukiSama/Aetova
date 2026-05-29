package butcher

import (
	server "aetova/server/src"
	"aetova/server/utils"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func RegisterGameManifest(gname string, guid uint, version uint) error {
	dLog := docLogger.NewLoggerWithGOpts("server/RegisterGameManifest")

	dLog.Info("Register start")

	// Create the manifest of the game
	manifestGame := utils.ManifestGame{
		Name:    gname,
		Version: version,
		Guid:    guid,
	}

	err := os.MkdirAll(fmt.Sprintf("%s%d", ToDir, guid), 0700)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fmt.Sprintf("%s%d/Manifest", ToDir, guid), 0700)
	if err != nil {
		return err
	}

	// create the manifest file
	file, err := os.Create(fmt.Sprintf("%s%d/AppManifest.json", ToDir, guid))
	if err != nil {
		return err
	}

	// save the data into the file
	err = utils.ToJson(manifestGame, file)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	dLog.Info("Register done")

	return nil
}

func ChopGame(zip string, gname string, guid uint, version uint) (err error) {
	dLog := docLogger.NewLoggerWithGOpts("Server/ChopGame")

	start := time.Now()

	dLog.Info("Unzip Start")

	// unzip the game send by the dev
	var path string
	if path, err = server.Unzip(zip, guid, version, gname); err != nil {
		return err
	}

	dLog.Info("Unzip Done")

	// create the manifest for download
	var manifestDir utils.ManifestDir
	manifestDir.Name = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	if manifestDir, err = chopDir(path); err != nil {
		return err
	}

	// create the manifest file
	file, err := os.Create(fmt.Sprintf("%s%d/Manifest/Mfs_%d.json", ToDir, guid, version))
	if err != nil {
		return err
	}

	// save the data into the file
	err = utils.ToJson(manifestDir, file)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	// Clear the unzip file
	err = os.RemoveAll(fmt.Sprintf("%s%d/", FromDir, guid))
	if err != nil {
		return err
	}

	dLog.Info(fmt.Sprintln("The process took", time.Since(start)))

	return err
}
