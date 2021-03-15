package main

import (
	"nipo"
)

func (databse *Database) SyncSlave(config *Config) bool{
	return true
}

func SetOnSlaves(config *Config,key,value string) bool {
	for _, slave:= range config.Slaves {
		nipoconfig := nipo.CreateConfig(slave.Token, slave.Ip, slave.Port)
		nipo.Set(nipoconfig, key, value) 
	}
	return true
}
