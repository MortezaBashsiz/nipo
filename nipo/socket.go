package main

import (
	"net"
	"os"
	"fmt"
)

func (config *Config) OpenSocket() {
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
            config.logger("Error accepting: "+err.Error())
            os.Exit(1)
        }
        go handleRequest(connection)
    }
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 4096)
	reqLen, err := conn.Read(buf)
	if err != nil {
	  fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Adas: "+string(reqLen)+string(buf))
	conn.Write([]byte("Message received."))
	conn.Close()
}