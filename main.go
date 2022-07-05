package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	// Create leveldb folder
	db, err := leveldb.OpenFile("db", nil)
	if err != nil {
		_ = fmt.Errorf("error opening db: %v", err)
	}
	defer db.Close()

	// modify database
	if err = db.Put([]byte("key"), []byte("hello world"), nil); err != nil {
		_ = fmt.Errorf("error inserting key: %v", err)
	}

	// modifu database
	if err = db.Put([]byte("key"), []byte("hello and bye"), nil); err != nil {
		_ = fmt.Errorf("error inserting key: %v", err)
	}

	// get key from database
	value, err := db.Get([]byte("key"), nil)
	if err != nil {
		_ = fmt.Errorf("error getting key: %v", err)
	}
	fmt.Println("value: ", string(value))

	// delete key from database
	if err = db.Delete([]byte("key"), nil); err != nil {
		_ = fmt.Errorf("error deleting key: %v", err)
	} else {
		fmt.Println("deleted key")
	}

	// get not exist key from database
	value, err = db.Get([]byte("key"), nil)
	if err != nil {
		_ = fmt.Errorf("error getting key: %v", err)
	}
	fmt.Println("value: ", string(value))
}
