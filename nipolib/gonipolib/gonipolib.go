package gonipolib

import (
	"net"
	"fmt"
	"strings"
)

func OpenConnection(connectionString string) (net.Conn,bool) {
	// user pass IP port
	connection,ok := Connect(connectionString)
	if !ok {
		fmt.Println("nipolib Error connecting to socket: ")
	}
	return connection,ok
}

func CloseConnection(connection net.Conn) bool {
	connection.Close()
	return true
}

func Login(connectionString string) bool {
	connection,_ := OpenConnection(connectionString)
	connectionStringFields := strings.Fields(connectionString)
	username := connectionStringFields[0]
	password := connectionStringFields[1]
	connection.Write([]byte("login " + username + " " + password))
	return true
}