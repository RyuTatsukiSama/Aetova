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
		MxConn.closeConnection()
		return true
	case download:
		HandlePostDownload(MxConn)
	case pause:
		// stop download goroutine
	case resume:
		// handlePostResume
	case cancel:
		// pause case + delete all file downloaded so fare
	default:
		dLog.Error("Error 6 : Message not listed")
	}

	return false
}
