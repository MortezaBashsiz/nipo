package main

import (
    "os"
)

func main() {
    database := CreateDatabase()
    config := GetConfig(os.Args[1])
    database.OpenSocket(config);
}