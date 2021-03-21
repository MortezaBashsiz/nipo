package main

import (
	"nipo"
	"time"
	"strconv"
)

type Slave struct {
	Node *Node
	Status, CheckedAt string
	Database *Database
}

type Cluster struct {
	Slaves []Slave
	Status string
}

func (cluster *Cluster) GetStatus() string {
	result := "{ "
	for index, slave := range cluster.Slaves {
		tempStr := "{ id : "+strconv.Itoa(slave.Node.Id) + " , ip : " + slave.Node.Ip + " , status : " + slave.Status + " , checkedat : " + slave.CheckedAt + " }"
		if !(index == len(cluster.Slaves)-1) {
			tempStr += ","
		}
		result += tempStr
	}
	result += " }"
	return result
}

func (config *Config) CreateCluster() *Cluster {
	cluster := Cluster {}
	for _, slave := range config.Slaves {
		tempSlave := Slave {}
		tempSlave.Node = slave
		tempSlave.Status = "none"
		tempSlave.CheckedAt = "none"
		tempSlave.Database = CreateDatabase()
		cluster.Slaves = append(cluster.Slaves, tempSlave)
	}
	cluster.Status = "none"
	return &cluster
}

func (cluster *Cluster) HealthCheck(config *Config) {
	for index, slave:= range cluster.Slaves {
		nipoconfig := nipo.CreateConfig(slave.Node.Token, slave.Node.Ip, slave.Node.Port)
		result,_ := nipo.Ping(nipoconfig) 
		if result == "pong\n" {
			if slave.Status == "unhealthy" {
				slave.Status = "recover"
				config.logger("slave by id : " + strconv.Itoa(slave.Node.Id) + " becomes healthy", 1)
				config.logger("slave by id : " + strconv.Itoa(slave.Node.Id) + " is in recovery", 1)
				nipoconfig := nipo.CreateConfig(slave.Node.Token, slave.Node.Ip, slave.Node.Port)
				slave.Database.Foreach(func (key,value string) {
					config.logger(key+value, 1)
					_, ok := nipo.Set(nipoconfig, key, value) 
					if !ok {
						config.logger("Set command on slave does not work correctly",2)
					}
				})
				slave.Database = CreateDatabase()
				config.logger("slave by id : " + strconv.Itoa(slave.Node.Id) + " recovery compleated", 1)
			}
			cluster.Slaves[index].Status = "healthy"
			cluster.Slaves[index].CheckedAt = time.Now().Format("2006-01-02 15:04:05.000")
			cluster.Status = "healthy"
		} else {
			cluster.Slaves[index].Status = "unhealthy"
			cluster.Slaves[index].CheckedAt = time.Now().Format("2006-01-02 15:04:05.000")
			cluster.Status = "unhealthy"
			config.logger("slave by id : " + strconv.Itoa(slave.Node.Id) + " is not healthy", 1)
		}
	}
	time.Sleep(time.Duration(config.Global.Checkinterval) * time.Millisecond)
}

func (cluster *Cluster) SetOnSlaves(config *Config,key,value string) {
	for _, slave:= range cluster.Slaves {
		if slave.Status == "healthy" {
			nipoconfig := nipo.CreateConfig(slave.Node.Token, slave.Node.Ip, slave.Node.Port)
			_, ok := nipo.Set(nipoconfig, key, value) 
			if !ok {
				config.logger("Set command on slave does not work correctly",2)
			}
		} 
		if slave.Status == "unhealthy" {
			slave.Database.Set(key,value)
		}
	}
}

func (database *Database) RunCluster(config *Config, cluster *Cluster) {
	for {
		cluster.HealthCheck(config)
	}
}
