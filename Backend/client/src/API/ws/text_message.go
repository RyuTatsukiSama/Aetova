package ws

import (
	mc "aetova/client/src/API/MutexConnection"
	"os"
	"path/filepath"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

const (
	download   string = "download"
	pause      string = "pause"
	resume     string = "resume"
	cancel     string = "cancel"
	update     string = "update"
	resume_upt string = "resume_upt"
)

func handleStringMessage(message string) {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	switch message {
	case download:
		dLog.Debug("download has been called")
		handleDownload(MxConn)
	case pause:
		dLog.Debug("pause has been called")
		mc.CancelFunc()
	case resume:
		dLog.Debug("resume has been called")
		handlePostResume(MxConn)
	case cancel:
		dLog.Debug("cancel has been called")
		mc.CancelFunc()
		dLog.Warning("Doesn't work currently")
		err := deleteDownloadFiles("BuildOranys") // TODO: hardcode need to be yeet
		if err != nil {
			dLog.Error("Error 9 : " + err.Error())
			return
		}
	case update:
		dLog.Debug("update has been called")
		handleUpdate(MxConn)
	case resume_upt:
		dLog.Debug("Resume Update has been called")
		handleResumeUpdate(MxConn)
	default:
		dLog.Error("Error 6 : Message " + message + " not listed")
	}
}

func deleteDownloadFiles(folderName string) error { // TODO: Find a way for it to work
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
