package main

import (
    "os"
    "flag"
)

func main() {
    serverflags := flag.Bool("server", false, "-server CONFIG_PATH")
    cliflags := flag.Bool("cli", false, "-cli CONFIG_PATH")
    flag.Parse()
    database := CreateDatabase()
    config := GetConfig(os.Args[2])
    if *serverflags {
        database.OpenSocket(config);
    }
    if *cliflags {
        ConnectSocket(config)
    }
}