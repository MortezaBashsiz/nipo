package main

import (
	"nipo"
	// "runtime"
	// "sync"
	"fmt"
	"strconv"
	"os"
)

func main() {
	// runtime.GOMAXPROCS(10)
    // var wg sync.WaitGroup
	if os.Args[1] == "set" {
		for n:=0 ; n < 10 ; n++ {
    		// wg.Add(1)
			// go func() {
				connection,ok := nipo.Login("admin admin 127.0.0.1 2323")
				if ok {
					for n := 0; n <= 1000000; n++ {
						connection.Set(strconv.Itoa(n), strconv.Itoa(n))
					}
					result,_ := connection.Avg(".*")
					fmt.Println(result)
				}
			// }()
		}
	}
	if os.Args[1] == "get" {
		// for n:=0 ; n < 10 ; n++ {
    		// wg.Add(1)
			// go func() {
				connection,ok := nipo.Login("admin admin 192.168.100.194 2323")
				if ok {
					for n := 0; n <= 1000000; n++ {
						connection.Get(strconv.Itoa(n))
					}
					result,_ := connection.Avg(".*")
					fmt.Println(result)
				}
			// }()
		// }
	}
	// wg.Wait()
	// connection.Logout()
}
