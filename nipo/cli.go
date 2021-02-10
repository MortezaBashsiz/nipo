package main

import (
	"strings"
	"fmt"
)

func (database *Database) cmdCheck(cmd string) *Database {
	cmdSplit := strings.Split(cmd," ")
	cmdType := ""
	argOne := ""
    argTwo := ""
    db := CreateDatabase()
	for _, split := range cmdSplit {
        if split == "set" {
            cmdType = "set"
            continue
        }
        if split == "get" {
            cmdType = "get"
            continue
        }
        if split == "select" {
            cmdType = "select"
            continue
        }
        if split != "set" && cmdType == "set" && argOne == "" && argTwo == "" {
            argOne = split
            continue
        }
        if split != "set" && cmdType == "set" && argOne != "" && argTwo == "" {
            argTwo = split
            continue
        }
        if split != "get" && cmdType == "get" && argOne == "" {
            argOne = split
            continue
        }
        if split != "select" && cmdType == "select" && argOne == "" {
            argOne = split
            continue
        }
    }
    if cmdType == "set" && argOne != "" && argTwo != "" {
        ok := database.Set(argOne,argTwo)
        if ok {
            db.items[argOne] = argTwo
            return db
        }
    }
    if cmdType == "get" && argOne != "" {
		value,ok := database.Get(argOne)
		if ok {
            db.items[argOne] = value
            return db
		}
    }
    if cmdType == "select" && argOne != "" {
        db,err := database.Select(argOne)
        if err != nil {
            fmt.Println(err)
        }
        db.Foreach(func (key,value string) {
            db.items[key] = value
        })
        return db
    }
    return db
}