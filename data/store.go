package data

import (
	"github.com/coreos/bbolt"
)

type (
	//Store represents the model store base
	Store struct {
		db *bolt.DB //a field with a lowercase first letter will not be exported out of the package
	}
)

//Update mutates the database
//Bolt DB wraps mutations (data changes) in an Update transaction
//The transaction is marked successful if fn returns nil for error
func (store *Store) Update(fn func(tx *bolt.Tx) error) error {
	return store.db.Update(fn)
}

//View reads the database
//Bolt DB wraps reads in a readonly transaction
func (store *Store) View(fn func(tx *bolt.Tx) error) error {
	return store.db.View(fn)
}
