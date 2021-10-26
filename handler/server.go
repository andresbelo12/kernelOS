package handler

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/andresbelo12/KernelOS/model"
)

const serverPORT = 8080

func InitServerConnection() (*net.Conn, error) {

	PORT := ":" + strconv.Itoa(serverPORT)

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer listener.Close()

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(connection.RemoteAddr().String())
	//GetConnectionType(connection)
	return &connection, err

}

func GetConnectionType(connection net.Conn) string {
	netData, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(netData)
	t := time.Now()
	myTime := t.Format(time.RFC3339) + "\n"
	connection.Write([]byte(myTime))
	return ""
}

func ListenConnection(connection net.Conn) {
	defer connection.Close()

	for {

		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("-> ", string(netData))
		a := model.Message{Destination: "aa", Source: "kernel"}
		connection.Write(a.ToJson())
	}

}


