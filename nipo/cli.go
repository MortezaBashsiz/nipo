package main

import (
	"strings"
	"fmt"
)

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
    key := cmdFields[1]
    db := CreateDatabase()
    value,ok := database.Get(key)
    if ok {
        db.items[key] = value
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
    db.Foreach(func (key,value string) {
        db.items[key] = value
    })
    return db
}

func (database *Database) cmd(cmd string) *Database {
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
        }
    }
    return db
}