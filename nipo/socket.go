package main

import (
	// "sync"
	"net"
	"os"
	"fmt"
	"bufio"
	"strings"
	"encoding/json"
)

type Client struct {
	Connection net.Conn
	User User
	Authorized bool
}


func CreateClient() *Client {
	return &Client {}
}

func (client *Client)Login(config *Config) *Client {
	client.Connection.Write([]byte("nipo > "))
	logincmd, err := bufio.NewReader(client.Connection).ReadString('\n')
	client.Authorized = false
	loginCmdFields := strings.Fields(string(logincmd))
	if err != nil {
		fmt.Println(err)
		return client
	}
	if len(loginCmdFields) < 3 {
		fmt.Println("Error : login command needs 2 arg ")
		config.logger("Error : login command needs 2 arg ", 1)
		client.Connection.Write([]byte("Error : login command needs 2 arg \n"))
		return client
	}
	if loginCmdFields[0] != "login" {
		fmt.Println("Error : Wrong command : "+loginCmdFields[0])
		config.logger("Error : Wrong command : "+loginCmdFields[0], 1)
		client.Connection.Write([]byte("Error : Wrong command : "+loginCmdFields[0]+"\n"))
		return client
	}
	username := loginCmdFields[1]
	password := loginCmdFields[2]
	for _, tempuser := range config.Users {
		if username == tempuser.Username {
			if password == tempuser.Password {
				client.Authorized = true
				client.User = *tempuser
				client.Connection.Write([]byte("OK\n"))
				return client
			} 
		} 
	}
	if client.Authorized == false {
		client.Connection.Write([]byte("Error : wrong user or password \n"))
		client.Connection.Write([]byte("\n"))
	}
	return client
}

func (database *Database) HandelSocket(config *Config, client *Client) {
	defer client.Connection.Close()
	strRemoteAddr := client.Connection.RemoteAddr().String()
	client.Login(config)
	if client.Authorized {
		for {
			client.Connection.Write([]byte("nipo > "))
			input, err := bufio.NewReader(client.Connection).ReadString('\n')
			if err != nil {
					fmt.Println(err)
					return
			}
			if strings.TrimSpace(string(input)) == "exit" {
					config.logger("Client closed the connection from " + strRemoteAddr, 1)
					return
			}
			if strings.TrimSpace(string(input)) == "EOF" {
				config.logger("Client terminated the connection from " + strRemoteAddr, 2)
				return
			}
			returneddb,message := database.cmd(string(input), config, &client.User)
			jsondb, err := json.Marshal(returneddb.items)
			if message != ""{
				client.Connection.Write([]byte(message))
				client.Connection.Write([]byte("\n"))
			}
			if err != nil {
				config.logger("Error in converting to json" , 1)
			}
			if len(jsondb) > 2 {
				client.Connection.Write([]byte(message))
				client.Connection.Write(jsondb)
				client.Connection.Write([]byte("\n"))
			}
		}
	} else {
		config.logger("Wrong user pass from "+strRemoteAddr, 1)
		client.Connection.Close()
	}

}

func (database *Database) OpenSocket(config *Config) {
	config.logger("Opennig Socket on "+config.Listen.Ip+":"+config.Listen.Port+"/"+config.Listen.Protocol, 1)
	socket,err := net.Listen(config.Listen.Protocol, config.Listen.Ip+":"+config.Listen.Port)
	if err != nil {
        config.logger("Error listening: "+err.Error(), 2)
		os.Exit(1)
	}
	defer socket.Close()
	for {
		client := CreateClient()
		var err error
		client.Connection, err = socket.Accept()
		if err != nil {
			config.logger("Error accepting socket: "+err.Error(), 2)
		}
		go database.HandelSocket(config, client)
	}
}