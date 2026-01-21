package ws

import (
	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func webSocketReader(mxConn MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	// For receiving message
	for {
		go handleJsonMessage(mxConn)

		<-chanClose
		dLog.Info("Close order given")
		// close all the goroutine
		return
	}
}
