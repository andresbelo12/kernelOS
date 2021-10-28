package main

import (
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
	listener := handler.CreateListener(&server)
	for {
		server.InitServerConnection(listener)
	}

}
