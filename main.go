package main

import (
	"fmt"

	"github.com/andresbelo12/KernelOS/handler"
	"github.com/andresbelo12/KernelOS/model"
)

func main() {
	a := model.Message{Command: "a"}
	fmt.Println(a)

	communicationServer()

}

func communicationServer() {
	server := handler.InitServer()

	for {
		server.InitServerConnection()
		fmt.Println(server.Dictionary)
	}

}
