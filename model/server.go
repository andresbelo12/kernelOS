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

	Server interface {
		InitServerConnection() (err error)
		AddConnection(connection *ServerConnection)
		RegisterConnection(netConnection *net.Conn, message Message) *ServerConnection
	}
)
