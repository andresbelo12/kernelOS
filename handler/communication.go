package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/andresbelo12/KernelOS/model"
)

func EstablishClient(clientInfo *model.ClientConnection, message model.Message)(err error) {
	connection, err := net.Dial("tcp", clientInfo.ServerHost + ":" + clientInfo.ServerPort)
	if err != nil {
		return
	}
	
	if _, err = connection.Write(message.ToJson()); err != nil{
		return
	}

	clientInfo.ServerConnection = &connection

	return
}

func ReadMessage(connection *net.Conn) (message model.Message, err error) {
	buffer := make([]byte, 1024)
	bufferLength, err := (*connection).Read(buffer)

	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err)
		}
		return
	}

	if err = json.Unmarshal(buffer[:bufferLength], &message); err != nil {
		return
	}

	return
}

func ListenConnection(connection *model.ServerConnection) {
	for {
		message, err := ReadMessage(connection.ClientConnection)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(message)
	}
}

func WriteConnection(connection *model.ServerConnection, message *model.Message) {
	if _, err := (*connection.ClientConnection).Write(message.ToJson()); err != nil {
		panic(err)
	}
}
