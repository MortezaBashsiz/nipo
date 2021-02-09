package main

import (
    "log"
	"os"
)

func (config *Config) logger(strLog string) {
	file, err := os.OpenFile(config.Log.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(file)
    log.Println(strLog)
}
