package main

import (
	"net"
	"strings"
	"fmt"
	"bufio"
)

type Connection struct {
	connection net.Conn
	connectionString string
}

func CreateConnection() *Connection {
	return &Connection {}
}

func socketConnect(connectionString string) (*Connection,bool) {
	connection := CreateConnection()
	connectionStringFields := strings.Fields(connectionString)
	socket,err := net.Dial("tcp", connectionStringFields[1]+":"+connectionStringFields[2])
	if err != nil {
        fmt.Println("nipolib Error connecting to socket: "+err.Error())
	}
	connection.connection = socket
	connection.connectionString = connectionString
	return connection,true
}

func (connection *Connection) socketWrite(cmd string) (string, bool) {
	connection.connection.Write([]byte(cmd+"\n"))
	response,_ := bufio.NewReader(connection.connection).ReadString('\n')
	ok := false
	if string(response) != "" {
		return string(response),true
	}
	return string(response),ok
}

func (connection *Connection) socketClose() (bool) {
	socket := connection.connection
	socket.Close()
	return true
}