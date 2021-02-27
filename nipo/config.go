package main

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type User struct {
	Username string
	Password string
	Keys string
	Cmds string
}

type Config struct {  

	Listen struct {
		Ip		string
		Port	string
		Protocol string
	}

	Log struct {
		Level 	string
		Path  	string
	}

	Users []*User

}

func CreateConfig() *Config {
	return &Config {}
}

func GetConfig(path string) *Config {  
	file, err := os.Open(path)
	config := CreateConfig()
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()
	
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
	    log.Fatal(err)
	}
	config.logger("Nipo service is starting ...")
	config.logger("Reading config file on :" +path)
	return config
}