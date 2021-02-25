package main

import (
	"strings"
	"fmt"
    "strconv"
    "regexp"
)

func validateKey(key string, config *Config) bool {
    keyWildCard := config.Access.Wildcard
    keyMatched, _ := regexp.MatchString(keyWildCard, key)
    return keyMatched
}

func validateCmd(cmd string, config *Config) bool {
    cmdWildCard := config.Access.Wildcard
    cmdMatched, _ := regexp.MatchString(cmdWildCard, cmd)
    return cmdMatched
}

func (database *Database) cmdSet(cmd string) *Database {
    cmdFields := strings.Fields(cmd)
    key := cmdFields[1]
    db := CreateDatabase() 
    if len(cmdFields) >= 3 {
        value := cmdFields[2]
        for n:=3; n<len(cmdFields); n++ {
            value += " "+cmdFields[n]
        }   
        database.Set(key,value)
        db.items[key] = value
    }
    return db
}

func (database *Database) cmdGet(cmd string) *Database {
    cmdFields := strings.Fields(cmd)
    db := CreateDatabase()
    for _, key := range cmdFields {
        if cmdFields[0] != key {
            value,ok := database.Get(key)
            if ok {
                db.items[key] = value
            }
        }
    }
    return db
}

func (database *Database) cmdSelect(cmd string) *Database {
    cmdFields := strings.Fields(cmd)
    key := cmdFields[1]
    db,err := database.Select(key)
    if err != nil {
        fmt.Println(err)
    }
    return db
}

func (database *Database) cmdSum(cmd string) *Database {
    cmdFields := strings.Fields(cmd)
    key := cmdFields[1]
    db,err := database.Select(key)
    returndb := CreateDatabase()
    var sum float64 = 0
    if err != nil {
        fmt.Println(err)
    }
    db.Foreach(func (key,value string) {
        valFloat,_ :=  strconv.ParseFloat(value, 64)
        sum += valFloat
    })
    returndb.items[key] = fmt.Sprintf("%f", sum)
    return returndb
}

func (database *Database) cmdAvg(cmd string) *Database {
    cmdFields := strings.Fields(cmd)
    key := cmdFields[1]
    db,err := database.Select(key)
    returndb := CreateDatabase()
    var sum float64 = 0
    count := 0
    if err != nil {
        fmt.Println(err)
    }
    db.Foreach(func (key,value string) {
        valFloat,_ :=  strconv.ParseFloat(value, 64)
        sum += valFloat
        count ++
    })
    avg := (float64(sum))/(float64(count))
    returndb.items[key] = fmt.Sprintf("%f", avg)
    return returndb
}

func (database *Database) cmd(cmd string, config *Config) *Database {
    config.logger("client executed command : "+cmd)
    cmdFields := strings.Fields(cmd)
    db := CreateDatabase()
    if len(cmdFields) >= 2 {
        switch cmdFields[0] {
        case "set":
            db = database.cmdSet(cmd)
            break
        case "get":
            db = database.cmdGet(cmd)
            break
        case "select":
            db = database.cmdSelect(cmd)
            break
        case "sum":
            db = database.cmdSum(cmd)
            break
        case "avg":
            db = database.cmdAvg(cmd)
            break
        }
    }
    return db
}