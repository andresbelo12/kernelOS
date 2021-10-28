package handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/andresbelo12/KernelOS/model"
)

type ServerListener struct {
	Server *Server
}

func CreateListener(server *Server) model.CommunicationListener {
	return ServerListener{Server: server}
}

func (listener ServerListener) ProcessMessage(connection interface{}, message *model.Message) error {

	if message.Destination != model.MD_FILES {
		if err := listener.CreateLog(message); err != nil {
			log.Println("Could not create log for " + string(message.ToJson()) + " because: " + err.Error())
		}
	}

	clientConnection := connection.(**model.ServerConnection)

	if message.Destination == model.MD_KERNEL {
		listener.KernelFunctions(*clientConnection, message)
	}

	failureMessage := model.Message{
		Command:     model.CMD_SEND,
		Source:      model.MD_FILES,
		Destination: message.Source,
		Message:     "response:false;operation:" + message.Message + ";message:Reason ",
	}

	fmt.Println(message)
	if message.Destination != model.MD_KERNEL {
		if err := listener.ClientGateway(message); err != nil {
			failureMessage.Message += err.Error()
			if err = WriteClient(*clientConnection, &failureMessage); err != nil {
				log.Fatal(err.Error())
				return err
			}
		}
	}
	return nil
}

func (listener ServerListener) KernelFunctions(connection *model.ServerConnection, message *model.Message) {
	fmt.Println(message)
}

func (listener ServerListener) CreateLog(message *model.Message) (err error) {
	if nextConnection, exist := listener.Server.Dictionary[model.MD_FILES]; exist {
		err = WriteClient(nextConnection, message)
	} else {
		return errors.New("module " + message.Destination + " not registered")
	}
	return
}

func (listener ServerListener) ClientGateway(message *model.Message) (err error) {
	if nextConnection, exist := listener.Server.Dictionary[message.Destination]; exist {
		err = WriteClient(nextConnection, message)
	} else {
		return errors.New("module " + message.Destination + " not registered")
	}
	return
}
