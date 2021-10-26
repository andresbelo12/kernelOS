package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/andresbelo12/KernelOS/model"
)

func ReadMessage(connection *net.Conn)(message model.Message){
	buffer := make([]byte, 1024)

	_, err := (*connection).Read(buffer)
	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err)
		}
		panic(err)
	}
	
	if err := json.Unmarshal(buffer, &message); err != nil {
		panic(err)
	}
	return 
}
