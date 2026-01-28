package mutexconnection

import (
	"context"
	"sync"
	"time"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
	"github.com/gorilla/websocket"
)

type MutexConnection struct {
	Conn    *websocket.Conn
	WriteMx *sync.Mutex
	ReadMx  *sync.Mutex
}

var CancelFunc context.CancelFunc

func (this *MutexConnection) WriteText(message string) {
	dLog := docLogger.NewLoggerWithGOpts("Client/MutexConnection")
	this.WriteMx.Lock()
	err := this.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		this.WriteMx.Unlock()
		dLog.Error(err.Error())
	}
	this.WriteMx.Unlock()
}

func (this *MutexConnection) WriteJSON(message interface{}) {
	dLog := docLogger.NewLoggerWithGOpts("Client/MutexConnection")
	this.WriteMx.Lock()
	err := this.Conn.WriteJSON(message)
	if err != nil {
		this.WriteMx.Unlock()
		dLog.Error(err.Error())
	}
	this.WriteMx.Unlock()
}

func (mxConn *MutexConnection) CloseConnection() {
	dLog := docLogger.NewLoggerWithGOpts("Client/MutexConnection")
	dLog.Debug("Stop Gouroutine")

	CancelFunc()

	dLog.Debug("Close connection")

	closeNormalClosure := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	mxConn.WriteMx.Lock()
	err := mxConn.Conn.WriteControl(websocket.CloseMessage, closeNormalClosure, time.Now().Add(time.Millisecond*125))
	if err != nil {
		dLog.Error("Error 1 " + err.Error())
	}

	err = mxConn.Conn.Close()
	if err != nil {
		dLog.Error("Error 2 " + err.Error())
	}
	mxConn.WriteMx.Unlock()

	dLog.Info("Connection closed")
}
