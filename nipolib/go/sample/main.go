package main

import (
	"nipo"
	// "runtime"
	// "sync"
	"fmt"
	"strconv"
	"os"
	// "time"
)

func main() {
	max,_ := strconv.Atoi(os.Args[2])
	if os.Args[1] == "set" {
		for n:=0 ; n <= max ; n++ {
			connection,_ := nipo.OpenConnection(os.Args[3]+" 127.0.0.1 2323")
			connection.Set(os.Args[3], strconv.Itoa(n), strconv.Itoa(n))
			connection.Logout()
		}
		connection,_ := nipo.OpenConnection(os.Args[3]+" 127.0.0.1 2323")
		result,_ := connection.Avg(os.Args[3], ".*")
		fmt.Println(result)
		connection.Logout()
	}
	if os.Args[1] == "get" {
		for n:=0 ; n <= max ; n++ {
			connection,_ := nipo.OpenConnection(os.Args[3]+" 127.0.0.1 2323")
			connection.Get(os.Args[3], strconv.Itoa(n))
			connection.Logout()
		}
		connection,_ := nipo.OpenConnection(os.Args[3]+" 127.0.0.1 2323")
		result,_ := connection.Avg(os.Args[3], ".*")
		fmt.Println(result)
		connection.Logout()
	}
}
