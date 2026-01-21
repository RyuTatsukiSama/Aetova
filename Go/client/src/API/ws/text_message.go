package ws

import (
	"os"
	"path/filepath"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

const (
	download string = "download"
	pause    string = "pause"
	resume   string = "resume"
	cancel   string = "cancel"
)

func handleStringMessage(message string) {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	switch message {
	case download:
		dLog.Debug("download has been called")
		handleDownload(MxConn)
	case pause:
		dLog.Debug("pause has been called")
		cancelFunc()
	case resume:
		dLog.Debug("resume has been called")
		handlePostResume(MxConn)
	case cancel:
		dLog.Debug("cancel has been called")
		cancelFunc()
		dLog.Warning("Doesn't work currently")
		// TODO : Do it after Worker refactor
		err := deleteDownloadFiles("BuildOranys") // TODO : hardcode need to be yeet
		if err != nil {
			dLog.Error("Error 9 : " + err.Error())
			return
		}
	default:
		dLog.Error("Error 6 : Message " + message + " not listed")
	}
}

func deleteDownloadFiles(folderName string) error {
	// remove the game folder
	err := os.RemoveAll("downloads/" + folderName)
	if err != nil {
		return err
	}

	// remove the game manifest
	err = os.Remove("Manifest.json")
	if err != nil {
		return err
	}

	// remove the files manifest
	files, err := filepath.Glob("Manifest_*.bin")
	if err != nil {
		return err
	}

	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			return err
		}
	}

	return nil
}
