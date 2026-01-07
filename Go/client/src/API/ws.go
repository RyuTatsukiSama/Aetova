package client_api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

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

// Create the connection with the webSocket
func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // TODO : Change that with the API key, or other to check security

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	MxConn = MutexConnection{
		conn: ws,
		mx:   &sync.Mutex{},
	}

	fmt.Println("Client Connected!")

	err = ws.WriteMessage(websocket.TextMessage, []byte("Client Connected!"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	webSocketReader(MxConn)
}

func webSocketReader(mxConn MutexConnection) {
	lastTime := time.Now()
	var timer time.Duration = 0
	ctx, cancel = context.WithCancel(context.Background())
	chanClose := make(chan bool)

	// send message every seconds
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				chanClose <- true
				return
			default:
				currentTime := time.Now()
				deltaTime := currentTime.Sub(lastTime)
				lastTime = currentTime

				if timer > time.Duration(time.Second) {

					timer = 0
					mxConn.mx.Lock()
					err := mxConn.conn.WriteMessage(websocket.TextMessage, []byte("U there ? :)"))
					if err != nil {
						log.Println(err)
						mxConn.mx.Unlock()
						return
					}

					mxConn.mx.Unlock()
				} else {
					timer += deltaTime
				}
			}
		}
	}(ctx)

	// For receiving message
	for {
		_, message, err := mxConn.conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}

		fmt.Println(string(message))

		switch string(message) {
		case "download":
			// HandlePostDownload(mxConn)
		case "stop":
			closeConnection(mxConn, chanClose)
			return
		default:
			mxConn.mx.Lock()
			err = mxConn.conn.WriteMessage(websocket.TextMessage, []byte(""))
			if err != nil {
				mxConn.mx.Unlock()
				log.Fatal(err)
			}
			mxConn.mx.Unlock()
		}
	}
}

func closeConnection(mxConn MutexConnection, chanClose chan bool) {
	cancel()
	<-chanClose
	closeNormalClosure := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	err := mxConn.conn.WriteControl(websocket.CloseMessage, closeNormalClosure, time.Now().Add(time.Millisecond*125))
	if err != nil {
		log.Fatal(err)
	}

	err = mxConn.conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
