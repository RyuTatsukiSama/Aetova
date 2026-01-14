package ws

import (
	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func webSocketReader(mxConn MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	// For receiving message
	for {
		if isClosed := handleJsonMessage(mxConn); isClosed {
			dLog.Info("Close order given")
			return
		}
	}
}
