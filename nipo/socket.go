package main

import (
	"net"
	"os"
	"fmt"
	"bufio"
	"strings"
)

func (database *Database) OpenSocket(config *Config) {
	config.logger("Opennig Socket on "+config.Listen.Ip+":"+config.Listen.Port+"/"+config.Listen.Protocol)
	socket,err := net.Listen(config.Listen.Protocol, config.Listen.Ip+":"+config.Listen.Port)
	if err != nil {
        config.logger("Error listening: "+err.Error())
		os.Exit(1)
    }
	defer socket.Close()
	connection, err := socket.Accept()
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