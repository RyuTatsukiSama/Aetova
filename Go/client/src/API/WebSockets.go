package client_api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Create the connection with the webSocket
func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // TODO : Change that with the API key, or other to check security

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Client Connected!")

	err = ws.WriteMessage(1, []byte("Client Connected!"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	webSocketReader(ws)
}

func webSocketReader(conn *websocket.Conn) {
	lastTime := time.Now()
	var timer time.Duration = 0

	for {
		currentTime := time.Now()
		deltaTime := currentTime.Sub(lastTime)
		lastTime = currentTime

		if timer > time.Duration(time.Second*5) {
			timer = 0
			if err := conn.WriteMessage(1, []byte("Tu es là ? :)")); err != nil {
				log.Println(err)
				return
			}
		} else {
			timer += deltaTime
		}
	}
}
