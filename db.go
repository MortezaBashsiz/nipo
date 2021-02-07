package main

type Db struct {
	items map [string] string
}

func CreateDB () *Db {
	return &Db{ items : map [string] string {} }
}

func (db *Db) Get(key string) (string, bool) {
	value,ok := db.items[key]
	return value,ok
}

func (db *Db) Set(key string, value string) (bool) {
	_,ok := db.Get(key)
	db.items[key] = value
	return ok
}

func (db *Db) Foreach (action func (string, string)) {
	for key,value := range db.items {
		action (key,value)
	}
}

func (db *Db) Accumulate (state interface{}, action func (interface{}, string, string) interface{}) interface{} {
	for key,value := range db.items {
		state = action (state, key, value)
	}
	return state
}