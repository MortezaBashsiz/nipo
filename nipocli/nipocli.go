package main

import (
	"fmt"
	"bufio"
	"os"
	"nipo"
	"strings"
)

func checkCmd(cmd string, config *nipo.Config) string {
	cmdFields := strings.Fields(cmd)
    result := ""
    if len(cmdFields) >= 2 {
        switch cmdFields[0] {
        case "set":
			value := ""
			for count:=2 ; count < len(cmdFields); count++ {
				value += cmdFields[count]+" "
			}
            result,_ = nipo.Set(config, cmdFields[1], value)
            break
        case "get":
			keys := ""
			for count:=1 ; count < len(cmdFields); count++ {
				keys += cmdFields[count]+" "
			}
            result,_ = nipo.Get(config, keys)
            break
		case "sum":
            result,_ = nipo.Sum(config, cmdFields[1])
			break
        case "select":
            result,_ = nipo.Select(config, cmdFields[1])
            break
        case "avg":
            result,_ = nipo.Avg(config, cmdFields[1])
        }
    } 
	return result
}

func Start(config *nipo.Config) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("nipo > ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		result := checkCmd(cmd, config)
		result = strings.TrimSuffix(result, "\n")
		fmt.Println(result)
	}
}