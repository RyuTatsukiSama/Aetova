package ws

import (
	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

const (
	close    string = "close"
	download string = "download"
	pause    string = "pause"
	resume   string = "resume"
	cancel   string = "cancel"
)

func handleStringMessage(message string) bool {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	switch message {
	case close:
		dLog.Debug("close has been called")
		MxConn.closeConnection()
		return true
	case download:
		dLog.Debug("download has been called")
		HandleDownload(MxConn)
	case pause:
		dLog.Debug("pause has been called")
		cancelFunc()
	case resume:
		dLog.Debug("resume has been called")
		// handlePostResume
	case cancel:
		dLog.Debug("cancel has been called")
		// pause case + delete all file downloaded so fare
	default:
		dLog.Error("Error 6 : Message not listed")
	}

	return false
}
