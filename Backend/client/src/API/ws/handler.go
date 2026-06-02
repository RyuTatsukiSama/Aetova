package ws

import (
	mc "aetova/client/src/API/MutexConnection"
	"net/http"
	"sync"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	MxConn mc.MutexConnection
)

// Create the connection with the webSocket
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	dLog := docLogger.NewLoggerWithGOpts("CLient/websocket")

	initGoroutineVar(r)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // TODO : Change that with the API key, or other to check security

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		dLog.Log(docLogger.Critical, err.Error())
		return
	}

	MxConn = mc.MutexConnection{
		Conn:    ws,
		WriteMx: &sync.Mutex{},
	}

	dLog.Log(docLogger.Info, "Client connected!")
	MxConn.WriteJSON("Client Connected!", mc.Text)

	err = CheckUpdate(MxConn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		dLog.Error(err.Error())
		return
	}

	webSocketReader(MxConn)
}
