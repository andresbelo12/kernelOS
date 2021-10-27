package handler

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/andresbelo12/KernelOS/model"
)

const serverPORT = 8080

type Server struct {
	Dictionary map[string]*model.ServerConnection
}

func InitServer() Server {
	return Server{Dictionary: make(map[string]*model.ServerConnection)}
}

func (server Server) InitServerConnection() (err error) {

	PORT := ":" + strconv.Itoa(serverPORT)

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer listener.Close()

	netConnection, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return err
	}

	message, err := ReadMessage(&netConnection)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	connection := server.RegisterConnection(&netConnection, message)

	go ListenClient(connection)

	return
}

func (server Server) AddConnection(connection *model.ServerConnection) {
	server.Dictionary[(*connection).ClientName] = connection
}

func (server Server) RegisterConnection(netConnection *net.Conn, message model.Message) *model.ServerConnection {
	var connection model.ServerConnection

	connectionInfo := strings.Split((*netConnection).RemoteAddr().String(), ":")
	connection.ClientHost = connectionInfo[0]
	connection.ClientPort = connectionInfo[1]
	connection.ClientName = message.Source
	connection.ClientConnection = netConnection

	server.AddConnection(&connection)
	return &connection
}

