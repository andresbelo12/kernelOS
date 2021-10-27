package model

import (
	"encoding/json"
	"fmt"
)

const(
	MD_FILES    = "FILES"
	MD_KERNEL   = "KERNEL"
	MD_GUI      = "GUI"
	CMD_START   = "start"
	CMD_SEND    = "send"
	CMD_INFO    = "info"
	CMD_STOP    = "stop"
)

type Message struct {
	Command     string `json:"cmd"`
	Source      string `json:"src"`
	Destination string `json:"dst"`
	Message     string `json:"msg"`
}

func (msg Message) ToJson() (byteMessage []byte) {
	byteMessage, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	return
}

func ToMessage(byteMessage []byte, bufferLength int)(message Message,err error){
	err = json.Unmarshal(byteMessage[:bufferLength], &message)
	return
}