package gonipolib

import (
	"net"
	"strings"
	"fmt"
)

func Connect(connectionString string) (net.Conn,bool) {
	connectionStringFields := strings.Fields(connectionString)
	connection,err := net.Dial("tcp", connectionStringFields[2]+":"+connectionStringFields[3])
	ok := false
	if err != nil {
        fmt.Println("nipolib Error connecting to socket: "+err.Error())
	} else {
		ok = true
	}
	return connection,ok
}

func Write(cmd string, connection net.Conn) (bool) {
	response := make([]byte, 4096)
	connection.Write([]byte(cmd+"\n"))
	connection.Read(response)
	ok := false
	if string(response) == "OK" {
		fmt.Println("Login success")
		return true
	}
	return ok
}