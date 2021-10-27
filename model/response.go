package model

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	CodTerm int8   `json:"codterm"`
	Message string `json:"msg"`
}

func (rsp Response) ToJson() (byteMessage []byte) {
	byteMessage, err := json.Marshal(rsp)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	return
}

func ToResponse(byteResponse []byte, bufferLength int) (response Response, err error) {
	err = json.Unmarshal(byteResponse[:bufferLength], &response)
	return
}
