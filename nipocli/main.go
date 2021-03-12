package main

import (
	"flag"
	"nipo"
)

func main(){
	Token := flag.String("t", "token", "Token if authorization is enabled")
	Server := flag.String("s", "127.0.0.1", "Server IP or Hostname")
	Port := flag.String("p", "2323", "Server port number")
	flag.Parse()
	config := nipo.CreateConfig(*Token, *Server, *Port)
	Start(config)
}