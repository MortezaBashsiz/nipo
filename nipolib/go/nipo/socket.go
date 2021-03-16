package nipo

import (
	"net"
	"strings"
	"fmt"
	"bufio"
)

func socketConnect(connectionString string) (*Connection,bool) {
	connection := CreateConnection()
	connectionStringFields := strings.Fields(connectionString)
	socket,err := net.Dial("tcp", connectionStringFields[1]+":"+connectionStringFields[2])
	if err != nil {
        fmt.Println("nipolib Error connecting to socket: "+err.Error())
		return connection,false
	}
	connection.socket = socket
	connection.connectionString = connectionString
	return connection,true
}

func (connection *Connection) socketWrite(cmd string) (string, bool) {
	connection.socket.Write([]byte(cmd+"\n"))
	response,_ := bufio.NewReader(connection.socket).ReadString('\n')
	ok := false
	if string(response) != "" {
		return string(response),true
	}
	return string(response),ok
}

func (connection *Connection) socketClose() (bool) {
	socket := connection.socket
	socket.Close()
	return true
}