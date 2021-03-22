package main

import (
    "os"
)

func main() {
    database := CreateDatabase()
    config := GetConfig(os.Args[1])
    cluster := config.CreateCluster()
    database.Run(config, cluster)
}