package ws

import (
	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

const (
	close    string = "close"
	download string = "download"
	pause    string = "pause"
	resume   string = "resume"
)

func handleStringMessage(message string) bool {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	switch message {
	case close:
		MxConn.closeConnection()
		return true
	case download:

	case pause:
	case resume:
	default:
		dLog.Error("Error 6 : Message not listed")
	}

	return false
}
