package main

import (
	"nipo"
	// "fmt"
	"strconv"
)

func main() {
	//connection := gonipolib.CreateConnection()
	//ok := false
	connection,ok := nipo.Login("admin admin 127.0.0.2 2323")
	if ok {
		for n := 0; n <= 500000; n++ {
			go connection.Set(strconv.Itoa(n), strconv.Itoa(n))
		}
		for n := 0; n <= 500000; n++ {
			go connection.Get(strconv.Itoa(n))
		}
	}
	connection.Logout()
}
