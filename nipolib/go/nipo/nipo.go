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
		fmt.Println("nipolib Error connecting to socket: " + connectionString)
	}
	return *connection,ok
}

func (connection *Connection) Logout() bool {
	err := connection.socket.Close()
	if err != nil {
		fmt.Println("nipolib Error logout from connection : " + err.Error())
		return false	
	}
	return true
}

func Ping(config *Config) (string, bool) {
	connection,ok := OpenConnection(config)
	result := ""
	if ok {
		result,ok := connection.socketWrite(config.token + " ping ")
		connection.Logout()
		return result,ok	
	} 
	return result,ok
}

func Status(config *Config) (string, bool) {
	connection,ok := OpenConnection(config)
	result := ""
	if ok {
		result,ok := connection.socketWrite(config.token + " status ")
		connection.Logout()
		return result,ok	
	} 
	return result,ok
}

func Set(config *Config, key string, value string) (string, bool) {
	connection,ok := OpenConnection(config)
	result := ""
	if ok {
		result,ok := connection.socketWrite(config.token + " set "+ key + " " + value)
		connection.Logout()
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

func Sum(config *Config, key string) (string, bool) {
	connection,ok := OpenConnection(config)
	result := ""
	if ok {
		result,ok := connection.socketWrite(config.token + " sum "+ key)
		return result,ok
	}
	connection.Logout()
	return result,ok
}

func Count(config *Config, key string) (string, bool) {
	connection,ok := OpenConnection(config)
	result := ""
	if ok {
		result,ok := connection.socketWrite(config.token + " count "+ key)
		return result,ok
	}
	connection.Logout()
	return result,ok
}