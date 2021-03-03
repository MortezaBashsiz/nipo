package nipo

import (
	"net"
	"strings"
	"fmt"
)

func socketConnect(connectionString string) (*Connection,bool) {
	connection := CreateConnection()
	connectionStringFields := strings.Fields(connectionString)
	socket,err := net.Dial("tcp", connectionStringFields[2]+":"+connectionStringFields[3])
	if err != nil {
        fmt.Println("nipolib Error connecting to socket: "+err.Error())
	}
	connection.socket = socket
	connection.connectionString = connectionString
	return connection,true
}

func (connection *Connection) Close() (bool) {
	socket := connection.socket
	socket.Close()
	return true
}

func (connection *Connection) socketLogin(cmd string) (bool) {
	response := make([]byte, 4096)
	connection.socket.Write([]byte(cmd+"\n"))
	connection.socket.Read(response)
	ok := false
	if string(response) == "OK" {
		fmt.Println("Login success")
		return true
	}
	return ok
}

func (connection *Connection) socketWrite(cmd string) (string, bool) {
	connection.socket.Write([]byte(cmd+"\n"))
	response := make([]byte, 4096)
	connection.socket.Read(response)
	return string(response),true
}