package nipo

import (
	"net"
	"fmt"
	"strings"
)

type Connection struct {
	socket net.Conn
	connectionString string
}

func CreateConnection() *Connection {
	return &Connection {}
}

func OpenConnection(connectionString string) (Connection, bool) {
	// user pass IP port
	connection,ok := socketConnect(connectionString)
	connection.connectionString = connectionString
	if !ok {
		fmt.Println("nipolib Error connecting to socket: ")
	}
	return *connection,ok
}

func (connection *Connection) Logout() bool {
	connection.Close()
	return true
}

func Login(connectionString string) (Connection, bool) {
	connection,ok := OpenConnection(connectionString)
	connectionStringFields := strings.Fields(connectionString)
	username := connectionStringFields[0]
	password := connectionStringFields[1]
	cmdLogin := "login " + username + " " + password
	if connection.socketLogin(cmdLogin){
		return connection,true
	}
	return connection,ok
}

func (connection *Connection) Set(key string, value string) (string, bool) {
	result,ok := connection.socketWrite("set "+ key + " " + value)
	return result,ok
}

func (connection *Connection) Get(key string) (string, bool) {
	result,ok := connection.socketWrite("get "+ key)
	return result,ok
}