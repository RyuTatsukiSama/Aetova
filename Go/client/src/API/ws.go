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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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

	mxConn := MutexConnection{
		conn: ws,
		mx:   &sync.Mutex{},
	}

	fmt.Println("Client Connected!")

	err = ws.WriteMessage(websocket.TextMessage, []byte("Client Connected!"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	webSocketReader(mxConn)
}

func webSocketReader(mxConn MutexConnection) {
	lastTime := time.Now()
	var timer time.Duration = 0
	ctx, cancel = context.WithCancel(context.Background())
	chanClose := make(chan bool)

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
					if err := mxConn.conn.WriteMessage(1, []byte("U there ? :)")); err != nil {
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

	for {
		_, message, err := mxConn.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			cancel()
			return
		}

		fmt.Println(string(message))

		if string(message) == "stop" {
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
			return
		}
	}
}
