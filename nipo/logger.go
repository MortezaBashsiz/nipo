package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
writes the given log into log file and stdout
*/
func (config *Config) logger(strLog string, level int) {
	file, err := os.OpenFile(config.Log.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	prefix := ""
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	if config.Log.Level >= level {
		if level == 1 {
			prefix = "INFO "
		}
		if level == 2 {
			prefix = "DEBUG "
		}
		fmt.Println(strings.TrimSuffix((prefix + strLog), "\n"))
		log.Println(strings.TrimSuffix((prefix + strLog), "\n"))
	}
}
