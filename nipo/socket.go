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

func Login(config *Config, connection net.Conn) (bool, *User) {
	user := CreateUser()
	connection.Write([]byte("nipo > "))
	logincmd, err := bufio.NewReader(connection).ReadString('\n')
	authorized := false
	loginCmdFields := strings.Fields(string(logincmd))
	if err != nil {
		fmt.Println(err)
		return false,nil
	}
	if len(loginCmdFields) < 3 {
		fmt.Println("Error : login command needs 2 arg ")
		config.logger("Error : login command needs 2 arg ", 1)
		connection.Write([]byte("Error : login command needs 2 arg \n"))
		return false,nil
	}
	if loginCmdFields[0] != "login" {
		fmt.Println("Error : Wrong command : "+loginCmdFields[0])
		config.logger("Error : Wrong command : "+loginCmdFields[0], 1)
		connection.Write([]byte("Error : Wrong command : "+loginCmdFields[0]+"\n"))
		return false,nil
	}
	username := loginCmdFields[1]
	password := loginCmdFields[2]
	for _, tempuser := range config.Users {
		if username == tempuser.Username {
			if password == tempuser.Password {
				authorized = true
				user = tempuser
				connection.Write([]byte("OK\n"))
				return authorized,user
			} 
		} 
	}
	if authorized == false {
		connection.Write([]byte("Error : wrong user or password \n"))
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
				config.logger("Client closed the connection from " + strRemoteAddr, 1)
				return
		}
		if strings.TrimSpace(string(input)) == "EOF" {
			config.logger("Client terminated the connection from " + strRemoteAddr, 2)
			return
		}
		returneddb,message := database.cmd(string(input), config, user)
		jsondb, err := json.Marshal(returneddb.items)
		if message != ""{
			connection.Write([]byte(message))
			connection.Write([]byte("\n"))
		}
		if err != nil {
			config.logger("Error in converting to json" , 1)
		}
		if len(jsondb) > 2 {
			connection.Write([]byte(message))
			connection.Write(jsondb)
			connection.Write([]byte("\n"))
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
    var wg sync.WaitGroup
	for thread := 0; thread < config.Proc.Threads; thread++ {
		wg.Add(1)
		go func() {
        	defer wg.Done()
			for {
				connection, err := socket.Accept()
				strRemoteAddr := connection.RemoteAddr().String()
				if err != nil {
					config.logger("Error accepting socket: "+err.Error(), 2)
				}
				loginResult,user := Login(config,connection)
				if loginResult {
					database.HandelSocket(config, connection, user)
				} else {
					config.logger("Wrong user pass from "+strRemoteAddr, 1)
					connection.Close()
				}
			}
		}()
	}
	wg.Wait()
}