package ws

import (
	"fmt"
	"sync"

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
