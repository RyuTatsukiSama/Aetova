package client_api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

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

// Create the connection with the webSocket
func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	dLog := docLogger.NewLoggerWithGOpts("CLient/websocket")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // TODO : Change that with the API key, or other to check security

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		dLog.Log(docLogger.Critical, err.Error())
		return
	}

	MxConn = MutexConnection{
		conn: ws,
		mx:   &sync.Mutex{},
	}

	dLog.Log(docLogger.Info, "Client connected!")

	err = ws.WriteMessage(websocket.TextMessage, []byte("Client Connected!"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	webSocketReader(MxConn)
}

func webSocketReader(mxConn MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	ctx, cancel = context.WithCancel(context.Background())
	chanClose := make(chan bool)

	// For receiving message
	for {
		messageType, message, err := mxConn.conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}

		switch messageType {
		case websocket.CloseMessage:
			closeConnection(mxConn, chanClose)
			return
		case websocket.TextMessage:
			dLog.Log(docLogger.Debug, string(message))
			handleTextMessage(string(message), mxConn, chanClose)
		default:
			mxConn.mx.Lock()
			err := mxConn.conn.WriteMessage(websocket.TextMessage, []byte("Message type not allowed"))
			if err != nil {
				mxConn.mx.Unlock()
				log.Fatal(err)
			}
			mxConn.mx.Unlock()
		}
	}
}

func handlePongMessage(mxConn MutexConnection) {
	mxConn.mx.Lock()
	err := mxConn.conn.WriteMessage(websocket.TextMessage, []byte("Pong"))
	if err != nil {
		mxConn.mx.Unlock()
		log.Fatal(err)
	}
	mxConn.mx.Unlock()
}

func handleTextMessage(message string, mxConn MutexConnection, chanClose chan bool) {
	switch message {
	case "download":
		go HandlePostDownload(mxConn)
	case "close":
		closeConnection(mxConn, chanClose)
		return
	case "stop":
		cancel()
	case "resume":
		go HandlePostResume(mxConn)
	case "ping":
		handlePongMessage(mxConn)
	default:
		mxConn.mx.Lock()
		err := mxConn.conn.WriteMessage(websocket.TextMessage, []byte("Message not allowed"))
		if err != nil {
			mxConn.mx.Unlock()
			log.Fatal(err)
		}
		mxConn.mx.Unlock()
	}
}

func closeConnection(mxConn MutexConnection, chanClose chan bool) {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")
	dLog.Log(docLogger.Debug, "Try to close connection")

	cancel()

	dLog.Log(docLogger.Debug, "Cancel")

	<-chanClose

	dLog.Log(docLogger.Debug, "Send close to channel")

	closeNormalClosure := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	err := mxConn.conn.WriteControl(websocket.CloseMessage, closeNormalClosure, time.Now().Add(time.Millisecond*125))
	if err != nil {
		log.Fatal(err)
	}

	dLog.Log(docLogger.Debug, "Send close message")

	err = mxConn.conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	dLog.Log(docLogger.Info, "Connection Closed!")
}
