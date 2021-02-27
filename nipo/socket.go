package main

import (
	"net"
	"os"
	"fmt"
	"bufio"
	"strings"
	"encoding/json"
)

func Login(config *Config, connection net.Conn) (bool, *User) {
	user := CreateUser()
	strRemoteAddr := connection.RemoteAddr().String()
	connection.Write([]byte("Welcome to NIPO"+"\n"))
	connection.Write([]byte("You are connecting from "+strRemoteAddr+"\n"))
	connection.Write([]byte("Enter username : "))
	username, err := bufio.NewReader(connection).ReadString('\n')
	authorized := false
	if err != nil {
		fmt.Println(err)
		return false,nil
	}
	connection.Write([]byte("Enter password : "))
	password, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return false,nil
	}
	username = strings.TrimSuffix(username, "\n")
	password = strings.TrimSuffix(password, "\n")
	for _, tempuser := range config.Users {
		if username == tempuser.Username {
			if password == tempuser.Password {
				authorized = true
				user = tempuser
				return authorized,user
			} 
		} 
	}
	if authorized == false {
		connection.Write([]byte("nipo > wrong user or password"))
		connection.Write([]byte("\n"))
	}
	return authorized,user
}

func (database *Database) HandelSocket(config *Config, connection net.Conn, user *User) {
	defer connection.Close()
	strRemoteAddr := connection.RemoteAddr().String()
	for {
		connection.Write([]byte("nipo > "))
		input, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
				fmt.Println(err)
				return
		}
		if strings.TrimSpace(string(input)) == "exit" {
				config.logger("Client closed the connection from "+strRemoteAddr)
				return
		}
		if strings.TrimSpace(string(input)) == "EOF" {
			config.logger("Client terminated the connection from "+strRemoteAddr)
			return
		}
		returneddb,message := database.cmd(string(input), config, user)
		jsondb, err := json.Marshal(returneddb.items)
		if message != ""{
			connection.Write([]byte(message))
			connection.Write([]byte("\n"))
		}
		if err != nil {
			config.logger("Error in converting to json")
		}
		if len(jsondb) > 2 {
			connection.Write([]byte(message))
			connection.Write(jsondb)
			connection.Write([]byte("\n"))
		}
	}
}

func (database *Database) OpenSocket(config *Config) {
	config.logger("Opennig Socket on "+config.Listen.Ip+":"+config.Listen.Port+"/"+config.Listen.Protocol)
	socket,err := net.Listen(config.Listen.Protocol, config.Listen.Ip+":"+config.Listen.Port)
	if err != nil {
        config.logger("Error listening: "+err.Error())
		os.Exit(1)
	}
	defer socket.Close()
	for {
		connection, err := socket.Accept()
		strRemoteAddr := connection.RemoteAddr().String()
		if err != nil {
			config.logger("Error accepting socket: "+err.Error())
		}
		loginResult,user := Login(config,connection)
		if loginResult {
			connection.Write([]byte("Here you go with NIPO"+"\n"))
			go database.HandelSocket(config, connection, user)
		} else {
			config.logger("Wrong user pass from "+strRemoteAddr)
			connection.Close()
		}
	}
}

func ConnectSocket(config *Config) {
	socket,err := net.Dial(config.Listen.Protocol, config.Listen.Ip+":"+config.Listen.Port)
	if err != nil {
        config.logger("Error Connecting: "+err.Error())
		os.Exit(1)
	}
	defer socket.Close()
	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(socket)
 
	for {
		serverResponse, err := serverReader.ReadString('\n')
		if err != nil {
			config.logger("Server did not responsed to request")
		}
		socket.Write([]byte(serverResponse + "\n"))
		clientRequest, err := clientReader.ReadString('\n')
		if err != nil {
			config.logger("Client request was not correct")
		}
		socket.Write([]byte(clientRequest + "\n"))
	}
}