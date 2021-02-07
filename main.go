package main

import (
	"fmt"
	"strconv"
)

func main() {
	database := CreateDB()

	for n := 0; n < 10000000; n++ {
		database.Set(strconv.Itoa(n),strconv.Itoa(n))
	}
	database.Set("a","b")
	database.Foreach(func (key,_ string) {
		fmt.Println(key)
	})

	defer func () {
		if r := recover ();r != nil {
			fmt.Println("Paniced", r)
		}
	}()

	sum := database.Accumulate(0, func (state interface{}, key,value string) interface{} {
		n := state.(int)
		if v, err := strconv.Atoi(value); err == nil {
			n += v 
		} else {
			panic (err)
		}
		return n
	}).(int)

	fmt.Printf("%d\n",sum)
}