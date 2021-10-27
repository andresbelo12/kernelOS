package handler

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/andresbelo12/KernelOS/model"
)

const serverPORT = "8080"

type Server struct {
	Dictionary map[string]*model.ServerConnection
}

func InitServer() Server {
	return Server{Dictionary: make(map[string]*model.ServerConnection)}
}

func (server Server) InitServerConnection(listener model.CommunicationListener) (err error) {

	PORT := ":" + serverPORT

	connectionDoor, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer connectionDoor.Close()

	netConnection, err := connectionDoor.Accept()
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

	go ListenClient(listener, connection)
	go sendMessage(connection)
	return
}

func sendMessage(connection *model.ServerConnection) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Print("--->" + text)
		WriteConnection(connection, &model.Message{Command: model.CMD_SEND, Source: model.MD_KERNEL, Destination: model.MD_FILES, Message: "prueba"})
	}
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
