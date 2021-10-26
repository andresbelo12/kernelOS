package model

import (
	"encoding/json"
	"fmt"
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
