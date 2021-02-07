package main

import (
	"fmt"
	"strconv"
)

func main() {
	db := CreateDB()

	for n := 0; n < 10; n++ {
		db.Set(strconv.Itoa(n),strconv.Itoa(n))
	}
	db.Set("a","b")
	db.Foreach(func (key,_ string) {
		fmt.Println(key)
	})

	defer func () {
		if r := recover ();r != nil {
			fmt.Println("Pannnnnnnnnnnnnnnnnnnniced", r)
		}
	}()

	sum := db.Accumulate(0, func (state interface{}, key,value string) interface{} {
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