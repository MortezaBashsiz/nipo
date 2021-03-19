package main

import (
    "os"
)

func main() {
    database := CreateDatabase()
    config := GetConfig(os.Args[1])
    if config.Global.Master == "true" {
        go database.RunCluster(config)
    }
    database.Run(config)
}