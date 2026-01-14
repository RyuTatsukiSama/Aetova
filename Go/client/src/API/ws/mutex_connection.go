package ws

import (
	"fmt"
	"sync"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
	"github.com/gorilla/websocket"
)

type MutexConnection struct {
	conn *websocket.Conn
	mx   *sync.Mutex
}

func (this *MutexConnection) WriteText(message string) {
	this.mx.Lock()
	err := this.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		this.mx.Unlock()
		fmt.Print(err.Error())
	}
	this.mx.Unlock()
}

func (mxConn *MutexConnection) closeConnection() {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")
	dLog.Debug("Stop Gouroutine")

	cancelFunc()

	dLog.Debug("Close connection")

	closeNormalClosure := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	mxConn.mx.Lock()
	err := mxConn.conn.WriteControl(websocket.CloseMessage, closeNormalClosure, time.Now().Add(time.Millisecond*125))
	if err != nil {
		dLog.Error("Error 1 " + err.Error())
	}

	err = mxConn.conn.Close()
	if err != nil {
		dLog.Error("Error 2 " + err.Error())
	}
	mxConn.mx.Unlock()

	dLog.Info("Connection closed")
}
