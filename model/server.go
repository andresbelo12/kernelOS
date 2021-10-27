package model

import "net"

type (

	ClientConnection struct{
		ServerHost string
		ServerPort string
		ServerConnection *net.Conn
	}

	ServerConnection struct {
		ClientName       string
		ClientPort       string
		ClientHost       string
		ClientConnection *net.Conn
	}

	CommunicationListener interface {
		ProcessMessage(processorTools interface{}, message *Message)(error)
	}
)
