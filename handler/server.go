package handler

import (
	"fmt"
	"net"
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
	//go sendMessage(connection)
	return
}

/*func sendMessage(connection *model.ServerConnection) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Println(strings.ReplaceAll(text, "\r\n", ""))
		a := strings.ReplaceAll(text, "\r\n", "")
		var err error
		if a == "a" {
			err = WriteClient(connection, &model.Message{Command: model.CMD_SEND, Source: model.MD_KERNEL, Destination: model.MD_FILES, Message: "log:created"})
		}
		if a == "b" {
			err = WriteClient(connection, &model.Message{Command: model.CMD_SEND, Source: model.MD_KERNEL, Destination: model.MD_GUI, Message: "un mensaje x"})
		}
		if a == "c" {
			err = WriteClient(connection, &model.Message{Command: model.CMD_SEND, Source: model.MD_KERNEL, Destination: model.MD_GUI, Message: "action:created"})
		}
		if a == "d" {
			err = WriteClient(connection, &model.Message{Command: model.CMD_SEND, Source: model.MD_GUI, Destination: model.MD_FILES, Message: "create:sample"})
		}
		if a == "e" {
			err = WriteClient(connection, &model.Message{Command: model.CMD_SEND, Source: model.MD_GUI, Destination: model.MD_FILES, Message: "delete:sample"})
		}
		if err != nil {
			return
		}

	}
}*/

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
