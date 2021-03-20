package main

import (
    "os"
)

func main() {
    database := CreateDatabase()
    config := GetConfig(os.Args[1])
    cluster := config.CreateCluster()
    if config.Global.Master == "true" {
        go database.RunCluster(config, cluster)
    }
    database.Run(config, cluster)
}