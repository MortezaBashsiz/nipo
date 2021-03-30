package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
validates that user have permission to run the command or not
*/
func validateCmd(cmd string, user *User) bool {
	cmds := strings.Split(user.Cmds, "||")
	allowed := false
	for count := range cmds {
		if cmds[count] == "all" {
			return true
		}
		if cmds[count] == cmd {
			return true
		}
	}
	return allowed
}

/*
validates that user have access to this key (regex) or not
*/
func validateKey(key string, user *User) bool {
	keys := strings.Split(user.Keys, "||")
	allowed := false
	for count := range keys {
		matched, err := regexp.MatchString(keys[count], key)
		if err != nil {
			fmt.Println(err)
		}
		if matched {
			return true
		}
	}
	return allowed
}

/*
sets the key and value into database
*/
func (database *Database) cmdSet(config *Config, cluster *Cluster, cmd string) (*Database, bool) {
	cmdFields := strings.Fields(cmd)
	key := cmdFields[1]
	db := CreateDatabase()
	ok := false
	if len(cmdFields) >= 3 {
		value := cmdFields[2]
		for n := 3; n < len(cmdFields); n++ {
			value += " " + cmdFields[n]
		}
		cluster.SetOnSlaves(config, key, value)
		ok = database.Set(key, value)
		if !ok {
			return db, false
		} else {
			db.items[key] = value
			return db, true
		}
	}
	return db, ok
}

/*
gets the value of given keys, it will create a new databse as result
*/
func (database *Database) cmdGet(cmd string) *Database {
	cmdFields := strings.Fields(cmd)
	db := CreateDatabase()
	for _, key := range cmdFields {
		if cmdFields[0] != key {
			value, ok := database.Get(key)
			if ok {
				db.items[key] = value
			}
		}
	}
	return db
}

/*
selects the keys and values which are matched in given regex
it will create a new databse as result
*/
func (database *Database) cmdSelect(cmd string) *Database {
	cmdFields := strings.Fields(cmd)
	key := cmdFields[1]
	db, err := database.Select(key)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

/*
summeries the values which mathed in given regex
it will ignore non number values
*/
func (database *Database) cmdSum(cmd string) *Database {
	cmdFields := strings.Fields(cmd)
	key := cmdFields[1]
	db, err := database.Select(key)
	returndb := CreateDatabase()
	var sum float64 = 0
	if err != nil {
		fmt.Println(err)
	}
	db.Foreach(func(key, value string) {
		valFloat, _ := strconv.ParseFloat(value, 64)
		sum += valFloat
	})
	returndb.items[key] = fmt.Sprintf("%f", sum)
	return returndb
}

/*
calculate the average of the values which mathed in given regex
it will ignore non number values
*/
func (database *Database) cmdAvg(cmd string) *Database {
	cmdFields := strings.Fields(cmd)
	key := cmdFields[1]
	db, err := database.Select(key)
	returndb := CreateDatabase()
	var sum float64 = 0
	count := 0
	if err != nil {
		fmt.Println(err)
	}
	db.Foreach(func(key, value string) {
		valFloat, _ := strconv.ParseFloat(value, 64)
		sum += valFloat
		count++
	})
	avg := (float64(sum)) / (float64(count))
	returndb.items[key] = fmt.Sprintf("%f", avg)
	return returndb
}

/*
counts the keys which mathed in given regex
it will ignore non number values
*/
func (database *Database) cmdCount(cmd string) *Database {
	cmdFields := strings.Fields(cmd)
	key := cmdFields[1]
	db, err := database.Select(key)
	returndb := CreateDatabase()
	// var sum float64 = 0
	count := 0
	if err != nil {
		fmt.Println(err)
	}
	db.Foreach(func(key, value string) {
		count++
	})
	returndb.items["count"] = strconv.Itoa(count)
	return returndb
}

/*
the main function to handle the command
checks the validation and autorization of user to access the keys and commands
*/
func (database *Database) cmd(cmd string, config *Config, cluster *Cluster, user *User) (*Database, string) {
	config.logger("client executed command : "+cmd, 2)
	config.logger("cmd.go - func cmd - with cmd : "+cmd, 2)
	config.logger("cmd.go - func cmd - with user : "+user.Name, 2)
	cmdFields := strings.Fields(cmd)
	db := CreateDatabase()
	ok := false
	message := ""
	if len(cmdFields) >= 2 {
		switch cmdFields[0] {
		case "count":
			db = database.cmdCount(cmd)
			break
		case "set":
			if config.Global.Authorization == "true" {
				if validateCmd("set", user) {
					if validateKey(cmdFields[1], user) {
						db, ok = database.cmdSet(config, cluster, cmd)
						if !ok {
							message = ("set failed by user " + user.Name + " for command : " + cmd)
							config.logger(message, 1)
						}
					} else {
						message = ("User " + user.Name + " not allowed to use regex : " + cmdFields[1])
						config.logger(message, 1)
					}
				} else {
					message = ("User " + user.Name + " not allowed to use command : " + cmd)
					config.logger(message, 1)
				}
			} else {
				db, ok = database.cmdSet(config, cluster, cmd)
				if !ok {
					message = ("set failed by user " + user.Name + " for command : " + cmd)
					config.logger(message, 1)
				}
			}
			break
		case "get":
			if config.Global.Authorization == "true" {
				if validateCmd("get", user) {
					if validateKey(cmdFields[1], user) {
						db = database.cmdGet(cmd)
					} else {
						message = ("User " + user.Name + " not allowed to use regex : " + cmdFields[1])
						config.logger(message, 1)
					}
				} else {
					message = ("User " + user.Name + " not allowed to use command : " + cmd)
					config.logger(message, 1)
				}
			} else {
				db = database.cmdGet(cmd)
			}
			break
		case "select":
			if config.Global.Authorization == "true" {
				if validateCmd("select", user) {
					if validateKey(cmdFields[1], user) {
						db = database.cmdSelect(cmd)
					} else {
						message = ("User " + user.Name + " not allowed to use regex : " + cmdFields[1])
						config.logger(message, 1)
					}
				} else {
					message = ("User " + user.Name + " not allowed to use command : " + cmd)
					config.logger(message, 1)
				}
			} else {
				db = database.cmdSelect(cmd)
			}
			break
		case "sum":
			if config.Global.Authorization == "true" {
				if validateCmd("sum", user) {
					if validateKey(cmdFields[1], user) {
						db = database.cmdSum(cmd)
					} else {
						message = ("User " + user.Name + " not allowed to use regex : " + cmdFields[1])
						config.logger(message, 1)
					}
				} else {
					message = ("User " + user.Name + " not allowed to use command : " + cmd)
					config.logger(message, 1)
				}
			} else {
				db = database.cmdSum(cmd)
			}
			break
		case "avg":
			if config.Global.Authorization == "true" {
				if validateCmd("avg", user) {
					if validateKey(cmdFields[1], user) {
						db = database.cmdAvg(cmd)
					} else {
						message = ("User " + user.Name + " not allowed to use regex : " + cmdFields[1])
						config.logger(message, 1)
					}
				} else {
					message = ("User " + user.Name + " not allowed to use command : " + cmd)
					config.logger(message, 1)
				}
			} else {
				db = database.cmdAvg(cmd)
			}
			break
		}
	}
	return db, message
}
