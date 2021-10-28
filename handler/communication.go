package handler

import (
	"fmt"
	"io"
	"net"

	"github.com/andresbelo12/KernelOS/model"
)

func EstablishClient(clientInfo *model.ClientConnection, message model.Message) (err error) {
	connection, err := net.Dial("tcp", clientInfo.ServerHost+":"+clientInfo.ServerPort)
	if err != nil {
		return
	}

	if _, err = connection.Write(message.ToJson()); err != nil {
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

	if message, err = model.ToMessage(buffer, bufferLength); err != nil {
		return
	}

	return
}

func ListenClient(listener model.CommunicationListener, connection *model.ServerConnection) (err error) {
	defer (*connection.ClientConnection).Close()

	for {
		message, err := ReadMessage(connection.ClientConnection)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		listener.ProcessMessage(&connection, &message)
	}
}

func ListenServer(listener model.CommunicationListener, connection *model.ClientConnection) (err error) {
	defer (*connection.ServerConnection).Close()

	for {
		message, err := ReadMessage(connection.ServerConnection)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		listener.ProcessMessage(&connection, &message)
	}
}

func WriteServer(connection *model.ClientConnection, message *model.Message) (err error) {
	if _, err := (*connection.ServerConnection).Write(message.ToJson()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return
}

func WriteClient(connection *model.ServerConnection, message *model.Message) (err error) {
	if _, err := (*connection.ClientConnection).Write(message.ToJson()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return
}
