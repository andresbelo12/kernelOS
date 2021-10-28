package handler

import (
	"fmt"

	"github.com/andresbelo12/KernelOS/model"
)

type ServerListener struct{}

func CreateListener() model.CommunicationListener {
	return ServerListener{}
}

func (listener ServerListener) ProcessMessage(conncetion interface{}, message *model.Message) (err error) {
	fmt.Println(message)

	return
}
