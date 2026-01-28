package ws

import (
	mc "aetova/client/src/API/MutexConnection"
	"encoding/json"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

type Message struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"Data"`
}

type MessageType uint

const (
	Text MessageType = iota
	Close
	Exit
)

func handleJsonMessage(mxConn mc.MutexConnection) {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	// TODO : Lock the mutex here soft lock the programm, but not so thread safe, find a way if possible to be more clean
	var message Message
	if err := mxConn.Conn.ReadJSON(&message); err != nil {
		dLog.Error("Error 7 : " + err.Error())
		return
	}

	dLog.Debug("message has been read")

	// start another goroutine to handle multiple message
	if message.Type != Close && message.Type != Exit {
		go handleJsonMessage(mxConn)
	}

	switch message.Type {
	case Text:
		var str string
		err := json.Unmarshal(message.Data, &str)
		if err != nil {
			dLog.Error("Error 8 :" + err.Error())
			return
		}
		handleStringMessage(str)
	case Close:
		dLog.Debug("close has been called")
		MxConn.CloseConnection()
		chanClose <- false
	case Exit:
		dLog.Debug("Exit has been called")
		MxConn.CloseConnection()
		chanClose <- true
	default:
		dLog.Error("Error 5 : Message type not allowed")
	}
}
