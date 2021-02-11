package main

import (
	"net"
	"os"
	"fmt"
	"bufio"
	"strings"
)

func (database *Database) HandelSocket(connection net.Conn) {
	defer connection.Close()
	for {
		connection.Write([]byte("nipo > "))
		input, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
				fmt.Println(err)
				return
		}
		if strings.TrimSpace(string(input)) == "exit" {
				fmt.Println("Exiting TCP server!")
				return
		}
		returneddb := database.cmd(string(input))
        returneddb.Foreach(func (key,value string) {
            connection.Write([]byte("> "+key+" "+value))
		})
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
		if err != nil {
			config.logger("Error accepting socket: "+err.Error())
		}
		go database.HandelSocket(connection)
	}
}