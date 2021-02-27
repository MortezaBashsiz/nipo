package main

import (
    // "os"
    "flag"
	"net"
	"fmt"
	// "bufio"
	// "strings"
)

func main() {
    cliFlagsServer := flag.String("s", "127.0.0.1:2323", "IP:PORT")
	cliFlagsUser := flag.String("u", "admin", "username")
	cliFlagsPass := flag.String("p", "admin", "password")
	response := make([]byte, 4096)
    flag.Parse()
    connection, err := net.Dial("tcp", *cliFlagsServer)
	defer connection.Close()
	if err != nil {
		fmt.Println(err)
	}
	connection.Write([]byte(*cliFlagsUser))
	connection.Write([]byte(*cliFlagsPass))
	for {
		connection.Read(response)
		fmt.Println(string(response))
	}
}