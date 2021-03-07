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
	// IP Port
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

func (connection *Connection) Set(token string, key string, value string) (string, bool) {
	result,ok := connection.socketWrite(token + " set "+ key + " " + value)
	return result,ok
}

func (connection *Connection) Get(token string, key string) (string, bool) {
	result,ok := connection.socketWrite(token + " get "+ key)
	return result,ok
}

func (connection *Connection) Select(token string, key string) (string, bool) {
	result,ok := connection.socketWrite(token + " select "+ key)
	return result,ok
}

func (connection *Connection) Avg(token string, key string) (string, bool) {
	result,ok := connection.socketWrite(token + " avg "+ key)
	return result,ok
}