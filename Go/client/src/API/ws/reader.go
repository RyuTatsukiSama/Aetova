package ws

import (
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func webSocketReader(mxConn MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	// For receiving message
	for {
		go func() {
			for {
				select {
				case <-ctx.Done():
					dLog.Debug("Connection being close")
					return
				default:
					dLog.Debug("U there ? :)")
					mxConn.WriteText("U there ? :)")
					time.Sleep(time.Second)
				}
			}
		}()

		if isClosed := handleJsonMessage(mxConn); isClosed {
			dLog.Info("Close order given")
			return
		}
	}
}
