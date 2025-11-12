package butcher

import (
	server "aetova/server/src"
	"aetova/server/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

func ChopGame(zip string) (err error) {
	start := time.Now()

	fmt.Println("Unzip start")

	// unzip the game send by the dev
	var unzipPath string
	if unzipPath, err = server.Unzip(zip); err != nil {
		return err
	}

	fmt.Println("Unzip done")

	// create the manifest.json for the game
	var manifestDir utils.ManifestDir
	manifestDir.Name = strings.Split(unzipPath, "/")[len(strings.Split(unzipPath, "/"))-1]
	if manifestDir, err = chopDir(unzipPath); err != nil {
		return err
	}

	// create the manifest file
	var file *os.File
	if file, err = os.Create("manifest.json"); err != nil {
		return err
	}

	// save the data into the file
	if err = utils.ToJson(manifestDir, file); err != nil {
		return err
	}

	fmt.Println("The process takes", time.Since(start))

	return err
}
