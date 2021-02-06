package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {  
	Prev  	*Item
	Next	*Item
	Key  	string
	Value	string
}

type Chain struct {  
	Head  	*Item
	Tail	*Item
}

func Display(item *Item) string {
	return item.Value
}

func (chain *Chain) Get(key string) *Item {
	current := chain.Head
	var item *Item
	for current != nil {
		if current.Key == key {
			item = current
			break
		} else {
			current = current.Next
			continue
		}
	}
	return item
}

func (chain *Chain) Set(key string, value string) {
	getItem := chain.Get(key)
	if getItem != nil {
		getItem.Value = value
	} else {
		item := &Item{
			Next : chain.Head,
			Key : key,
			Value : value,
		}
		if chain.Head != nil {
			chain.Head.Prev = item
		}
		chain.Head = item
	
		temp := chain.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		chain.Tail = temp
	}
}

func (chain *Chain) cmdCheck(cmd string) {
	cmdSplit := strings.Split(cmd," ")
	cmdType := ""
	argOne := ""
	argTwo := ""

	for _, split := range cmdSplit {
		if split == "set" {
			cmdType = "set"
			continue
		}
		if split == "get" {
			cmdType = "get"
			continue
		}
		if split != "set" && cmdType == "set" && argOne == "" && argTwo == "" {	
			argOne = split
			continue
		}
		if split != "set" && cmdType == "set" && argOne != "" && argTwo == "" {
			argTwo = split
			continue
		}

		if split != "get" && cmdType == "get" && argOne == "" {
			argOne = split
			continue
		}
	}
	
	if cmdType == "set" && argOne != "" && argTwo != "" {
		chain.Set(argOne,argTwo)
	}
	if cmdType == "get" && argOne != "" {
		fmt.Println(Display(chain.Get(argOne)))
	}
}

func main() {
	chain := Chain{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("nipo > ")
 		cmd, err := reader.ReadString('\n')
 		if err != nil {
  			fmt.Fprintln(os.Stderr, err)
 		}
 		cmd = strings.TrimSuffix(cmd, "\n")
 		chain.cmdCheck(cmd)
    }
}