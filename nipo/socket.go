package main

import (
	"runtime"
	"sync"
	"net"
	"os"
	"fmt"
	"bufio"
	"strings"
	"encoding/json"
)

var Wait sync.WaitGroup
var Lock sync.Mutex

type Client struct {
	Connection net.Conn
	User User
	Authorized bool
}


func CreateClient() *Client {
	return &Client {}
}

func (client *Client) Validate(token string, config *Config) bool {
	client.Authorized = false
	for _, tempuser := range config.Users {
		if token == tempuser.Token {
			client.Authorized = true
			client.User = *tempuser
			return client.Authorized
		} 
	}
	return client.Authorized
}

func (database *Database) HandelSocket(config *Config, client *Client) {
	defer client.Connection.Close()
	strRemoteAddr := client.Connection.RemoteAddr().String()
	client.Connection.Write([]byte("nipo > "))
	input, err := bufio.NewReader(client.Connection).ReadString('\n')
	if err != nil {
			fmt.Println(err)
			return
	}
	inputFields := strings.Fields(string(input))
	if inputFields[0] == "exit" {
		config.logger("Client closed the connection from " + strRemoteAddr, 2)
		return
	}
	if inputFields[0] == "EOF" {
		config.logger("Client terminated the connection from " + strRemoteAddr, 2)
		return
	}

	if config.Global.Authorization == "true" {
		if client.Validate(inputFields[0], config) {
			cmd := ""
			if len(inputFields) >= 3 {
				cmd = inputFields[1]
				for n:=2; n<len(inputFields); n++ {
					cmd += " "+inputFields[n]
				}   
			}
			returneddb,message := database.cmd(cmd, config, &client.User, true)
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
		} else {
			config.logger("Wrong token "+strRemoteAddr, 1)
			client.Connection.Close()
		}
	} else {
		cmd := ""
		if len(inputFields) >= 3 {
			cmd = inputFields[1]
			for n:=2; n<len(inputFields); n++ {
				cmd += " "+inputFields[n]
			}   
		}
		returneddb,message := database.cmd(cmd, config, &client.User, false)
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
}

func (database *Database) OpenSocket(config *Config) {
	config.logger("Opennig Socket on "+config.Listen.Ip+":"+config.Listen.Port+"/"+config.Listen.Protocol, 1)
	socket,err := net.Listen(config.Listen.Protocol, config.Listen.Ip+":"+config.Listen.Port)
	if err != nil {
        config.logger("Error listening: "+err.Error(), 2)
		os.Exit(1)
	}
	defer socket.Close()
	runtime.GOMAXPROCS(config.Proc.Cores)
	for thread := 0; thread < config.Proc.Threads; thread++ {
		Wait.Add(1)
		go func() {
        	defer Wait.Done()
			for {
				Lock.Lock()
				client := CreateClient()
				var err error
				client.Connection, err = socket.Accept()
				if err != nil {
					config.logger("Error accepting socket: "+err.Error(), 2)
				}
				database.HandelSocket(config, client)
				Lock.Unlock()
			}
		}()
	}
	Wait.Wait()
}