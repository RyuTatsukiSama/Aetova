package ws

import (
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func webSocketReader(mxConn MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	// For receiving message
	for {
		go handleJsonMessage(mxConn)

		ExitOrder := <-chanClose
		dLog.Info("Close order given")
		// close all the goroutine
		if ExitOrder {
			os.Exit(0)
		} else {
			return
		}
	}
}
