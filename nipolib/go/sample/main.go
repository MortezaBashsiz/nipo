package main

import (
	"nipo"
	"runtime"
	"sync"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//connection := gonipolib.CreateConnection()
	//ok := false
	connection,ok := nipo.Login("admin admin 127.0.0.2 2323")
	runtime.GOMAXPROCS(2)
    var wg sync.WaitGroup
    wg.Add(1)
	go func() {
		if ok {
			for n := 0; n <= 50; n++ {
				fmt.Println(n)
				connection.Set(strconv.Itoa(n), strconv.Itoa(n))
				time.Sleep(1 * time.Microsecond)
			}
			// for n := 0; n <= 50; n++ {
			// 	connection.Get(strconv.Itoa(n))
			// }
		}
	}()
	wg.Wait()
	connection.Logout()
}
