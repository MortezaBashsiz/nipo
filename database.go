package main

type Database struct {
	items map [string] string
}

func CreateDB () *Database {
	return &Database{ items : map [string] string {} }
}

func (db *Database) Get(key string) (string, bool) {
	value,ok := db.items[key]
	return value,ok
}

func (db *Database) Set(key string, value string) (bool) {
	_,ok := db.Get(key)
	db.items[key] = value
	return ok
}

func (db *Database) Foreach (action func (string, string)) {
	for key,value := range db.items {
		action (key,value)
	}
}

func (db *Database) Accumulate (state interface{}, action func (interface{}, string, string) interface{}) interface{} {
	for key,value := range db.items {
		state = action (state, key, value)
	}
	return state
}