package nipo

import (
	"net"
	"fmt"
)

type Connection struct {
	socket net.Conn
	connectionString string
}

type Config struct {
	token, server,	port	string
}

func CreateConfig(token, server, port string) *Config {
	return &Config {
		token	: token ,
		server 	: server ,
		port	: port ,	
	}
}

func CreateConnection() *Connection {
	return &Connection {}
}

func OpenConnection(config *Config) (Connection, bool) {
	connectionString := config.token + " " + config.server + " " + config.port
	connection,ok := socketConnect(connectionString)
	connection.connectionString = connectionString
	if !ok {
		fmt.Println("nipolib Error connecting to socket: ")
	}
	return *connection,ok
}

func (connection *Connection) Logout() bool {
	connection.socket.Close()
	return true
}

func Set(config *Config, key string, value string) (string, bool) {
	connection,ok := OpenConnection(config)
	result := ""
	if ok {
		result,ok := connection.socketWrite(config.token + " set "+ key + " " + value)
		return result,ok	
	}
	connection.Logout()
	return result,ok
}

func Get(config *Config, key string) (string, bool) {
	connection,ok := OpenConnection(config)
	result := ""
	if ok {
		result,ok := connection.socketWrite(config.token + " get "+ key)
		return result,ok
	}
	connection.Logout()
	return result,ok
}

func Select(config *Config, key string) (string, bool) {
	connection,ok := OpenConnection(config)
	result := ""
	if ok {
		result,ok := connection.socketWrite(config.token + " select "+ key)
		return result,ok
	}
	connection.Logout()
	return result,ok
}

func Avg(config *Config, key string) (string, bool) {
	connection,ok := OpenConnection(config)
	result := ""
	if ok {
		result,ok := connection.socketWrite(config.token + " avg "+ key)
		return result,ok
	}
	connection.Logout()
	return result,ok
}