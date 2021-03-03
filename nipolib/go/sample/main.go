package main

import (
	"nipo"
	"fmt"
)

func main() {
	//connection := gonipolib.CreateConnection()
	//ok := false
	connection,ok := nipo.Login("admin admin 127.0.0.2 2323")
	if ok {
		fmt.Println("aaaaaaaaaaaa")
		setres,_ := connection.Set("name", "adas")
		fmt.Println(setres)
		getres,_ := connection.Get("name")
		fmt.Println(getres)
	}
	connection.Logout()
}
