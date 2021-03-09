package main

import (
    "os"
)

func main() {
    database := CreateDatabase()
    config := GetConfig(os.Args[1])
    if config.Global.Authorization == "true" {
        database.OpenSocketWithAutorization(config);
    } else {
        database.OpenSocketWithoutAutorization(config);
    }
}