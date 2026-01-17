package ws

import (
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
)

func handleJsonMessage(mxConn MutexConnection) bool {
	dLog := docLogger.NewLoggerWithGOpts("Client/websocket")

	// TODO : Lock the mutex here soft lock the programm, but not so thread safe, find a way if possible to be more clean
	var message Message
	if err := mxConn.conn.ReadJSON(&message); err != nil {
		dLog.Error("Error 7 : " + err.Error())
		return false
	}

	dLog.Debug("message has been read")

	switch message.Type {
	case Text:
		var str string
		err := json.Unmarshal(message.Data, &str)
		if err != nil {
			dLog.Error("Error 8 :" + err.Error())
			return false
		}
		return handleStringMessage(str)
	default:
		dLog.Error("Error 5 : Message type not allowed")
		return false
	}

	return false
}
