package main

import (
	"nipo"
	// "runtime"
	// "sync"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//connection := gonipolib.CreateConnection()
	//ok := false
	connection,ok := nipo.Login("admin admin 127.0.0.1 2323")
	time.Sleep(10000 * time.Microsecond)
	// runtime.GOMAXPROCS(16)
    // var wg sync.WaitGroup
	// for n:=0 ; n <= 16 ; n++ {
    	// wg.Add(1)
		// go func() {
	fmt.Println("---------------------------")
			if ok {
				for n := 0; n <= 5; n++ {
					connection.Set(strconv.Itoa(n), strconv.Itoa(n))
					// fmt.Println(result)
					// fmt.Println(err)
					// time.Sleep(1000 * time.Microsecond)
				}
				// for n := 0; n <= 50; n++ {
				// 	connection.Get(strconv.Itoa(n))
				// }
			}
		// }()
	// }
	// wg.Wait()
	connection.Logout()
}
