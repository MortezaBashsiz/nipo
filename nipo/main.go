package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "flag"
)

func main() {
    serverflags := flag.Bool("server", false, "-server CONFIG_PATH")
    cliflags := flag.Bool("cli", false, "-cli CONFIG_PATH")
    flag.Parse()
    database := CreateDatabase()
    returneddb := CreateDatabase()
    if *serverflags {
        config := GetConfig(os.Args[2])
        database.OpenSocket(config);
    }
    if *cliflags {
        reader := bufio.NewReader(os.Stdin)
        for {
            fmt.Print("nipo > ")
            cmd, err := reader.ReadString('\n')
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
            }
            cmd = strings.TrimSuffix(cmd, "\n")
            returneddb = database.cmdCheck(cmd)
            returneddb.Foreach(func (key,value string) {
                fmt.Println(key,value)
            })
        }
    }
}