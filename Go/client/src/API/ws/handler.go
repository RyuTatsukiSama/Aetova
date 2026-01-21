package ws

import (
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
	MxConn MutexConnection
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

	MxConn = MutexConnection{
		conn:    ws,
		writeMx: &sync.Mutex{},
	}

	dLog.Log(docLogger.Info, "Client connected!")

	err = ws.WriteMessage(websocket.TextMessage, []byte("Client Connected!"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		dLog.Log(docLogger.Critical, err.Error())
		return
	}

	webSocketReader(MxConn)
}
