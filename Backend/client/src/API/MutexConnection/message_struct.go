package mutexconnection

import "encoding/json"

type Message struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"Data"`
}

type MessageType uint

const (
	Text MessageType = iota
	Close
	Exit
	Monitoring
	DownloadDone
)
