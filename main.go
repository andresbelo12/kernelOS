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

	for {
		server.InitServerConnection()
		fmt.Println(server.Dictionary)
	}

}
