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
    clientflags := flag.Bool("client", true, "-client CONFIG_PATH")
    flag.Parse()
    database := CreateDB()
    if *serverflags {
        config := GetConfig(os.Args[2])
        config.OpenSocket();
    }
    if *clientflags {
        reader := bufio.NewReader(os.Stdin)
        for {
            fmt.Print("nipo > ")
            cmd, err := reader.ReadString('\n')
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
            }
            cmd = strings.TrimSuffix(cmd, "\n")
            database.cmdCheck(cmd)
        }
    }
}