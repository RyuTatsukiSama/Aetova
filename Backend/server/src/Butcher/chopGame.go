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

func ChopGame(zip string) (err error) {
	dLog := docLogger.NewLoggerWithGOpts("Server/ChopGame")

	start := time.Now()

	// create the row in data base
	guid := 0 // TODO: Change it when postgreSQL is setup
	// Generate a random GUID here, the GUID != ID in database

	dLog.Info("Unzip Start")

	// unzip the game send by the dev
	var path string
	if path, err = server.Unzip(zip, guid); err != nil {
		return err
	}

	dLog.Info("Unzip Done")

	// create the manifest.json for the game
	var manifestDir utils.ManifestDir
	manifestDir.Name = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	if manifestDir, err = chopDir(path); err != nil {
		return err
	}

	// create the manifest file
	var file *os.File
	if file, err = os.Create(fmt.Sprintf("Mfs_%d.json", guid)); err != nil {
		return err
	}

	// save the data into the file
	if err = utils.ToJson(manifestDir, file); err != nil {
		return err
	}

	dLog.Info(fmt.Sprintln("The process took", time.Since(start)))

	return err
}
