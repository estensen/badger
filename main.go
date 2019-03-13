package main

import (
	"fmt"
	"github.com/dgraph-io/badger"
)

func main() {
	opts := badger.DefaultOptions
	opts.Dir = "./data"
	opts.ValueDir = "./data"
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("answer"), []byte("42"))
		return err
	})
	if err != nil {
		panic(err)
	}

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("answer"))
		if err != nil {
			panic(err)
		}

		err = item.Value(func(val []byte) error {
			fmt.Printf("The answer is: %s\n", val)
			return nil
		})
		if err != nil {
			panic(err)
		}

		return nil
	})
}
