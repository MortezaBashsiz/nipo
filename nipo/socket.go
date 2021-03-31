/*
Written by Morteza Bashsiz <morteza.bashsiz@gmail.com>
This file contains all functions and objects related to handling network, sockets,
	multi processing and multi threading
*/

package main

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
)

var Wait sync.WaitGroup
var Lock sync.Mutex

type Client struct {
	Connection net.Conn
	User       User
	Authorized bool
}

func CreateClient() *Client {
	return &Client{}
}

/*
validate the client with given token
*/
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

/*
handles opened socket. after checking the authorization field at config, validates
given token, checks the command fields count, executes the command, converts to json
and finally writes on opened socket
*/
func (database *Database) HandleSocket(config *Config, cluster *Cluster, client *Client) {
	defer client.Connection.Close()
	strRemoteAddr := client.Connection.RemoteAddr().String()
	input, err := bufio.NewReader(client.Connection).ReadString('\n')
	if err != nil {
		config.logger("Read from socket error : "+err.Error(), 2)
		return
	}
	inputFields := strings.Fields(input)
	if len(inputFields) >= 2 {
		if inputFields[1] == "ping" {
			_, _ = client.Connection.Write([]byte("pong"))
			_, _ = client.Connection.Write([]byte("\n"))
			return
		}
		if inputFields[1] == "status" {
			status := ""
			if config.Global.Master == "true" {
				status = cluster.GetStatus()
			} else {
				status = "Not Clustered"
			}
			_, _ = client.Connection.Write([]byte(status))
			_, _ = client.Connection.Write([]byte("\n"))
			return
		}
		if inputFields[1] == "exit" {
			config.logger("Client closed the connection from "+strRemoteAddr, 2)
			return
		}
		if inputFields[1] == "EOF" {
			config.logger("Client terminated the connection from "+strRemoteAddr, 2)
			return
		}
	}
	if config.Global.Authorization == "true" {
		if client.Validate(inputFields[0], config) {
			cmd := ""
			if len(inputFields) >= 3 {
				cmd = inputFields[1]
				for n := 2; n < len(inputFields); n++ {
					cmd += " " + inputFields[n]
				}
			}
			returneddb, message := database.cmd(cmd, config, cluster, &client.User)
			jsondb, err := json.Marshal(returneddb.items)
			if message != "" {
				_, _ = client.Connection.Write([]byte(message))
				_, _ = client.Connection.Write([]byte("\n"))
			}
			if err != nil {
				config.logger("Error in converting to json : "+err.Error(), 1)
			}
			if len(jsondb) > 2 {
				_, _ = client.Connection.Write([]byte(message))
				_, _ = client.Connection.Write(jsondb)
				_, _ = client.Connection.Write([]byte("\n"))
			}
		} else {
			config.logger("Wrong token "+strRemoteAddr, 1)
			client.Connection.Close()
		}
	} else {
		cmd := ""
		if len(inputFields) >= 3 {
			cmd = inputFields[1]
			for n := 2; n < len(inputFields); n++ {
				cmd += " " + inputFields[n]
			}
		}
		returneddb, message := database.cmd(cmd, config, cluster, &client.User)
		jsondb, err := json.Marshal(returneddb.items)
		if message != "" {
			_, _ = client.Connection.Write([]byte(message))
			_, _ = client.Connection.Write([]byte("\n"))
		}
		if err != nil {
			config.logger("Error in converting to json : "+err.Error(), 1)
		}
		if len(jsondb) > 2 {
			_, _ = client.Connection.Write([]byte(message))
			_, _ = client.Connection.Write(jsondb)
			_, _ = client.Connection.Write([]byte("\n"))
		}
	}
}

/*
called from main function, runs the service, multi-thread and multi-process handles here
calls the HandleSocket function
*/
func (database *Database) Run(config *Config, cluster *Cluster) {
	if config.Global.Master == "true" {
		go database.RunCluster(config, cluster)
	}
	config.logger("Opennig Socket on "+config.Listen.Ip+":"+config.Listen.Port+"/"+config.Listen.Protocol, 1)
	socket, err := net.Listen(config.Listen.Protocol, config.Listen.Ip+":"+config.Listen.Port)
	if err != nil {
		config.logger("Error listening: "+err.Error(), 1)
		os.Exit(1)
	}
	defer socket.Close()
	runtime.GOMAXPROCS(config.Proc.Cores)
	for thread := 0; thread < config.Proc.Threads; thread++ {
		Wait.Add(1)
		go func() {
			defer Wait.Done()
			for {
				client := CreateClient()
				var err error
				client.Connection, err = socket.Accept()
				if err != nil {
					config.logger("Error accepting socket : "+err.Error(), 2)
				}
				Lock.Lock()
				database.HandleSocket(config, cluster, client)
				Lock.Unlock()
			}
		}()
		Wait.Wait()
	}
}
