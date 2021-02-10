package main

import (
	"regexp"
)

type Database struct {
	items map [string] string
}

func CreateDatabase() *Database {
	return &Database{ items : map [string] string {} }
}

func (database *Database) Get(key string) (string, bool) {
	value,ok := database.items[key]
	return value,ok
}

func (database *Database) Set(key string, value string) (bool) {
	_,ok := database.Get(key)
	database.items[key] = value
	return ok
}

func (database *Database) Foreach(action func (string, string)) {
	for key,value := range database.items {
		action (key,value)
	}
}

func (database *Database) Select(keyregex string) (*Database, error) {
	selected := CreateDatabase()
	var err error
	for key,value := range database.items {
		matched,err := regexp.MatchString(keyregex, key)
		if err != nil {
			return selected,err
		}
		if matched {
			selected.items[key] = value
		}
	}
	return selected,err
}

func (database *Database) Accumulate(state interface{}, action func (interface{}, string, string) interface{}) interface{} {
	for key,value := range database.items {
		state = action (state, key, value)
	}
	return state
}