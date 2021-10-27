package main

import (
	"fmt"

	"github.com/andresbelo12/KernelOS/handler"
)

func main() {

	LaunchModules()
	communicationServer()

}

func LaunchModules() {

}

func communicationServer() {
	server := handler.InitServer()
	listener := handler.CreateListener()
	for {
		server.InitServerConnection(listener)
		fmt.Println(server.Dictionary)
	}

}
